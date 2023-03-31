package models

type Login struct {
	Correo string `json:"correo" validate:"required,email"`
	Clave  string `json:"clave" validate:"required,min=5"`
}
