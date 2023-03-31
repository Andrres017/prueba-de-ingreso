package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/controllers"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/utils"
	"github.com/tu-usuario/mi-proyecto/validations"
)

func SetCamion(router *mux.Router) {
	validate := validator.New()
	validate.RegisterValidation("format_placa", validations.ValidarFormato)
	// validate.RegisterValidation("format_placa", func(fl validator.FieldLevel) bool {
	// 	return true
	// })

	subRoute := router.PathPrefix("/camion/api").Subrouter()
	subRoute.HandleFunc("/all", utils.WithAuth(controllers.GetAllCamiones)).Methods("GET")
	subRoute.HandleFunc("/find/{id}", utils.WithAuth(controllers.GetCamion)).Methods("GET")
	subRoute.HandleFunc("/findfilter/{campo}/{valor}", utils.WithAuth(controllers.GetFilterCamiones)).Methods("GET")
	subRoute.HandleFunc("/save", utils.WithAuth(validations.ValidateUniversal(&models.LogisticaCamione{}, "camion", validate)(controllers.SaveCamion))).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", utils.WithAuth(controllers.DeleteCamion)).Methods("DELETE")
}
