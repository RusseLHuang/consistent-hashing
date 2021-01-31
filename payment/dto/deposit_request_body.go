package dto

type DepositRequestBody struct {
	UserID uint `json:"userId"`
	Amount uint `json:"amount"`
}
