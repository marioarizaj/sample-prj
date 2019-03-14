package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/marioarizaj/sample-project/models"
	"github.com/marioarizaj/sample-project/services"
	"net/http"
	"strconv"
)

type ArticleController struct {
	articleService *services.ArticleService
}

func NewArticleController(service *services.ArticleService) *ArticleController {
	return &ArticleController{
		articleService:service,
	}
}


func (ac *ArticleController) CreateArticle(c *gin.Context) {
	var article *models.Article

	err := c.MustBindWith(&article, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}


	err = ac.articleService.CreateArticle(article)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (ac *ArticleController) UpdateArticle(c *gin.Context) {
	var article *models.Article

	err := c.MustBindWith(&article, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = ac.articleService.UpdateArticle(article)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	var article *models.Article

	err := c.MustBindWith(&article, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = ac.articleService.DeleteArticle(article)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (ac *ArticleController) GetArticles(c *gin.Context) {
	articles, err := ac.articleService.GetArticles()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (ac *ArticleController) GetArticlesById(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if idint == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Id should be bigger than 0",
		})
	}

	article, err := ac.articleService.GetArticleById(idint)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, article)
}