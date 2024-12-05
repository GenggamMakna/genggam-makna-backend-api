package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type MLResponse struct {
	Confidence        float64     `json:"confidence,omitempty"`
	PredictedAlphabet interface{} `json:"predicted_alphabet,omitempty"`
	Error             string      `json:"error,omitempty"`
}
