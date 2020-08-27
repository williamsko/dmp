package usager

import (
	"context"
	"dmp/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// FindUsagerByPhoneNumber : Find usager
func FindUsagerByPhoneNumber(phoneNumber string) (Usager, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"phonenumber": phoneNumber}).Decode(&usager)
	return usager, err
}

// FindUsagerByMatricule : Find usager
func FindUsagerByMatricule(matricule string) (Usager, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"matricule": matricule}).Decode(&usager)
	return usager, err
}

// CreateNewUsager : create a new usager
func CreateNewUsager(usager *Usager) (*Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if usager.CreatedAt.IsZero() {
		usager.CreatedAt = time.Now()
	}
	usagerCollection := db.ConnectDb().Collection("usager")
	_, err := usagerCollection.InsertOne(ctx, usager)
	return usager, err
}
