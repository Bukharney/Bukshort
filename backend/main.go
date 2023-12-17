package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/Bukharney/shorterURL/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type link struct {
	gorm.Model
	ShortURL    string `gorm:"unique"`
	OriginalURL string `gorm:"unique"`
}

type ShortenReq struct {
	URL string `json:"url" binding:"required"`
}

type ShortenRes struct {
	URL string `json:"url"`
}

type ErrorRes struct {
	Error string `json:"error"`
}

func RandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// CreateTags godoc
// @Summary Shorten a URL
// @Description Shorten a URL
// @Param url body ShortenReq true "URL to shorten"
// @Produce application/json
// @Tags shorten
// @Router /shorten [post]
// @Success 200 {object} ShortenRes
// @Failure 400 {object} ErrorRes
// @Failure 500 {object} ErrorRes
func Shorten(c *gin.Context, db *gorm.DB) {
	var data ShortenReq
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var url link
	err := db.Where("original_url = ?", data.URL).First(&url).Error
	if err == nil {
		c.JSON(200, gin.H{"url": url.ShortURL})
		return
	}

	url = link{
		ShortURL:    RandomString(6),
		OriginalURL: data.URL,
	}

	re := db.Create(&url)
	for re.Error != nil {
		url.ShortURL = RandomString(6)
		re = db.Create(&url)
	}

	if re.Error != nil {
		c.JSON(500, gin.H{"error": err.Error})
		return
	}

	c.JSON(200, gin.H{"url": url.ShortURL})
}

// GetURL godoc
// @Summary Get a URL
// @Description Get a URL
// @Produce application/json
// @Tags shorten
// @Router /{shortURL} [get]
// @Param shortURL path string true "Short URL"
// @Success 200 {object} ShortenRes
// @Failure 404 {object} ErrorRes
func GetURL(c *gin.Context, db *gorm.DB) {
	var url link
	if err := db.Where("short_url = ?", c.Param("shortURL")).First(&url).Error; err != nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(200, gin.H{"url": url.OriginalURL})
}

func GetSomethings(c *gin.Context, db *gorm.DB) {
	c.Redirect(302, "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley")
}

func DbConfig() (*gorm.DB, error) {

	db_host := "localhost"
	db_port := "5432"
	db_user := "postgres"
	db_name := "postgres"
	db_password := "postgres"

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	if host == "railway" {
		db_host = os.Getenv("DB_HOST")
		db_port = os.Getenv("DB_PORT")
		db_user = os.Getenv("DB_USERNAME")
		db_name = os.Getenv("DB_DATABASE")
		db_password = os.Getenv("DB_PASSWORD")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&link{})
	return db, nil
}

// @title URL Shortener API
// @version 1
// @description This is a URL shortener API

// @host localhost:8080
func main() {
	log.Println("Starting server...")

	db, err := DbConfig()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	r := gin.Default()

	r.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://bukshort.bukharney.tech", "https://shorter-url-bukharney.vercel.app", "*"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Content-Type"},
			AllowCredentials: true,
		}),
	)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		GetSomethings(c, db)
	})

	r.POST(
		"/shorten", func(c *gin.Context) {
			Shorten(c, db)
		},
	)

	r.GET(
		"/:shortURL",
		func(c *gin.Context) {
			GetURL(c, db)
		},
	)

	r.Run()
}
