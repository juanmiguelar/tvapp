package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// LoadConfig loads configuration using Viper
func LoadConfig() {
	viper.SetConfigName(".env")    // Name of the config file
	viper.SetConfigType("env")     // File type (dotenv format)
	viper.AddConfigPath(".")       // Look for the file in the current directory

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Optionally, read from environment variables
	viper.AutomaticEnv()
}

// Connect establishes a connection to the MongoDB instance
func Connect() {
	LoadConfig()

	// Read the MongoDB URI from configuration
	uri := viper.GetString("MONGO_URI")
	if uri == "" {
		log.Fatalf("MONGO_URI is not set in the configuration")
	}

	// Set up MongoDB connection
	var err error
	clientOptions := options.Client().ApplyURI(uri)
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
}
