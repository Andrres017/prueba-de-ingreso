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

func ValidateUniversal(model interface{}, typeStr string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			validate := validator.New()
			var err error
			switch typeStr {
			case "persona":
				err = json.NewDecoder(r.Body).Decode(model.(*models.Persona))
			case "maritima":
				err = json.NewDecoder(r.Body).Decode(model.(*models.LogisticaMaritima))

			case "camion":
				err = json.NewDecoder(r.Body).Decode(model.(*models.LogisticaCamione))
			}

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err = validate.Struct(model)
			// Validamos el modelo de la persona
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, typeStr, model)
			next(w, r.WithContext(ctx))
		}
	}
}
