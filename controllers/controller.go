package controllers

import (
	"errors"
	"net/http"

	"github.com/Santiagososs/MyGolangApi/models"
	"github.com/gin-gonic/gin"
)

func GetCandidatos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Candidatos)
}

func NewCandidate(c *gin.Context) {
	var Candidato models.Candidato

	if err := c.BindJSON(&Candidato); err != nil {
		return
	}

	models.Candidatos = append(models.Candidatos, Candidato)
	c.IndentedJSON(http.StatusCreated, models.Candidatos)
}

var uniqCandidate models.Candidato

func CandidateByTicket(c *gin.Context) {
	ticketnumber := c.Param("ticketnumber")

	uniqCandidate, err := GetCandidateById(ticketnumber)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Ticket inválido!!!"})
		return
	}

	c.IndentedJSON(http.StatusOK, uniqCandidate)
}

func GetCandidateById(ticketnumber string) (*models.Candidato, error) {
	for i, b := range models.Candidatos {
		if b.TicketNumber == ticketnumber {
			return &models.Candidatos[i], nil
		}
	}
	return nil, errors.New("Ticket inválido...")
}
