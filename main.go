package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marioarizaj/sample-project/controllers"
	"github.com/marioarizaj/sample-project/services"
	"github.com/marioarizaj/sample-project/utilities"
)

func main() {
	//utilities
	configUtil := utilities.NewConfigUtil()

	dbOrm := utilities.NewOrmDB(configUtil)

	articleService := services.NewArticleService(dbOrm)

	articleController := controllers.NewArticleController(articleService)

	router := gin.Default()

	api := router.Group("/api")

	api.POST("", articleController.CreateArticle)
	api.PUT("", articleController.UpdateArticle)
	api.GET("", articleController.GetArticles)
	api.GET("/:id", articleController.GetArticlesById)
	api.DELETE("", articleController.DeleteArticle)

	err := router.Run(configUtil.GetConfig("port"))
	if err != nil {
		panic(err)
	}
}
