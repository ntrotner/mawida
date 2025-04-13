package openapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupSwaggerUi(mux *mux.Router) {
	sh := http.StripPrefix("/swaggerui", http.FileServer(http.Dir("./open-api/swaggerui/")))
	mux.PathPrefix("/swaggerui/").Handler(sh)
}
