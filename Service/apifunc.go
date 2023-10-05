package Service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetBooksApi(c *gin.Context) {
	c.JSON(200,
		GetBook())
}

func GetGenresApi(c *gin.Context) {
	c.JSON(200,
		GetGenres())
}

func SelectedFormGenres(c *gin.Context) {
	var selectedGenres GenreStruct
	err := c.ShouldBind(&selectedGenres)
	if err != nil {
		fmt.Println(err)
		return
	}
	selectedGenresGlobal = append(selectedGenresGlobal[:len(selectedGenresGlobal)], selectedGenresGlobal[:]...)
	for _, i := range selectedGenres.Genres {
		selectedGenresGlobal = append(selectedGenresGlobal, Genres{Title: i})
	}
	c.JSON(200, gin.H{"SelGen": selectedGenres.Genres})
}

func GetOrderWhere(c *gin.Context) {
	var selectedGenres GenreStruct
	c.ShouldBind(&selectedGenres)
	c.JSON(200, GetBookWhere())
}

func GetOrderDesc(c *gin.Context) {
	var selectedGenres GenreStruct
	c.ShouldBind(&selectedGenres)
	c.JSON(200, GetBookWhereOrderDesc())
}

func GetOrderNormal(c *gin.Context) {
	var selectedGenres GenreStruct
	c.ShouldBind(&selectedGenres)
	c.JSON(200, GetBookWhereOrderNormal())
}
