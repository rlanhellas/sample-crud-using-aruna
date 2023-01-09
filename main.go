package main

import (
	"github.com/rlanhellas/aruna"
	"github.com/rlanhellas/aruna/httpbridge"
	"github.com/rlanhellas/sample-crud-using-aruna/client"
	"github.com/rlanhellas/sample-crud-using-aruna/shared"
	"net/http"
)

func main() {
	aruna.Run(&aruna.RunRequest{
		RoutesHttp: []*httpbridge.RouteHttp{
			{Path: "/v1/client", Method: http.MethodPost, Handler: client.Create,
				HandlerInputGenerator: shared.NewClient},
			{Path: "/v1/client/:id", Method: http.MethodDelete, Handler: client.Delete},
			{Path: "/v1/client", Method: http.MethodPut, Handler: client.Update, HandlerInputGenerator: shared.NewClient},
			{Path: "/v1/client/:id", Method: http.MethodGet, Handler: client.GetById},
		},
		MigrateTables: []any{&shared.Client{}},
	})
}
