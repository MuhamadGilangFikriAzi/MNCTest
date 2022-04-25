package api

import (
	"github.com/gin-gonic/gin"
	"mnctest.com/api/authenticator"
	"mnctest.com/api/delivery/apprequest"
	"mnctest.com/api/delivery/common_resp"
	"mnctest.com/api/usecase"
	"net/http"
)

type loginApi struct {
	usecase     usecase.LoginCustomerUsecase
	configToken authenticator.Token
}

func (l *loginApi) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		l.configToken.RemoveToken()
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("success logout", ""))
	}
}

func (l *loginApi) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataLogin apprequest.CustomerRequest
		if errBind := c.ShouldBindJSON(&dataLogin); errBind != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			return
		}
		dataAdmin, is_available, err := l.usecase.LoginCustomer(dataLogin)
		if err != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		if !is_available {
			common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage("not register"))
			return
		}
		tokenString, errToken := l.configToken.CreateToken(dataAdmin)
		if errToken != nil {
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage("Token Failed"))
			return
		}
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("login admin", gin.H{
			"token": tokenString,
		}))
	}
}

func NewLoginApi(routeGroup *gin.RouterGroup, adminUsecase usecase.LoginCustomerUsecase, configToken authenticator.Token) {
	api := &loginApi{
		adminUsecase,
		configToken,
	}

	routeGroup.POST("", api.Login())
	routeGroup.POST("/logout", api.Logout())
}
