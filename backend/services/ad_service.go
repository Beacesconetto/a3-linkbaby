package services

import (
	"fmt"
	"linkbaby/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAnuncio(newAnuncio models.Anuncio) (models.Anuncio, error) {
	collection := models.DB.Database("linkbaby").Collection("anuncios")

	_, err := collection.InsertOne(models.Ctx, bson.M{
		"descricao":  newAnuncio.Descricao,
		"usuario_id": newAnuncio.UsuarioID,
		"preco":      newAnuncio.Preco,
		"localidade": newAnuncio.Localidade,
	})

	if err != nil {
		return models.Anuncio{}, fmt.Errorf("erro ao criar anúncio: %v", err)
	}

	return newAnuncio, nil
}

func GetAnunciosByEmail(email string) ([]models.Anuncio, error) {
	var anuncios []models.Anuncio

	collection := models.DB.Database("linkbaby").Collection("anuncios")

	cursor, err := collection.Aggregate(models.Ctx, []bson.M{
		{
			"$lookup": bson.M{
				"from":         "usuarios",
				"localField":   "usuario_id",
				"foreignField": "_id",
				"as":           "usuario",
			},
		},
		{
			"$match": bson.M{
				"usuario.email": email,
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar anúncios: %v", err)
	}

	if err := cursor.All(models.Ctx, &anuncios); err != nil {
		return nil, fmt.Errorf("erro ao decodificar os anúncios: %v", err)
	}

	return anuncios, nil
}

func GetAllAnuncios() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio

	collection := models.DB.Database("linkbaby").Collection("anuncios")

	cursor, err := collection.Find(models.Ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar todos os anúncios: %v", err)
	}

	if err := cursor.All(models.Ctx, &anuncios); err != nil {
		return nil, fmt.Errorf("erro ao decodificar os anúncios: %v", err)
	}

	return anuncios, nil
}
