package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"pulsy/internal/firebase"
)

func UploadFile(bucketName, fileName string, file *multipart.FileHeader) error {

	client := firebase.StorageClient

	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return fmt.Errorf("storageBucket is not initialized")
	}

	// Crear el escritor para el archivo en el bucket
	writer := bucket.Object(fileName).NewWriter(firebase.GetNewContext())

	// Abrir el archivo para lectura
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer src.Close()

	_, err = io.Copy(writer, src)
	if err != nil {
		return fmt.Errorf("error writing to bucket: %v", err)
	}
	defer writer.Close()

	return nil
}

func DownloadFile(bucketName, fileName string, localPath string) error {
	client := firebase.StorageClient

	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return fmt.Errorf("storageBucket is not initialized: %v", err)
	}

	// Crear el lector para leer el archivo del bucket
	reader, err := bucket.Object(fileName).NewReader(firebase.GetNewContext())
	if err != nil {
		return fmt.Errorf("failed to create reader for file %s: %v", fileName, err)
	}
	defer reader.Close()

	// Crear el archivo local
	file, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("failed to create local file %s: %v", localPath, err)
	}
	defer file.Close()

	// Copiar los datos del lector al archivo local
	if _, err := io.Copy(file, reader); err != nil {
		return fmt.Errorf("failed to write to local file %s: %v", localPath, err)
	}

	return nil
}

func DeleteFile(bucketName, fileName string) error {
	client := firebase.StorageClient

	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return fmt.Errorf("storageBucket is not initialized")
	}

	// Referenciar el archivo a eliminar
	object := bucket.Object(fileName)

	// Eliminar el archivo
	if err := object.Delete(firebase.GetNewContext()); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}
