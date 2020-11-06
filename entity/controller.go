package entity

import (
	"dmp/utils"
)

// CheckAgentCredentials : function to check agent credentials
func CheckAgentCredentials(agent Agent, password string) (bool, error) {
	match, err := utils.ComparePassword(password, agent.HashedPassword)
	if !match || err != nil {
		return false, err
	}
	return true, nil
}
