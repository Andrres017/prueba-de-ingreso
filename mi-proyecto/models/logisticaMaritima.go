package models

type LogisticaMaritima struct {
	ID               *int64 `json:"id" gorm:"primary_key;auto_increment"`
	TipoProducto     string `json:"tipoProducto" validate:"required,min=1,max=255"`
	CantidadProducto uint16 `json:"cantidadProducto" validate:"required,min=1,max=255,numeric"`
	FechaRegistro    string `json:"fechaRegistro" validate:"required"`
	FechaEntrega     string `json:"fechaEntrega" validate:"required"`
	PuertoEntrega    string `json:"puertoEntrega" validate:"required,min=1,max=255"`
	PrecioEnvio      string `json:"precioEnvio" validate:"required,min=1,max=255,number"`
	NumeroFlota      string `json:"numeroFlota" validate:"required,min=8,max=8,format_placa_maritima"`
	NumeroGuia       string `json:"numeroGuia"`
}
