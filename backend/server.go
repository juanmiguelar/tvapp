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

const defaultPort = "8080"

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
	r.POST("/query", gin.WrapH(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))))
	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// Start the server
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
