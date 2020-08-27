package entity

import (
	"context"
	"dmp/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
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
