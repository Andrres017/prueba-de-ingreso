package models

type Persona struct {
	ID        *int64 `json:"id" gorm:"primary_key;auto_increment"`
	Nombre    string `json:"nombre" validate:"required,min=1,max=255"`
	Apellido  string `json:"apellido" validate:"required,min=1,max=255"`
	Direccion string `json:"direccion" validate:"required,min=1,max=255"`
	Telefono  string `json:"telefono" validate:"required,numeric,min=1,max=255"`
	Correo    string `json:"correo" validate:"required,email"`
	Clave     string `json:"clave" validate:"required,min=5"`
}
