package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/controllers"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/utils"
	"github.com/tu-usuario/mi-proyecto/validations"
)

func SetMaritima(router *mux.Router) {
	validate := validator.New()
	validate.RegisterValidation("format_placa_maritima", validations.ValidarFormatoMaritima)
	subRoute := router.PathPrefix("/maritima/api").Subrouter()
	subRoute.HandleFunc("/all", utils.WithAuth(controllers.GetAllMaritimas)).Methods("GET")
	subRoute.HandleFunc("/find/{id}", utils.WithAuth(controllers.GetMaritima)).Methods("GET")
	subRoute.HandleFunc("/findfilter/{campo}/{valor}", utils.WithAuth(controllers.GetFilterMaritima)).Methods("GET")
	subRoute.HandleFunc("/save", utils.WithAuth(validations.ValidateUniversal(&models.LogisticaMaritima{}, "maritima", validate)(controllers.SaveMaritima))).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", utils.WithAuth(controllers.DeleteMaritima)).Methods("DELETE")
}
