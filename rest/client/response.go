package client

type ReferenceBasedResponse struct {
	Data    []string `json:"data"`
	Status  bool     `json:"status"`
	Message string   `json:"message"`
}
