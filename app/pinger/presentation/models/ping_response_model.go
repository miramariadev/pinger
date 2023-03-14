package models

type PingResponseModel struct {
	Method string `json:"method"`
	Result string `json:"result"`
	Error  string `json:"error"`
}

func NewPingResponseModel(result string) *PingResponseModel {
	return &PingResponseModel{
		Method: "ping",
		Result: result,
	}
}
