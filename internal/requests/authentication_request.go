package requests

type TokenRequestByUserID struct {
	UserID string `json:"userID" binding:"required"`
}
