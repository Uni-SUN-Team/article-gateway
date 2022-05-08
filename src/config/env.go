package config

import "os"

func ConfigENV() {
	os.Setenv("PORT", "8082")
	os.Setenv("CONTEXT_PATH", "/article-listening")
	// Strapi information gateway
	os.Setenv("HOST_STRAPI_SERVICE", "http://localhost:8081")
	os.Setenv("PATH_STRAPI_INFORMATION_GATEWAY", "/strapi-information-gateway/api/strapi")
	// Path
	os.Setenv("PATH_STRAPI_ARTICLE", "/api/articles")
}
