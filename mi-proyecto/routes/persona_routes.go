package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/controllers"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/utils"
	"github.com/tu-usuario/mi-proyecto/validations"
)

func SetPersona(router *mux.Router) {
	validate := validator.New()
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", utils.WithAuth(controllers.GetAll)).Methods("GET")
	subRoute.HandleFunc("/find/{id}", utils.WithAuth(controllers.Get)).Methods("GET")
	subRoute.HandleFunc("/findfilter/{campo}/{valor}", utils.WithAuth(controllers.GetFilterPersona)).Methods("GET")
	subRoute.HandleFunc("/save", validations.ValidateUniversal(&models.Persona{}, "persona", validate)(controllers.Save)).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", utils.WithAuth(controllers.Delete)).Methods("DELETE")
	subRoute.HandleFunc("/login", controllers.Getlogin).Methods("POST")
}
