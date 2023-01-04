package client

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rlanhellas/aruna/httpbridge"
	"github.com/rlanhellas/aruna/logger"
	"github.com/rlanhellas/sample-crud-using-aruna/shared"
	"net/http"
)

// Create client
func Create(ctx context.Context, req any, _ gin.Params) *httpbridge.HandlerHttpResponse {
	client := req.(*shared.Client)
	logger.Info(ctx, "Creating client: Client: %v", *client)
	return &httpbridge.HandlerHttpResponse{
		Data:       client,
		StatusCode: http.StatusCreated,
	}
}

func Delete(ctx context.Context, _ any, params gin.Params) *httpbridge.HandlerHttpResponse {
	if id, ok := params.Get("id"); ok {
		logger.Info(ctx, "Deleting client %s", id)
		return &httpbridge.HandlerHttpResponse{
			StatusCode: http.StatusOK,
			Data:       id,
		}
	} else {
		return &httpbridge.HandlerHttpResponse{
			Error:      errors.New("can not find param id"),
			StatusCode: http.StatusBadRequest,
		}
	}
}
func Update() {}
func Get()    {}
