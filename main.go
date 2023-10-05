package main

import (
	"ZapretiMne/Service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
	}))
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.GET("/bookslist", Service.GetBooksApi)
	route.GET("/genreslist", Service.GetGenresApi)
	route.POST("/forma", Service.SelectedFormGenres)
	route.GET("/orderwhere", Service.GetOrderWhere)
	route.GET("/ordernormal", Service.GetOrderNormal)
	route.GET("/orderdesc", Service.GetOrderDesc)

	err := route.Run(":8085")

	if err != nil {
		return
	}
}
