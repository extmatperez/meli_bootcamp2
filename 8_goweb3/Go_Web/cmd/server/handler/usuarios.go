package handler

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	usuarios "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/internal/usuarios"
	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/pkg/web"
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

// ListProducts godoc
// @Summary List usuarios
// @Tags Usuario
// @Description get usuarios
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios/get [get]
func (control *Usuario) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !validarToken(c) {
			return
		}

		usuarios, err := control.service.GetAll()
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
		} else {
			c.JSON(200, web.NewResponse(200, usuarios, ""))
		}
	}
}

// StoreProducts godoc
// @Summary Store usuario
// @Tags Usuario
// @Description store usuario
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param usuario body request true "Usuario to store"
// @Success 200 {object} web.Response
// @Router /usuarios/add [post]
func (control *Usuario) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !validarToken(c) {
			return
		}

		var newUser request
		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar un usuario %v", err)))
		} else {
			err := validarUsuario(newUser)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				response, err := control.service.Store(newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura, newUser.Activo, newUser.FechaCreacion)
				if err != nil {
					c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar la persona %v", err)))
				} else {
					c.JSON(200, web.NewResponse(200, response, ""))
				}
			}
		}
	}
}

// UpdateProducts godoc
// @Summary Update usuario
// @Tags Usuario
// @Description update usuario
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param usuario body request true "Usuario to update"
// @Success 200 {object} web.Response
// @Router /usuarios/update [put]
func (control *Usuario) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !validarToken(c) {
			return
		}

		var updateUser request
		err := c.ShouldBindJSON(&updateUser)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			err := validarUsuario(updateUser)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				response, err := control.service.Update(updateUser.ID, updateUser.Nombre, updateUser.Apellido, updateUser.Email, updateUser.Edad, updateUser.Altura, updateUser.Activo, updateUser.FechaCreacion)
				if err != nil {
					c.JSON(404, web.NewResponse(404, nil, err.Error()))
				} else {
					c.JSON(200, web.NewResponse(200, response, ""))
				}
			}
		}
	}
}

// DeleteUsuario godoc
// @Summary Delete usuario
// @Tags Usuario
// @Description delete usuario
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} web.Response
// @Router /usuarios/delete/:id [delete]
func (control *Usuario) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !validarToken(c) {
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "El ID ingresado no es válido"))
		} else {
			err := control.service.Delete(int(id))
			if err != nil {
				c.JSON(404, web.NewResponse(404, nil, err.Error()))
			} else {
				c.JSON(200, web.NewResponse(200, fmt.Sprintf("Usuario %d ha sido eliminado correctamente", id), ""))
			}
		}
	}
}

// PatchProducts godoc
// @Summary Patch usuario
// @Tags Usuario
// @Description patch usuario
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param usuario body request true "Usuario to patch"
// @Success 200 {object} web.Response
// @Router /usuarios/patch [patch]
func (control *Usuario) EditarNombreEdad() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !validarToken(c) {
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			var userPatch request
			err := c.ShouldBindJSON(&userPatch)
			fmt.Println(userPatch)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				user, err := control.service.EditarNombreEdad(int(id), userPatch.Nombre, userPatch.Edad)

				if err != nil {
					c.JSON(404, web.NewResponse(404, nil, err.Error()))
				} else {
					c.JSON(200, web.NewResponse(200, user, ""))
				}
			}
		}

	}
}

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")

	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "Imposible validar token"))
		return false
	}

	token_env := os.Getenv("TOKEN")

	if token != token_env {
		ctx.JSON(400, web.NewResponse(400, nil, "Token incorrecto"))
		return false
	}
	return true
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
