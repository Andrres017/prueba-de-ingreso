package main

import (
	"log"
	"net/http"

	"github.com/tu-usuario/mi-proyecto/common"
	"github.com/tu-usuario/mi-proyecto/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// func corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
// 		next.ServeHTTP(w, r)
// 	})
// }

func main() {
	common.Migrate()

	router := mux.NewRouter()
	//router.Use(corsMiddleware)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
	)
	routes.SetPersona(router)
	routes.SetCamion(router)
	routes.SetMaritima(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: cors(router),
	}

	log.Println("Servidor ejecutandoce sobre el puerto: 9000")
	log.Println(server.ListenAndServe())
}
