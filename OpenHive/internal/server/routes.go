package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.LoadHTMLGlob("./front/html/*.html")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": "Página inicial - Base de atendimento"})
	})
	r.GET("/persons", func(c *gin.Context) {
		c.HTML(200, "persons.html", gin.H{"title": "Pessoas - Base de atendimento"})
	})
	r.GET("/tickets", func(c *gin.Context) {
		c.HTML(200, "tickets.html", gin.H{"title": "Pessoas - Base de atendimento"})
	})

	return r
}
