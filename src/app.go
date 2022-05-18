package src

import (
	"os"
	"unisun/api/article-listening/src/constants"
	"unisun/api/article-listening/src/route"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(os.Getenv(constants.CONTEXT_PATH) + "/api")
	{
		route.Consumer(g)
	}
	return r
}
