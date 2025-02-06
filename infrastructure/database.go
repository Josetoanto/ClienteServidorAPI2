package infrastructure

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open("mysql", "root:tono1234@tcp(localhost:3306)/nombre_base_de_datos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	fmt.Println("Conexión exitosa a la base de datos")
}
