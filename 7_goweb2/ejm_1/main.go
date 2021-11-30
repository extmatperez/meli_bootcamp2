package main

import "github.com/gin-gonic/gin"

type Entity struct {
	Id   int   `json:"id"`
}

var entities []Entity = []Entity{};

func main() {
	r := gin.Default()
	r.POST("/entity", func(c *gin.Context) {
		var entity Entity
		if err := c.ShouldBindJSON(&entity); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		entity.Id = entities[len(entities)-1].Id + 1
		entities = append(entities, entity)
		c.JSON(200, entity)
	})
	r.GET("/entities", func(c *gin.Context) {
		c.JSON(200, entities)
	}) 
	r.Run()
}