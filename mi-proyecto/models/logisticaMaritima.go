package models

type LogisticaMaritima struct {
	ID               int64  `json:"id" gorm:"primary_key;auto_increment"`
	TipoProducto     string `json:"tipoProducto" validate:"required,min=1,max=255"`
	CantidadProducto uint16 `json:"cantidadProducto" validate:"required,min=1,max=255,numeric"`
	FechaRegistro    string `json:"fechaRegistro" validate:"required,datetime"`
	FechaEntrega     string `json:"fechaEntrega" validate:"required,datetime"`
	PuertoEntrega    string `json:"puertoEntrega" validate:"required,min=1,max=255"`
	PrecioEnvio      string `json:"precioEnvio" validate:"required,min=1,max=255"`
	NumeroFlota      string `json:"numeroFlota" validate:"required,min=8,max=8"`
	NumeroGuia       string `json:"numeroGuia" gorm:"uniqueIndex" validate:"required,min=10,max=10"`
}
