package api

import (
	"github.com/gin-gonic/gin"
	"mnctest.com/api/delivery/apprequest"
	"mnctest.com/api/delivery/common_resp"
	"mnctest.com/api/usecase"
	"net/http"
)

type transactionApi struct {
	usecase usecase.BalanceTransferUseCase
}

func (api *transactionApi) BalanceTransfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data apprequest.TransactionRequest
		c.ShouldBindJSON(&data)
		dataReturn, msg, err := api.usecase.BalanceTransfer(data)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage(msg, dataReturn))
	}
}

func NewTransactionApi(routeGroup *gin.RouterGroup, usecase usecase.BalanceTransferUseCase) {
	api := transactionApi{usecase}

	routeGroup.POST("", api.BalanceTransfer())
}
