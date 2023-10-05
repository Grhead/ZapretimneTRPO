package Service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Books struct {
	Title      string  `db:"Title" json:"title"`
	Author     string  `db:"Author" json:"author"`
	Year       int     `db:"Year" json:"year"`
	Publisher  string  `json:"publisher"`
	Genre      string  `json:"genre"`
	Popularity float32 `db:"Popularity" json:"popularity"`
}

type Genres struct {
	Title string `db:"Title" json:"genres"`
}

type GenreStruct struct {
	Genres []string `form:"genresbox[]"`
}

var db *gorm.DB
var selectedGenresGlobal []Genres

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("Hro.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetBook() []Books {
	var ResArrw []Books
	db.Joins("JOIN Genres on Genres.Id = Books.Genre").
		Joins("JOIN Publishers on Publishers.Id = Books.Publisher").
		Select("Books.title", "Author", "Year", "Publishers.Title as Publisher", "Genres.Title as Genre", "Popularity").
		Find(&ResArrw)
	return ResArrw
}

func GetGenres() []Genres {
	var ResArr []Genres
	db.Find(&ResArr)
	return ResArr
}

func GetBookWhere() []Books {
	var ResArr []Books
	var GenresArr []int
	if len(selectedGenresGlobal) < 1 {
		return ResArr
	}
	for _, i := range selectedGenresGlobal {
		log.Println(i)
		db.Raw("SELECT Id FROM Genres WHERE Title = ?", i.Title).Scan(&GenresArr)
	}
	db.Joins("JOIN Genres on Genres.Id = Books.Genre").
		Joins("JOIN Publishers on Publishers.Id = Books.Publisher").
		Select("Books.title", "Author", "Year", "Publishers.Title as Publisher", "Genres.Title as Genre", "Popularity").
		Where("Genres.Id IN ?", GenresArr).
		Find(&ResArr)
	return ResArr
}

func GetBookWhereOrderNormal() []Books {
	var ResArr []Books
	var GenresArr []int
	if len(selectedGenresGlobal) < 1 {
		return ResArr
	}
	for _, i := range selectedGenresGlobal {
		log.Println(i)
		db.Raw("SELECT Id FROM Genres WHERE Title = ?", i.Title).Scan(&GenresArr)
	}
	db.Joins("JOIN Genres on Genres.Id = Books.Genre").
		Joins("JOIN Publishers on Publishers.Id = Books.Publisher").
		Select("Books.title", "Author", "Year", "Publishers.Title as Publisher", "Genres.Title as Genre", "Popularity").
		Where("Genres.Id IN ?", GenresArr).
		Order("Popularity").
		Find(&ResArr)
	return ResArr
}

func GetBookWhereOrderDesc() []Books {
	var ResArr []Books
	var GenresArr []int
	if len(selectedGenresGlobal) < 1 {
		return ResArr
	}
	for _, i := range selectedGenresGlobal {
		log.Println(i)
		db.Raw("SELECT Id FROM Genres WHERE Title = ?", i.Title).Scan(&GenresArr)
	}
	db.Joins("JOIN Genres on Genres.Id = Books.Genre").
		Joins("JOIN Publishers on Publishers.Id = Books.Publisher").
		Select("Books.title", "Author", "Year", "Publishers.Title as Publisher", "Genres.Title as Genre", "Popularity").
		Order("Popularity DESC").
		Where("Genres.Id IN ?", GenresArr).
		Find(&ResArr)
	return ResArr
}
