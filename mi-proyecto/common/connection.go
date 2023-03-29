package common

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tu-usuario/mi-proyecto/models"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Andresfelipe01@/prueba?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando.....")

	db.AutoMigrate(&models.Persona{})
	db.AutoMigrate(&models.LogisticaCamione{})
	db.AutoMigrate(&models.LogisticaMaritima{})

}
