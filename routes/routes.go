package routes

import (
	"github.com/Santiagososs/MyGolangApi/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()
	router.GET("/candidatos", controllers.GetCandidatos)
	router.POST("/candidatos", controllers.NewCandidate)
	router.GET("/candidatos/:ticketnumber", controllers.CandidateByTicket)
	router.Run("localhost:3000")
}
