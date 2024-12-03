package models

type Usuario struct {
	ID        int    `json:"id" bson:"_id,omitempty"`
	Nome      string `json:"nome" bson:"nome"`
	Email     string `json:"email" bson:"email"`
	Telefone  string `json:"telefone" bson:"telefone"`
	Senha     string `json:"senha" bson:"senha"`
	Categoria string `json:"categoria" bson:"categoria"`
}
