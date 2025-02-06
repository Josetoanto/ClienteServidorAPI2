package application

import (
	"crud-carros-server2/infrastructure"
	"crud-carros-server2/domain"
	"fmt"
)



func ObtenerCarros() ([]domain.Car, error) {
	var cars []domain.Car
	if err := infrastructure.DB.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func ComprarCarro(id uint) error {
	var car domain.Car
	if err := infrastructure.DB.First(&car, id).Error; err != nil {
		return fmt.Errorf("carro no encontrado")
	}

	infrastructure.DB.Delete(&car)

	if err := infrastructure.DB.Exec("UPDATE cajas SET dinero = dinero + ?", car.Precio).Error; err != nil {
		return err
	}

	return nil
}
