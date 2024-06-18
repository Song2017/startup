package middleware

import (
	"net/http"
	"strings"

	_pkg "startup/_pkg"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(securityKey, securityVal string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if securityVal != "" {
			api_resp := map[string]interface{}{"Code": 401}
			var authorization string
			authorization = c.GetHeader(securityKey)
			authorization = _pkg.Ternary(authorization == "",
				c.GetHeader(strings.ToLower(securityKey)), authorization)
			if token, ok := c.GetQuery(securityKey); ok && authorization == "" {
				authorization = token
			}
			// Token is the default access key
			if c.GetHeader("Token") != "" && authorization == "" {
				authorization = c.GetHeader("Token")
			}

			if authorization == "" {
				api_resp["Message"] = "API token is required"
				c.JSON(
					http.StatusUnauthorized,
					api_resp,
				)
				c.Abort()
				return
			}
			if authorization != securityVal {
				api_resp["Message"] = "Provided token is not valid"
				c.JSON(
					http.StatusUnauthorized,
					api_resp,
				)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
