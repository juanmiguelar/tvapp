package main

import (
	"log"
	"os"
	"tvapp-backend/database"
	"tvapp-backend/graph"

	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Default port
const defaultPort = "8080"

// CORSMiddleware handles CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Continue to next handler
		c.Next()
	}
}

func main() {
	// Connect to MongoDB
	database.Connect()

	// Retrieve the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("PORT environment variable not set, using default port %s", port)
	}

	// Configure GraphQL server
	r := gin.Default()

	// Use the CORS middleware
	r.Use(CORSMiddleware())

	// Setup GraphQL routes
	r.POST("/query", gin.WrapH(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))))
	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// Start the server
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
