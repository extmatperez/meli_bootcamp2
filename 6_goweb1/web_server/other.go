
// ACEPTAR UN PARAMETRO POR URL
/* router.GET("/productname/:nombre", func(c *gin.Context) {
	nombre := c.Param("nombre")
	c.String(http.StatusOK, "Hello %s", nombre)
}) */

// func getAll(c *gin.Context) {

// 	var prodList []Products
// 	readProducts, _ := os.ReadFile("./products.json")

// 	if err := json.Unmarshal([]byte(readProducts), &prodList); err != nil {
// 		log.Fatal(err)
// 	}

// 	c.JSON(200, gin.H{
// 		"data": prodList,
// 	})
// }
