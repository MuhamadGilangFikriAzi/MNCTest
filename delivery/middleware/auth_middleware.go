package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"mnctest.com/api/authenticator"
	"mnctest.com/api/delivery/common_resp"
	"net/http"
	"strings"
)

type AuthTokenMiddleware struct {
	acctToken authenticator.Token
}

func NewAuthTokenMiddleware(configToken authenticator.Token) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		acctToken: configToken,
	}
}

func (a *AuthTokenMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/customers" || c.Request.URL.Path == "/customers/register" {
			c.Next()
		} else {
			h := authHeader{}
			err := c.ShouldBindHeader(&h)
			if err != nil {
				common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage(err.Error()))
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			if tokenString == "" {
				common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage("Unautherized"))
				return
			}
			isAvailable, err := a.acctToken.CheckTokenAvailable(tokenString)
			if err == redis.Nil || !isAvailable {
				common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage("Unautherized"))
				return
			}
			if !isAvailable {
				common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage("Unautherized"))
				return
			}
			token, errToken := a.acctToken.VerifAccessToken(tokenString)
			if errToken != nil {
				common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage(errToken.Error()))
				return
			}

			if token["iss"] == a.acctToken.GetAppName() {
				a.acctToken.UpdateToken(tokenString)
				c.Next()
			} else {
				common_resp.NewCommonResp(c).FailedResp(http.StatusUnauthorized, common_resp.FailedMessage("Unautherized"))
				return
			}

			c.Next()
		}
	}
}
