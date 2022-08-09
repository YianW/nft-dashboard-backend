package middlewares

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"tulip/backend/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Authorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := token.ExtractRole(c)
		if err != nil {
			c.String(http.StatusBadRequest, "TOKENERROR", err.Error())
			c.Abort()
		}

		// casbin rule enforcing
		res, err := e.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.String(http.StatusInternalServerError, "ERROR", err)
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			c.String(http.StatusForbidden, "FORBIDDEN", errors.New("unauthorized"))
			c.Abort()
			return
		}
	}

}
