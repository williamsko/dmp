package usager

import (
	"context"
	"dmp/db"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// IUsager : Interface de fonctions sur objet Usager
type IUsager interface {
	CreateNewUsager(usager *Usager) interface{}
	FindUsagerByPhoneNumber(phoneNumber string) Usager
}

// FindUsagerByPhoneNumber : Find usager
func FindUsagerByPhoneNumber(phoneNumber string) Usager {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	usager := Usager{}

	usagerCollection := db.ConnectDb().Collection("usager")
	result := usagerCollection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&usager)

	if result != nil {
		panic(result)
	}

	return usager
}

// CreateNewUsager : create a new usager
func CreateNewUsager(usager *Usager) interface{} {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if usager.CreatedAt.IsZero() {
		usager.CreatedAt = time.Now()
	}

	usagerCollection := db.ConnectDb().Collection("usager")
	result, err := usagerCollection.InsertOne(ctx, usager)

	if err != nil {
		panic(err)
	}

	fmt.Println("Nouvel usager crééé : ", usager.ID)

	return result.InsertedID
}
