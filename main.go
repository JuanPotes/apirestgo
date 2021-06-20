package main

import (
	"log"
	"net/http"

	"github.com/JuanPotes/api-rest-go/common"
	"github.com/JuanPotes/api-rest-go/routes"
	"github.com/gorilla/mux"
)

func main() {
	common.Migrate()

	router := mux.NewRouter()

	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose en puerto 9000")
	log.Println(server.ListenAndServe())
}
