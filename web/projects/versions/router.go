// Package versions implement /projects/{projectId}/versions routes.
package versions

import (
	"fmt"
	"net/http"

	"github.com/Palantir/palantir/web/pathvars"
	"github.com/Palantir/palantir/web/server"
	"github.com/gorilla/mux"
)

// HandleVersions define versions routes.
func HandleVersions(router *mux.Router, context *server.Context) {
	middlewares := context.ValidationMiddleware

	router.Handle("", middlewares(&createVersionHandler{context})).Methods(http.MethodPost)

	createFromRoute := fmt.Sprintf("/create/from/{%s}", pathvars.VersionID)
	router.Handle(createFromRoute,
		middlewares(&createFromExistingVersionHandler{context})).
		Methods(http.MethodPost)
}
