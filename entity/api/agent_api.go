package record

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
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	agent, err := entity.FindAgentByMatricule(payload.Matricule)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "unkonwn-agent-matricule")
		return
	}
	result, err := entity.CheckAgentCredentials(agent, payload.Password)
	if !result {
		utils.RespondWithError(c, http.StatusUnauthorized, "unkonwn-agent-credentials")
		return
	}
	// Create token for the agent and sent it in the result
	token, err := utils.CreateToken(agent.Matricule)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "authentication-token-creation-error")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"token": token, "agent": agent})
}
