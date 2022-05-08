package route

import (
	"unisun/api/article-listening/src/controller"

	"github.com/gin-gonic/gin"
)

func Consumer(r *gin.RouterGroup) {
	r.GET("/articles", controller.ArticleAll)
	r.GET("/articles/:id", controller.ArticleById)
}
