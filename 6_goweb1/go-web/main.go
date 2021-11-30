package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	datos, _ := os.ReadFile("./transaccion.json")
	json.Unmarshal(datos, &transacciones)

	router := gin.Default()

	router.GET("/saludar/:name", saludar)
	router.GET("/saludar", saludar)
	router.GET("/transacciones", getAll)
	router.GET("/transacciones/:id", getById)

	router.Run()

}

func String(file string) {
	panic("unimplemented")
}

// estructura

type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaCreacion     string  `json:"fecha_creacion"`
}

var transacciones []Transaccion
var filtrados []Transaccion

//

func saludar(c *gin.Context) {
	//queryName := c.Request.URL.Query()  // esto devuelve un map de string string
	queryName := c.Query("name")
	paramName := c.Param("name")

	if queryName != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola " + queryName,
		})
	} else if paramName != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola " + paramName,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola, ingresa tu nombre como query o param",
		})
	}

}

func getById(c *gin.Context) {
	filterId := c.Param("id")
	var filtId Transaccion

	if filterId != "" {
		for _, v := range transacciones {
			if filterId == strconv.Itoa(v.ID) {
				filtId = v
				break
			}
		}

	}

	if filtId.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"transacciones": filtId,
		})

	} else {
		c.JSON(http.StatusBadRequest, "el id que buscas no se encuentra")
	}

}

func getAll(c *gin.Context) {
	filterId := c.Query("id")
	filterCodTrans := c.Query("codigo_transaccion")
	filterMoneda := c.Query("moneda")
	filterMonto := c.Query("monto")
	filterEmisor := c.Query("emisor")
	filterReceptor := c.Query("receptor")
	filterFecha := c.Query("fecha_creacion")

	if filterId != "" {

		if len(filtrados) == 0 {
			for i, v := range transacciones {
				if c.Query("id") == strconv.Itoa(v.ID) {
					filtrados = append(filtrados, transacciones[i])
				}
			}

		} else {
			for i, v := range filtrados {
				if c.Query("id") == strconv.Itoa(v.ID) {
					filtrados = append(filtrados, filtrados[i])
				}
			}

		}

	}
	if filterCodTrans != "" {
		if len(filtrados) == 0 {
			for i, v := range transacciones {
				if c.Query("filterCodTrans") == v.CodigoTransaccion {
					filtrados = append(filtrados, transacciones[i])
				}
			}

		} else {
			for i, v := range filtrados {
				if c.Query("filterCodTrans") == v.CodigoTransaccion {
					filtrados = append(filtrados, filtrados[i])
				}
			}

		}
	}
	if filterMoneda != "" {
		if len(filtrados) == 0 {
			for i, v := range transacciones {
				if c.Query("filterMoneda") == v.Moneda {
					filtrados = append(filtrados, transacciones[i])
				}
			}

		} else {
			for i, v := range filtrados {
				if c.Query("filterMoneda") == v.Moneda {
					filtrados = append(filtrados, filtrados[i])
				}
			}

		}
	}
	if filterMonto != "" {
		if len(filtrados) == 0 {
			for i, v := range transacciones {
				if c.Query("filterMonto") == fmt.Sprint(v.Monto) {
					filtrados = append(filtrados, transacciones[i])
				}
			}

		} else {
			for i, v := range filtrados {
				if c.Query("filterMonto") == fmt.Sprint(v.Monto) {
					filtrados = append(filtrados, filtrados[i])
				}
			}

		}
	}
	if filterEmisor != "" {
		if len(filtrados) == 0 {
			for i, v := range transacciones {
				if c.Query("filterEmisor") == v.Emisor {
					filtrados = append(filtrados, transacciones[i])
				}
			}

		} else {
			for i, v := range filtrados {
				if c.Query("filterEmisor") == v.Emisor {
					filtrados = append(filtrados, filtrados[i])
				}
			}

		}
	}
	if filterReceptor != "" {
		if len(filtrados) == 0 {
			for i, v := range transacciones {
				if c.Query("filterReceptor") == v.Receptor {
					filtrados = append(filtrados, transacciones[i])
				}
			}

		} else {
			for i, v := range filtrados {
				if c.Query("filterReceptor") == v.Receptor {
					filtrados = append(filtrados, filtrados[i])
				}
			}

		}
	}
	if filterFecha != "" {
		if len(filtrados) == 0 {
			for i, v := range transacciones {
				if c.Query("filterFecha") == v.FechaCreacion {
					filtrados = append(filtrados, transacciones[i])
				}
			}

		} else {
			for i, v := range filtrados {
				if c.Query("filterFecha") == v.FechaCreacion {
					filtrados = append(filtrados, filtrados[i])
				}
			}

		}

	}

	if len(filtrados) != 0 && (filterId != "" || filterCodTrans != "" || filterMoneda != "" || filterMonto != "" || filterEmisor != "" || filterReceptor != "" || filterFecha != "") {
		c.JSON(http.StatusOK, gin.H{
			"transacciones": filtrados,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"transacciones": transacciones,
		})

	}

}
