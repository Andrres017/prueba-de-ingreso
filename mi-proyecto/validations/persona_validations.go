package validations

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tu-usuario/mi-proyecto/models"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

/*
func ValidateMiddlewarePersona(v *validator.Validate) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Parseamos el cuerpo de la solicitud en una estructura de la persona
			var persona models.Persona
			if err := json.NewDecoder(r.Body).Decode(&persona); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Validamos el modelo de la persona
			if err := v.Struct(persona); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Si la validación es exitosa, llamamos al controlador original
			next(w, r)
		}
	}
}*/

func ValidateUniversal(model interface{}, typeStr string, validate *validator.Validate) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var modelConvert any
			var err error
			switch typeStr {
			case "persona":
				modelConvert = new(models.Persona)
				err = json.NewDecoder(r.Body).Decode(modelConvert.(*models.Persona))
			case "maritima":
				modelConvert = new(models.LogisticaMaritima)
				err = json.NewDecoder(r.Body).Decode(modelConvert.(*models.LogisticaMaritima))

			case "camion":
				modelConvert = new(models.LogisticaCamione)
				err = json.NewDecoder(r.Body).Decode(modelConvert.(*models.LogisticaCamione))
			}

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest) //400
				return
			}
			err = validate.Struct(modelConvert)
			// Validamos el modelo de la persona
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity) //422
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, typeStr, modelConvert)
			next(w, r.WithContext(ctx))
		}
	}
}
