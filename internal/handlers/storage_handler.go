package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"pulsy/internal/firebase"
	"pulsy/internal/responses"
	"pulsy/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFileHandler(c *gin.Context) {
	// Obtener el archivo de la solicitud
	file, err := c.FormFile("file")

	if err != nil {
		message := "file not found"
		errData := map[string]interface{}{
			"field": "file",
			"cause": "file is required",
		}

		c.IndentedJSON(http.StatusBadRequest, responses.Error(message, errData))
		return
	}

	// Generar un UUID único y construir el nombre del archivo en el bucket
	fileUUID := uuid.New().String()
	fileFullName := file.Filename
	fileExtension := filepath.Ext(fileFullName)
	fileName := strings.Replace(fileFullName, fileExtension, "", -1)
	fileSize := file.Size
	bucketFileName := fileUUID + fileExtension

	metadata := map[string]interface{}{
		"uuid":         fileUUID[:8],
		"fullName":     fileFullName,
		"name":         fileName,
		"extension":    fileExtension,
		"size":         fileSize,
		"creationDate": time.Now(),
		"uuidOwner":    "rafiprogramando",
	}

	// Guardar archivo en Storage
	err = services.UploadFile(os.Getenv("FIREBASE_STORAGE_BUCKET"), bucketFileName, file)

	if err != nil {
		message := "failed to upload file"
		errData := map[string]interface{}{
			"fileName":      fileName,
			"fileExtension": fileExtension,
			"cause":         err.Error(),
		}

		c.IndentedJSON(http.StatusConflict, responses.Error(message, errData))
		return
	}

	// Guardar metadatos en Firestore
	err = services.InsertData("files", fileUUID, metadata)

	if err != nil {
		// Eliminar el archivo del bucket si hubo un error
		bucketErr := services.DeleteFile(os.Getenv("FIREBASE_STORAGE_BUCKET"), bucketFileName)

		var message string

		if bucketErr != nil {
			message = "failed to delete file"
			err = bucketErr
		} else {
			message = "failed to insert metadata"
		}

		errData := map[string]interface{}{
			"fileName":      fileName,
			"fileExtension": fileExtension,
			"cause":         err.Error(),
		}

		c.IndentedJSON(http.StatusConflict, responses.Error(message, errData))
		return
	}

	// Responder con éxito
	message := "files uploaded successfully"
	errData := map[string]interface{}{
		"uuid":          fileUUID,
		"fileName":      fileName,
		"fileExtension": fileExtension,
		"fileSize":      fileSize,
	}

	c.IndentedJSON(http.StatusAccepted, responses.Success(message, errData))
}

func DownloadFileHandler(c *gin.Context) {
	fileUUID := c.Param("fileUUID") // Nombre del archivo en Firebase

	var bucketFileName string
	var fileFullName string

	// Si UUID tiene 8 caracteres, se busca en Firestore
	if len(fileUUID) == 8 {
		// Criterios de búsqueda
		conditions := []firebase.QueryCondition{
			{Field: "uuid", Operator: "==", Value: fileUUID},
		}

		// Obtener metadatos del archivo
		metadata, err := services.ReadMultipleDocs("files", conditions)

		if err != nil {
			message := "failed to get file metadata"
			errData := map[string]interface{}{
				"fileUUID": fileUUID,
				"cause":    err.Error(),
			}

			c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
			return
		}

		if len(metadata) == 0 {
			message := "file not found"
			errData := map[string]interface{}{
				"fileUUID": fileUUID,
			}

			c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
			return
		}

		if len(metadata) > 1 {
			message := "multiple files found"
			errData := map[string]interface{}{
				"fileUUID": fileUUID,
				"action":   "insert a file full uuid",
			}

			c.IndentedJSON(http.StatusMultipleChoices, responses.Error(message, errData))
			return
		}

		fileUUID = metadata[0]["idDoc"].(string)
		bucketFileName = fileUUID + metadata[0]["extension"].(string)
		fileFullName = metadata[0]["fullName"].(string)
	} else {
		_, err := uuid.Parse(fileUUID)

		if err != nil {
			message := "uuid not valid"
			errData := map[string]interface{}{
				"fileUUID": fileUUID,
				"action":   "verify the uuid",
			}

			c.IndentedJSON(http.StatusBadRequest, responses.Error(message, errData))
			return
		}

		// Obtener metadatos del archivo
		metadata, err := services.ReadDoc("files", fileUUID)

		if err != nil {
			message := "failed to get metadata"
			errData := map[string]interface{}{
				"fileUUID": fileUUID,
				"cause":    err.Error(),
			}

			c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
			return
		}

		bucketFileName = fileUUID + metadata["extension"].(string)
		fileFullName = metadata["fullName"].(string)
	}

	/// Crear archivo temporal seguro
	tempFile, err := os.CreateTemp("", fileFullName)
	if err != nil {
		message := "failed to create temporary file"
		errData := map[string]interface{}{
			"fileUUID": fileUUID,
			"cause":    err.Error(),
		}

		c.IndentedJSON(http.StatusInternalServerError, responses.Error(message, errData))
		return
	}
	defer os.Remove(tempFile.Name()) // Se asegura que el archivo sea eliminado al finalizar

	// Descargar el archivo al sistema local
	err = services.DownloadFile(os.Getenv("FIREBASE_STORAGE_BUCKET"), bucketFileName, tempFile.Name())
	if err != nil {
		message := "failed to download file"
		errData := map[string]interface{}{
			"fileUUID": fileUUID,
			"cause":    err.Error(),
		}

		c.IndentedJSON(http.StatusInternalServerError, responses.Error(message, errData))
		return
	}

	// Enviar el archivo como respuesta
	c.Header("Content-Disposition", "attachment; filename="+fileFullName)
	c.File(tempFile.Name())

	/* message := "file found"
	errData := map[string]interface{}{
		"fileUUID": fileUUID,
	}

	c.IndentedJSON(http.StatusOK, responses.Success(message, errData)) */
}
