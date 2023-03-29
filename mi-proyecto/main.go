package main

import (
	"log"
	"net/http"

	"github.com/tu-usuario/mi-proyecto/common"
	"github.com/tu-usuario/mi-proyecto/routes"

	"github.com/gorilla/mux"
)

func main() {
	common.Migrate()

	router := mux.NewRouter()
	routes.SetPersona(router)
	routes.SetCamion(router)
	routes.SetMaritima(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandoce sobre el puerto: 9000")
	log.Println(server.ListenAndServe())
}
