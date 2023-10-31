package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

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

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	log.Println("Starting server...")

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USERNAME")
	db_name := os.Getenv("DB_DATABASE")
	db_password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&link{})

	r := gin.Default()

	r.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://urlshorter.bukharney.tech/"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			AllowCredentials: true,
		}),
	)

	r.POST(
		"/shorten",
		func(c *gin.Context) {
			var data struct {
				URL string `json:"url" binding:"required"`
			}

			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			var url link
			if err := db.Where("original_url = ?", data.URL).First(&url).Error; err != nil {
				url = link{
					ShortURL:    RandomString(6),
					OriginalURL: data.URL,
				}
				re := db.Create(&url)
				if re.Error != nil {
					c.JSON(500, gin.H{"error": err.Error})
					return
				}
			}

			c.JSON(200, gin.H{"url": url.ShortURL})
		},
	)

	r.GET(
		"/:shortURL",
		func(c *gin.Context) {
			var url link
			if err := db.Where("short_url = ?", c.Param("shortURL")).First(&url).Error; err != nil {
				c.JSON(404, gin.H{"error": "URL not found"})
				return
			}

			c.Redirect(302, url.OriginalURL)
		},
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run()
}
