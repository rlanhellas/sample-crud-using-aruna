package client

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rlanhellas/aruna/db"
	"github.com/rlanhellas/aruna/httpbridge"
	"github.com/rlanhellas/sample-crud-using-aruna/shared"
)

// @BasePath /v1/client

// Create godoc
//
// @Summary      Create client
// @Description  Create a new client in Postgres DB
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        client  body      shared.Client  true  "cliente a ser criado"
// @Success      201  {object}  shared.Client
// @Router       /v1/client [post]
func Create(ctx context.Context, req any, _ gin.Params) *httpbridge.HandlerHttpResponse {
	c := req.(*shared.Client)
	c.Id = int(db.GetSequenceId("client_table_id_seq"))

	l := shared.Language{Id: 10, Name: "portuguese"}
	db.Create(ctx, &l)

	//c.Languages = []shared.Language{{Name: "english"}}
	c.Languages = []shared.Language{l}

	return db.CreateWithBindHandlerHttp(ctx, c)
}

// @BasePath /v1/client

// Delete godoc
//
// @Summary      Delete client
// @Description  Delete a client from Postgres DB
// @Tags         v1
// @Produce      json
// @Param        id   path      int  true  "cliente a ser deletado"
// @Success      201     {object}  shared.Client
// @Router       /v1/client/{id} [delete]
func Delete(ctx context.Context, _ any, params gin.Params) *httpbridge.HandlerHttpResponse {
	client, err := configureClientWithParamID(params)
	if err != nil {
		return &httpbridge.HandlerHttpResponse{Error: err,
			StatusCode: http.StatusBadRequest}
	}

	return db.DeleteWithBindHandlerHttp(ctx, client)
}

// @BasePath /v1/client

// Update godoc
//
// @Summary      Update client
// @Description  Update a client in Postgres DB
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        client  body      shared.Client  true  "cliente a ser atualizado"
// @Success      200  {object}  shared.Client
// @Router       /v1/client [put]
func Update(ctx context.Context, req any, _ gin.Params) *httpbridge.HandlerHttpResponse {
	c := req.(*shared.Client)
	return db.UpdateWithBindHandlerHttp(ctx, c)
}

// @BasePath /v1/client

// GetById godoc
//
// @Summary      Get client by id
// @Description  Get a client from Postgres DB by id
// @Tags         v1
// @Produce      json
// @Param        id   path      int  true  "cliente a ser recuperado"
// @Success      200     {object}  shared.Client
// @Router       /v1/client/{id} [get]
func GetById(ctx context.Context, _ any, params gin.Params) *httpbridge.HandlerHttpResponse {
	client, err := configureClientWithParamID(params)
	if err != nil {
		return &httpbridge.HandlerHttpResponse{Error: err,
			StatusCode: http.StatusBadRequest}
	}

	return db.GetByIdWithBindHandlerHttp(ctx, client, []string{"Languages"})
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
