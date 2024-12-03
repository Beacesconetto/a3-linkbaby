package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var Ctx = context.Background()

// Conectar ao MongoDB
func ConnectDatabase() {
	mongoURI := "mongodb+srv://beacesconetto203:7Gw6Vusz92JKGllA@cluster0.w4ruo.mongodb.net/"
	if mongoURI == "" {
		mongoURI = "mongodb+srv://beacesconetto203:7Gw6Vusz92JKGllA@cluster0.w4ruo.mongodb.net/"
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
