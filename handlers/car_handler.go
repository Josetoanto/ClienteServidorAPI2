package handlers

import (
	"crud-carros-server2/application"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ExhibirCarros(c *gin.Context) {
    cars, err := application.ObtenerCarros()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo carros"})
        return
    }

    c.Header("Content-Type", "text/event-stream")
    c.Header("Cache-Control", "no-cache")
    c.Header("Connection", "keep-alive")
    c.Header("Transfer-Encoding", "chunked")

    flusher, ok := c.Writer.(http.Flusher)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "El streaming no es compatible con este servidor"})
        return
    }

    for _, car := range cars {
        carJSON, err := json.Marshal(car)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al convertir el carro a JSON"})
            return
        }

        _, err = c.Writer.Write(append(carJSON, '\n'))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escribir el carro en el stream"})
            return
        }

        flusher.Flush()

        time.Sleep(2 * time.Second)
    }
}





func ComprarCarro(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := application.ComprarCarro(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carro comprado exitosamente"})
}
