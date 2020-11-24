package types

//base response struct
type ResponseJson struct {
	Code      int         `json:"code"`
	RequestId string      `json:"requestId"`
	Message   interface{} `json:"message"`
	Payload   interface{} `json:"data"`
}
