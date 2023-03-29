package routes

import (
	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/controllers"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/validations"
)

func SetPersona(router *mux.Router) {

	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
	subRoute.HandleFunc("/save", validations.ValidateUniversal(&models.Persona{}, "persona")(controllers.Save)).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
}
