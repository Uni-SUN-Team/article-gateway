package config

import (
	"os"
	"unisun/api/article-listening/src/constants"
)

func ConfigENV() {
	os.Setenv(constants.PORT, "8080")
	os.Setenv(constants.CONTEXT_PATH, "/article-listening")
	// Strapi information gateway
	os.Setenv(constants.HOST_STRAPI_SERVICE, "https://api.unisun.dynu.com")
	os.Setenv(constants.PATH_STRAPI_INFORMATION_GATEWAY, "/strapi-information-gateway/api/strapi")
	// Path
	os.Setenv(constants.PATH_STRAPI_ARTICLE, "/api/articles")
}
