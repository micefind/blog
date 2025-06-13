package cate

import (
	"errors"
	"server/dao/mysql"
	db "server/dao/mysql/article/cate"
	model "server/model/article/cate"
)

var tableName = "article_cate"

func CreateCate(req model.Cate) error {
	// 检查分类是否存在
	result, err := mysql.CheckTableExist(tableName, "name", req.Name)
	if err != nil {
		return err
	}
	if result {
		return errors.New("分类已经存在")
	}
	return db.CreateCate(req)
}

func UpdateCate(req model.Cate) error {
	// 检查分类是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("分类不存在")
	}
	// 获取分类信息
	cateInfo, err := db.GetCateDetail(req)
	if err != nil {
		return err
	}
	if cateInfo.Name != req.Name {
		// 检查分类是否存在
		result, err := mysql.CheckTableExist(tableName, "name", req.Name)
		if err != nil {
			return err
		}
		if result {
			return errors.New("分类已经存在")
		}
	}

	return db.UpdateCate(req)
}

func DeleteCate(req model.Cate) error {
	// 检查分类是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("分类不存在")
	}
	return db.DeleteCate(req)
}

func GetCateDetail(req model.Cate) (model.CateDetail, error) {
	var cateInfo model.CateDetail
	// 检查分类是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return cateInfo, err
	}
	if !result {
		return cateInfo, errors.New("分类不存在")
	}
	return db.GetCateDetail(req)
}

func GetCateList(req model.Query) (model.List, error) {
	return db.GetCateList(req)
}
