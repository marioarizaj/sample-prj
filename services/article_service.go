package services

import (
	"github.com/go-pg/pg"
	"github.com/marioarizaj/sample-project/models"
)

type ArticleService struct {
	db *pg.DB
}

func NewArticleService(db *pg.DB) *ArticleService {
	return &ArticleService{
		db: db,
	}
}

func (ps *ArticleService) CreateArticle(article *models.Article) error {
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}
	err = tx.Insert(article)
	return tx.Commit()
}


func (ps *ArticleService) GetArticles() ([]models.Article,error) {
	var articles []models.Article

	err := ps.db.Model(&models.Article{}).Select(&articles)

	return articles, err
}

func (ps *ArticleService) UpdateArticle(newArticle *models.Article) error {
	return ps.db.Update(newArticle)
}


func (ps *ArticleService) GetArticleById(id int) (*models.Article,error) {
	article := &models.Article{}
	err := ps.db.Model(&models.Article{}).Where("id = ?", id).Select(article)

	return article, err
}

func (ps *ArticleService) DeleteArticle(article *models.Article) error {
	return ps.db.Delete(article)
}
