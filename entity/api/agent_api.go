package dossier

import (
	"dmp/entity"
	"dmp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// agentLoginPayloadValidator : for validation of agent login
type agentLoginPayloadValidator struct {
	Matricule string `json:"matricule" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

//AgentLoginAPI : API to login agent
func AgentLoginAPI(c *gin.Context) {
	var payload agentLoginPayloadValidator
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	agent, err := entity.FindAgentByMatricule(payload.Matricule)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"response_content": "unkonwn-agent-matricule", "response_code": "100"})
		return
	}
	result, err := entity.CheckAgentCredentials(agent, payload.Password)
	if !result {
		c.JSON(http.StatusUnauthorized, gin.H{"response_content": "unkonwn-agent-credentials", "response_code": "100"})
		return
	}
	// Create token for the agent and sent it in the result
	token, err := utils.CreateToken(agent.Matricule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response_content": "internal-error", "response_code": "100"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response_content": gin.H{"token": token, "agent": agent}, "response_code": "000"})
}
