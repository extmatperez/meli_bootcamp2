package handler

import (
	"fmt"
	"reflect"
	"strconv"

	usuarios "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(serv usuarios.Service) *Usuario {
	return &Usuario{service: serv}
}

func (control *Usuario) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarios, err := control.service.GetAll()
		if err != nil {
			c.String(400, "Hubo un error %v", err)
		} else {
			c.JSON(200, usuarios)
		}
	}
}

func (control *Usuario) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser request
		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			c.String(400, "Hubo un error al querer cargar un usuario %v", err)
		} else {
			err := validarUsuario(newUser)
			if err != nil {
				c.String(400, err.Error())
			} else {
				response, err := control.service.Store(newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura, newUser.Activo, newUser.FechaCreacion)
				if err != nil {
					c.String(400, "No se pudo cargar la persona %v", err)
				} else {
					c.JSON(200, response)
				}
			}
		}
	}
}

func (control *Usuario) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser request
		err := c.ShouldBindJSON(&updateUser)
		if err != nil {
			c.String(400, err.Error())
		} else {
			err := validarUsuario(updateUser)
			if err != nil {
				c.String(400, err.Error())
			} else {
				response, err := control.service.Update(updateUser.ID, updateUser.Nombre, updateUser.Apellido, updateUser.Email, updateUser.Edad, updateUser.Altura, updateUser.Activo, updateUser.FechaCreacion)
				if err != nil {
					c.String(404, err.Error())
				} else {
					c.JSON(200, response)
				}
			}
		}
	}
}

func (control *Usuario) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			c.String(400, "El ID ingresado no es válido")
		} else {
			err := control.service.Delete(int(id))
			if err != nil {
				c.String(404, err.Error())
			} else {
				c.String(200, "Usuario %d ha sido eliminado correctamente", id)
			}
		}
	}
}

func (control *Usuario) EditarNombreEdad() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			c.String(400, err.Error())
		} else {
			var userPatch request
			err := c.ShouldBindJSON(&userPatch)
			fmt.Println(userPatch)
			if err != nil {
				c.String(400, err.Error())
			} else {
				user, err := control.service.EditarNombreEdad(int(id), userPatch.Nombre, userPatch.Edad)

				if err != nil {
					c.String(404, err.Error())
				} else {
					c.JSON(200, user)
				}
			}
		}

	}
}

func validarUsuario(usuario request) error {
	message := ""
	var fields []string
	fields = append(fields, "ID", "Nombre", "Apellido", "Email", "Edad", "Altura", "Activo", "FechaCreacion")
	user := reflect.TypeOf(usuario)
	for i := 1; i < user.NumField(); i++ {
		nombre := user.Field(i).Name
		valor := reflect.ValueOf(usuario).FieldByName(fields[i]).Interface()

		if valor == "" || valor == 0 || valor == nil {
			message += "El campo " + nombre + " es requerido\n"
		}
	}
	if message == "" {
		return nil
	} else {
		return fmt.Errorf(message)
	}
}
