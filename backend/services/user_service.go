package services

import (
	"errors"
	"linkbaby/models"

	"go.mongodb.org/mongo-driver/bson"
)

var users []models.Usuario

func CreateUser(newUser models.Usuario) (models.Usuario, error) {
	var existingUser models.Usuario
	err := models.DB.Database("linkbaby").Collection("usuarios").FindOne(models.Ctx, map[string]interface{}{"email": newUser.Email}).Decode(&existingUser)
	if err == nil {
		return models.Usuario{}, errors.New("email already exists")
	}

	err = models.InsertUser(newUser)
	if err != nil {
		return models.Usuario{}, err
	}

	return newUser, nil
}
func LoginUser(email, senha string) (models.Usuario, error) {
	var user models.Usuario

	collection := models.DB.Database("linkbaby").Collection("usuarios")

	err := collection.FindOne(models.Ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.Usuario{}, errors.New("invalid email or password")
	}

	if user.Senha != senha {
		return models.Usuario{}, errors.New("invalid email or password")
	}

	return user, nil
}

func GetAllUsers() []models.Usuario {
	return users
}

func DeleteUser(id int) (bool, error) {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}
