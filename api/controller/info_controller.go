package controller

import (
	"bookshop/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)
type infoController struct{
	Info model.Info
}

func NewInfoController(info model.Info) infoController{
	return infoController{info}
}

func (controller infoController) GetInfo(context *gin.Context){
	jsonInfo,err := json.Marshal(controller.Info)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Info{})
	}
	context.JSON(http.StatusOK, string(jsonInfo))
}