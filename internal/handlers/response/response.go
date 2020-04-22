package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handle is the interface to handle responses at view level
type Handle interface {
	GetOkResponse(context *gin.Context, view string)
	GetBadResponse(context *gin.Context, view string)
	GetNotFoundResponse(context *gin.Context, view string)

	SetPayload(key string, val interface{})
}

// res is the struct that store data response for views
/* May saves dynamic data for example:
{
	"entity": entity, // object
	"error": "Error message" // error description
	"status: 200 // status response
}
*/
type res struct {
	payload map[string]interface{}
}

// NewHandle returns interface Handle
func NewHandle(payload map[string]interface{}) Handle {
	return &res{
		payload: payload,
	}
}

// GetResponse is the base response at view level
func (r *res) GetResponse(context *gin.Context, view string) {
	if _, ok := r.payload["status"]; !ok {
		r.SetPayload("status", http.StatusOK)
	}

	if _, ok := r.payload["error"]; ok {
		fmt.Println(r.payload["error"])
	}

	// supports rendering json/html
	content := SetType(
		context.DefaultQuery("type", TypeHTML),
	)

	if content == TypeHTML {
		context.HTML(http.StatusOK, view, r.payload)
	} else {
		context.JSON(http.StatusOK, r.payload)
	}
}

// GetOkResponse handles OK responses at view level
func (r *res) GetOkResponse(context *gin.Context, view string) {
	r.SetPayload("status", http.StatusOK)
	r.GetResponse(context, view)
}

// GetBadResponse handles Bad responses at view level
func (r *res) GetBadResponse(context *gin.Context, view string) {
	r.SetPayload("status", http.StatusBadRequest)
	r.GetResponse(context, view)
}

// GetNotFoundResponse handles Not Found responses at view level
func (r *res) GetNotFoundResponse(context *gin.Context, view string) {
	r.SetPayload("status", http.StatusNotFound)
	r.GetResponse(context, view)
}

// SetPayload sets payload by key
func (r *res) SetPayload(key string, val interface{}) {
	r.payload[key] = val
}