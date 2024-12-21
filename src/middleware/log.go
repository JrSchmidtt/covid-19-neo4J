package middleware

import (
	"github.com/JrSchmidtt/covid-19-neo4J/pkg/id"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := id.NewIdWithPrefix("T")
		c.Set("traceID", traceID)
		c.Next()
	}
}