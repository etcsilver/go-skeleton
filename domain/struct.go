package domain

type Hello struct {
	Name string `json:"name"`
}

type SuccessRespose struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Success Message"`
}

type ErrorRespose struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Error Message"`
}
