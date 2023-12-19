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

	c.Redirect(302, url.OriginalURL)
}

func GetSomethings(c *gin.Context, db *gorm.DB) {
	c.Redirect(302, "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley")
}

func DbConfig() (*gorm.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_unix.go: %s environment variable not set.\n", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep secrets safe.
	var (
		db_user     = mustGetenv("DB_USER")
		db_password = mustGetenv("DB_PASS")
		db_host     = mustGetenv("INSTANCE_UNIX_SOCKET")
		db_name     = mustGetenv("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", db_host, db_user, db_password, db_name)
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

// @host gbukshort.bukharney.tech
// @BasePath /
// @schemes https
func main() {
	log.Println("Connecting to database...")
	db, err := DbConfig()
	if err != nil {
		log.Fatalf("Fatal Error in main.go: %s\n", err)
	}

	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	gin.SetMode(gin.ReleaseMode)

	r.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://bukshort.bukharney.tech", "https://shorter-url-bukharney.vercel.app", "https://gbukshort.bukharney.tech/", "*"},
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

	log.Println("Starting server...")
	post := os.Getenv("PORT")
	r.Run(
		fmt.Sprintf(":%s", post),
	)
}
