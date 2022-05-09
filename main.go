package main

import (
	"os"
	"unisun/api/article-listening/src"
	"unisun/api/article-listening/src/config"
	"unisun/api/article-listening/src/constants"
)

func main() {
	if os.Getenv(constants.NODE) != constants.PRODUCTION {
		config.ConfigENV()
	}
	r := src.App()
	port := os.Getenv(constants.PORT)
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
