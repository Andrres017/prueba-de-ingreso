package models

type LogisticaCamione struct {
	ID               int64  `json:"id" gorm:"primary_key;auto_increment"`
	TipoProducto     string `json:"tipoProducto" validate:"required,min=1,max=255"`
	CantidadProducto uint16 `json:"cantidadProducto" validate:"required,min=1,max=255,numeric"`
	FechaRegistro    string `json:"fechaRegistro" validate:"required,datetime"`
	FechaEntrega     string `json:"fechaEntrega" validate:"required,datetime"`
	BodegaEntrega    string `json:"bodegaEntrega" validate:"required,min=1,max=255"`
	PrecioEnvio      string `json:"precioEnvio" validate:"required,min=1,max=255"`
	PlacaVehiculo    string `json:"placaVehiculo" validate:"required,min=6,max=6"`
	NumeroGuia       string `json:"numeroGuia" validate:"required,min=10,max=10"`
}
