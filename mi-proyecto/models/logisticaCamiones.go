package models

type LogisticaCamione struct {
	ID               *int64 `json:"id" gorm:"primary_key;auto_increment"`
	TipoProducto     string `json:"tipoProducto" validate:"required,min=1,max=255" error:"Tipo de producto se encuentra vacio"`
	CantidadProducto uint16 `json:"cantidadProducto" validate:"required,min=1,max=255,numeric" error:"Cantidad de producto vacio o no es un numero"`
	FechaRegistro    string `json:"fechaRegistro" validate:"required" error:"Fecha de registro es obligatorio"`
	FechaEntrega     string `json:"fechaEntrega" validate:"required" error:"Fecha de entrega es obligatorio"`
	BodegaEntrega    string `json:"bodegaEntrega" validate:"required,min=1,max=255" error:"Bodega es obligatorio"`
	PrecioEnvio      string `json:"precioEnvio" validate:"required,min=1,max=255,number" error:"Precio de envio es obligatorio y debe de ser un numero"`
	PlacaVehiculo    string `json:"placaVehiculo" validate:"required,min=6,max=6,format_placa" error:"Debe de ser obligatorio y una placa real"`
	NumeroGuia       string `json:"numeroGuia"`
}
