package models

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

var DB *mongo.Client
var Ctx = context.Background()

// Conectar ao MongoDB
func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb+srv://diegoMongo:08052021Bea@cluster0.yzp4w.mongodb.net/linkbady?retryWrites=true&w=majority&appName=Cluster0"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")

	DB = client
}

func InsertUser(newUser Usuario) error {
	collection := DB.Database("linkbaby").Collection("usuarios")

	// Inserindo o novo usuário no MongoDB
	_, err := collection.InsertOne(Ctx, bson.M{
		"nome":      newUser.Nome,
		"email":     newUser.Email,
		"telefone":  newUser.Telefone,
		"senha":     newUser.Senha,
		"categoria": newUser.Categoria,
	})

	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
		return err
	}

	return nil
}
