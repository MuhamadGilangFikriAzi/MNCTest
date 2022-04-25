package api

import (
	"github.com/gin-gonic/gin"
	"gokost.com/m/authenticator"
	"gokost.com/m/delivery/apprequest"
	"gokost.com/m/delivery/common_resp"
	"gokost.com/m/delivery/logger"
	"gokost.com/m/usecase"
	"net/http"
)

type loginApi struct {
	usecase     usecase.LoginAdminUsecase
	configToken authenticator.Token
}

func (l *loginApi) LoginAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataLogin apprequest.AdminRequest
		if errBind := c.ShouldBindJSON(&dataLogin); errBind != nil {
			logger.SendLogToDiscord("Login", errBind)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errBind.Error()))
			return
		}
		dataAdmin, is_available, err := l.usecase.LoginAdmin(dataLogin)
		if err != nil {
			logger.SendLogToDiscord("Login", err)
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

func NewLoginApi(routeGroup *gin.RouterGroup, adminUsecase usecase.LoginAdminUsecase, configToken authenticator.Token) {
	api := &loginApi{
		adminUsecase,
		configToken,
	}

	routeGroup.POST("/admin", api.LoginAdmin())
}
