package server

import "github.com/gin-gonic/gin"

type request struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      int     `json:"code"`
	Published string  `json:"published"`
	Created   string  `json:"created"`
}

func getAll(c *gin.Context) handlers {

	// page, _ := strconv.ParseInt(c.Request.URL.Query().Get("page"), 10, 64)
	// limit, _ := strconv.ParseInt(c.Request.URL.Query().Get("limit"), 10, 64)

	// startIndex := (page - 1) * limit
	// endIndex := page * limit

	// var paginatedResults []Products
	// paginatedResults = prodList[startIndex:endIndex]
	token := c.GetHeader("token")
	if token != tokenPrueba {
		c.JSON(401, gin.H{
			"error": "token inválido",
		})
	} else {

		c.JSON(200, gin.H{
			"data": prodList,
		})
	}

}
