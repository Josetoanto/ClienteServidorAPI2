package application

import (
	"crud-carros-server2/infrastructure"
	"fmt"
)

// Car estructura para el carro
type Car struct {
	ID     uint    `json:"id"`
	Brand  string  `json:"brand"`
	Model  string  `json:"model"`
	Year   int     `json:"year"`
	Precio float64 `json:"precio"`
}

// ObtenerCarros obtiene todos los carros
func ObtenerCarros() ([]Car, error) {
	var cars []Car
	if err := infrastructure.DB.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

// ComprarCarro elimina un carro por ID y actualiza la caja
func ComprarCarro(id uint) error {
	var car Car
	if err := infrastructure.DB.First(&car, id).Error; err != nil {
		return fmt.Errorf("carro no encontrado")
	}

	// Eliminar el carro
	infrastructure.DB.Delete(&car)

	// Actualizar la caja
	if err := infrastructure.DB.Exec("UPDATE cajas SET dinero = dinero + ?", car.Precio).Error; err != nil {
		return err
	}

	return nil
}
