package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"unisun/api/article-listening/src/constants"
	"unisun/api/article-listening/src/model"
	"unisun/api/article-listening/src/service"

	"github.com/gin-gonic/gin"
)

func ArticleAll(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "?populate[categories]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		payloadRequest.Path += "&" + strings.Trim(query, "?")
	}
	var articles model.Articles
	data := service.GetArticles(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &articles)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "result": articles.Data, "pagination": articles.Meta.Pagination})
}

func ArticleById(c *gin.Context) {
	id := c.Param("id")
	var articles model.Article
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "/" + id + "?populate[categories]=*&populate[advisors][populate]=thumnail&populate[courses][populate][0]=thumnail&populate[users_permissions_users]=*&populate[thumnail]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetArticles(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &articles)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "result": articles.Data})
}

func ArticleBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var articles model.ArticlesSlug
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "?populate[categories]=*&populate[advisors][populate]=*&populate[courses][populate][0]=thumnail&populate[users_permissions_users]=*&populate[thumnail]=*&filters[slug][$eq]=" + slug
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetArticles(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &articles)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	article := articles.Data[0]
	c.JSON(http.StatusOK, gin.H{"error": err, "result": article})
}
