// main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/your-username/project/handlers"
	"github.com/your-username/project/services"
)

func main() {
	// Initialize services
	embedService := services.NewEmbeddingService()
	dbService := services.NewChromaDBService()
	documentService := services.NewDocumentService(embedService, dbService)
	searchService := services.NewSearchService(embedService, dbService)

	// Initialize router
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Document routes
	router.POST("/documents", handlers.CreateDocument(documentService))
	router.GET("/documents/:id", handlers.GetDocument(documentService))
	router.PUT("/documents/:id", handlers.UpdateDocument(documentService))
	router.DELETE("/documents/:id", handlers.DeleteDocument(documentService))
	router.GET("/documents", handlers.ListDocuments(documentService))

	// Search route
	router.POST("/search", handlers.Search(searchService))

	log.Fatal(router.Run(":8080"))
}
