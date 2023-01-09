package client

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rlanhellas/aruna/db"
	"github.com/rlanhellas/aruna/httpbridge"
	"github.com/rlanhellas/sample-crud-using-aruna/shared"
	"net/http"
	"strconv"
)

// Create client
func Create(ctx context.Context, req any, _ gin.Params) *httpbridge.HandlerHttpResponse {
	c := req.(*shared.Client)
	return db.CreateWithBindHandlerHttp(ctx, c)
}

// Delete client
func Delete(ctx context.Context, _ any, params gin.Params) *httpbridge.HandlerHttpResponse {
	client, err := configureClientWithParamID(params)
	if err != nil {
		return &httpbridge.HandlerHttpResponse{Error: err,
			StatusCode: http.StatusBadRequest}
	}

	return db.DeleteWithBindHandlerHttp(ctx, client)
}

// Update client
func Update(ctx context.Context, req any, _ gin.Params) *httpbridge.HandlerHttpResponse {
	c := req.(*shared.Client)
	return db.UpdateWithBindHandlerHttp(ctx, c)
}

// GetById client
func GetById(ctx context.Context, _ any, params gin.Params) *httpbridge.HandlerHttpResponse {
	client, err := configureClientWithParamID(params)
	if err != nil {
		return &httpbridge.HandlerHttpResponse{Error: err,
			StatusCode: http.StatusBadRequest}
	}

	return db.GetByIdWithBindHandlerHttp(ctx, client)
}

func configureClientWithParamID(params gin.Params) (*shared.Client, error) {
	if id, ok := params.Get("id"); ok {
		client := shared.NewClient().(*shared.Client)
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		client.Id = idInt
		return client, nil
	} else {
		return nil, errors.New("param id not found")
	}
}
