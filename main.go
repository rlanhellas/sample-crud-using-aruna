package main

import (
	"context"
	"net/http"

	"github.com/rlanhellas/aruna"
	"github.com/rlanhellas/aruna/httpbridge"
	"github.com/rlanhellas/sample-crud-using-aruna/client"
	"github.com/rlanhellas/sample-crud-using-aruna/shared"
)

func main() {
	aruna.Run(&aruna.RunRequest{
		RoutesGroup: []*httpbridge.RouteGroupHttp{
			{
				Authenticated: true,
				Path:          "/v1/client",
				Routes: []*httpbridge.RouteHttp{
					{Path: "/", Method: http.MethodPost, Handler: client.Create,
						HandlerInputGenerator: shared.NewClient},
					{Path: "/:id", Method: http.MethodDelete, Handler: client.Delete},
					{Path: "/", Method: http.MethodPut, Handler: client.Update, HandlerInputGenerator: shared.NewClient},
					{Path: "/:id", Method: http.MethodGet, Handler: client.GetById},
					{Path: "/", Method: http.MethodGet, Handler: client.ListByName},
				},
			},
		},
		MigrateTables: []any{&shared.Client{}, &shared.Language{}},
		BackgroundTask: func(ctx context.Context) {
			//
		},
	})
}
