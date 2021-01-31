package dto

type TransferRequestBody struct {
	FromID uint `json:"fromId"`
	ToID   uint `json:"toId"`
	Amount uint `json:"amount"`
}
