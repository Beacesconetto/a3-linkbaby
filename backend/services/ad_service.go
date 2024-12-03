package services

import (
	"fmt"
	"linkbaby/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAnuncio(newAnuncio models.Anuncio) (models.Anuncio, error) {
	// Conecta à coleção de anúncios no MongoDB
	collection := models.DB.Database("linkbaby").Collection("anuncios")

	// Insere o novo anúncio no MongoDB
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

	// Conecta à coleção de anúncios no MongoDB
	collection := models.DB.Database("linkbaby").Collection("anuncios")

	// Consulta no MongoDB para encontrar os anúncios do usuário baseado no email
	cursor, err := collection.Aggregate(models.Ctx, []bson.M{
		{
			"$lookup": bson.M{
				"from":         "usuarios",   // Nome da coleção de usuários
				"localField":   "usuario_id", // Campo do anúncio que referencia o usuário
				"foreignField": "_id",        // Campo do usuário que será comparado
				"as":           "usuario",    // Nome do campo resultante da junção
			},
		},
		{
			"$match": bson.M{
				"usuario.email": email, // Condição de filtro para o email do usuário
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar anúncios: %v", err)
	}

	// Decodifica os resultados no slice de anúncios
	if err := cursor.All(models.Ctx, &anuncios); err != nil {
		return nil, fmt.Errorf("erro ao decodificar os anúncios: %v", err)
	}

	return anuncios, nil
}

func GetAllAnuncios() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio

	// Conecta à coleção de anúncios no MongoDB
	collection := models.DB.Database("linkbaby").Collection("anuncios")

	// Recupera todos os anúncios
	cursor, err := collection.Find(models.Ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar todos os anúncios: %v", err)
	}

	// Decodifica os resultados no slice de anúncios
	if err := cursor.All(models.Ctx, &anuncios); err != nil {
		return nil, fmt.Errorf("erro ao decodificar os anúncios: %v", err)
	}

	return anuncios, nil
}
