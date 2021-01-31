package dto

type CreateUserRequestBody struct {
	Name       string `json:"name"`
	NationalID string `json:"nationalId"`
}
