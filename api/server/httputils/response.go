package httputils

import (
	"net/http"

	"github.com/mojo-zd/helm-crabstick/types"
)

//NewSuccessResponse
func NewSuccessResponse(id string, payload interface{}) *types.ResponseJson {
	return &types.ResponseJson{
		RequestId: id,
		Payload:   payload,
		Code:      http.StatusOK,
		Message:   "success",
	}
}

//NewFailedResponse
func NewFailedResponse(id string, message string) *types.ResponseJson {
	return &types.ResponseJson{
		RequestId: id,
		Message:   message,
		Code:      http.StatusInternalServerError,
	}
}
