package api

import (
	"github.com/gin-gonic/gin"
	"mnctest.com/api/delivery/apprequest"
	"mnctest.com/api/delivery/common_resp"
	"mnctest.com/api/usecase"
	"net/http"
	"strconv"
)

type customerApi struct {
	usecase usecase.CustomerUseCase
}

func (api *customerApi) GetAllCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataMeta apprequest.Meta
		dataMeta.Limit, _ = strconv.Atoi(c.Query("limit"))
		dataMeta.Skip, _ = strconv.Atoi(c.Query("skip"))
		data, respMeta, err := api.usecase.GetAllCustomer(dataMeta)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", gin.H{
			"customers": data,
			"meta":      respMeta,
		}))
	}
}

func (api *customerApi) SearchCustomerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CustomerRequest
		c.ShouldBindJSON(&data)
		resp, err := api.usecase.DetailCustomer(data.CustomerId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", resp))
	}
}

func (api *customerApi) CreateCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CustomerRequest
		err := c.ShouldBindJSON(&data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		dataCreate, err := api.usecase.CreateCustomer(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", dataCreate))
	}
}

func (api *customerApi) UpdateCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CustomerRequest
		err := c.ShouldBindJSON(&data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		errUpdate := api.usecase.UpdateCustomer(data.CustomerId, data)
		if errUpdate != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errUpdate.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func (api *customerApi) DeleteCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CustomerRequest
		c.ShouldBindJSON(&data)
		err := api.usecase.DeleteCustomer(data.CustomerId)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func (api *customerApi) RegisterCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.CustomerRequest
		c.ShouldBindJSON(&data)
		errUpdate := api.usecase.RegisterCustomer(data.CustomerId)
		if errUpdate != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errUpdate.Error()))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("Success", ""))
	}
}

func NewCustomerApi(routeGroup *gin.RouterGroup, usecase usecase.CustomerUseCase) {
	api := &customerApi{usecase}

	routeGroup.GET("", api.GetAllCustomer())
	routeGroup.GET("/search", api.SearchCustomerById())
	routeGroup.POST("", api.CreateCustomer())
	routeGroup.PUT("", api.UpdateCustomer())
	routeGroup.PUT("/register", api.RegisterCustomer())
	routeGroup.DELETE("", api.DeleteCustomer())
}
