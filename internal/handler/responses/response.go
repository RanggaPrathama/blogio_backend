package responses


type Response struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorResponses struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

