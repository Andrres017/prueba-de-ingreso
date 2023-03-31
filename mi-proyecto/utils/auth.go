package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tu-usuario/mi-proyecto/models"
)

func GenerateToken(user models.Persona) (string, error) {
	// Generamos un nuevo token
	token := jwt.New(jwt.SigningMethodHS256)

	// Establecemos los claims del token, que son los datos que queremos guardar
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Correo
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // El token expirará en 24 horas

	// Firmamos el token con una clave secreta
	tokenString, err := token.SignedString([]byte("putaku"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el token desde el header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len("Bearer "):]

		// Validar el token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Aquí deberías retornar la clave secreta que utilizaste para firmar el token
			return []byte("putaku"), nil
		})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	}
}
