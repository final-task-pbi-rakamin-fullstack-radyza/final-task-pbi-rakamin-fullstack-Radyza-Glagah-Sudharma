package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "project-final/helpers"
    "strings"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. No token provided."})
            c.Abort()
            return
        }

        token = strings.TrimPrefix(token, "Bearer ")

        claims, err := helpers.VerifyToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}
