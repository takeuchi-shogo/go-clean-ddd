package usecase

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Code  int   `json:"code"`
	Error error `json:"error"`
}
