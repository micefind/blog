package article

import (
	"errors"
	"server/dao/mysql"
	db "server/dao/mysql/article"
	model "server/model/article"
)

var tableName = "article"

func CreateArticle(req model.Article) error {
	return db.CreateArticle(req)
}

func UpdateArticle(req model.Article) error {
	return db.UpdateArticle(req)
}

func DeleteArticle(req model.Article) error {
	// 检查文章是否存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("文章不存在")
	}
	return db.DeleteArticle(req)
}

func RecoverArticle(req model.Article) error {
	// 检查文章是否存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("文章不存在")
	}
	return db.RecoverArticle(req)
}

func GetArticleDetail(req model.Article) (model.ArticleDetail, error) {
	var articleInfo model.ArticleDetail
	// 检查文章是否存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return articleInfo, err
	}
	if !result {
		return articleInfo, errors.New("文章不存在")
	}
	return db.GetArticleDetail(req)
}

func GetArticleList(req model.Query) (model.List, error) {
	return db.GetArticleList(req)
}
