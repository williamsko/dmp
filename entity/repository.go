package entity

import (
	"context"
	"dmp/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAgentByMatricule : Find agent by using his matricule
func FindAgentByMatricule(matricule string) (Agent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var agent Agent
	agentCollection := db.ConnectDb().Collection("agent")
	err := agentCollection.FindOne(ctx, bson.M{"matricule": matricule}).Decode(&agent)
	return agent, err
}

// FindAgentByID : Find agent by using his ID
func FindAgentByID(id primitive.ObjectID) (Agent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var agent Agent
	agentCollection := db.ConnectDb().Collection("agent")
	err := agentCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&agent)
	return agent, err
}
