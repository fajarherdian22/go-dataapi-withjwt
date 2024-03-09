package controller

import (
	"gojwt/helper"
	"gojwt/model/domain"
	"gojwt/model/web"
	"gojwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataController interface {
	GetDataByFilter(c *gin.Context)
	GetDataAll(c *gin.Context)
}

type DataControllerImpl struct {
	DataService service.DataService
}

func NewDataController(DataService service.DataService) DataController {
	return &DataControllerImpl{
		DataService: DataService,
	}
}

func (controller *DataControllerImpl) GetDataByFilter(c *gin.Context) {
	var filter domain.Data4G
	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataResponse := controller.DataService.GetDataByFilter(c.Request.Context(), *filter.Kecamatan, *filter.Rpt_area)
	WebResponse := web.WebResponse{
		Code:   200,
		Data:   dataResponse,
		Status: true,
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}

func (controller *DataControllerImpl) GetDataAll(c *gin.Context) {
	dataResponse := controller.DataService.GetAllData(c.Request.Context())
	WebResponse := web.WebResponse{
		Code:   200,
		Data:   dataResponse,
		Status: true,
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}
