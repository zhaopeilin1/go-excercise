package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"home-system/pkg/app"
	"home-system/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func (c *gin.Context) {
		var (
			token string
			ecode=errcode.Success
		)
		if s,exists := c.GetQuery("token");exists {
			token =s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_,err:=app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode=errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response:=app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()

	}
}
