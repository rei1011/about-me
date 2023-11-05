package handler

import (
	"about-me/handler/response"
	"about-me/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	useCase := usecase.NewProfileUseCase()
	list := useCase.GetOrganizationProfile()
	c.IndentedJSON(http.StatusOK, response.NewProfileListRes(list))
}
