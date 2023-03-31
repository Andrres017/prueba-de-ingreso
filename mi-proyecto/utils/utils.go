package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Sha1Hex(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}

func Descuento(numeroStr string, porcentaje int) string {
	numero, err := strconv.Atoi(numeroStr)
	if err != nil {
		return ""
	}
	descuento := float64(numero) * float64(porcentaje) / 100.0
	total := float64(numero) - descuento
	return fmt.Sprintf("%.2f", total)
}

func GenerateCode(length int) string {
	// Definimos los caracteres que pueden estar en el código generado
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	// Creamos un source de caracteres aleatorios
	rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())

	// Creamos un nuevo slice de bytes con la longitud especificada
	bytes := make([]byte, length)

	// Generamos caracteres aleatorios para cada posición en el slice
	for i := range bytes {
		bytes[i] = charset[source.Int63()%int64(len(charset))]
	}

	// Convertimos el slice de bytes a un string y lo devolvemos
	return string(bytes)
}
