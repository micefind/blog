package tag

import (
	"errors"
	"server/dao/mysql"
	db "server/dao/mysql/article/tag"
	model "server/model/article/tag"
)

var tableName = "article_tag"

func CreateTag(req model.Tag) error {
	// 检查标签是否存在
	result, err := mysql.CheckTableExist(tableName, "name", req.Name)
	if err != nil {
		return err
	}
	if result {
		return errors.New("标签已存在")
	}
	return db.CreateTag(req)
}

func UpdateTag(req model.Tag) error {
	// 检查标签是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("标签不存在")
	}
	// 获取标签信息
	tagInfo, err := GetTagDetail(req)
	if err != nil {
		return err
	}
	// 如果标签名称有修改
	if tagInfo.Name != req.Name {
		// 检查标签是否存在
		result, err := mysql.CheckTableExist(tableName, "name", req.Name)
		if err != nil {
			return err
		}
		if result {
			return errors.New("标签已存在")
		}
	}
	return db.UpdateTag(req)
}

func DeleteTag(req model.Tag) error {
	// 检查标签是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("标签不存在")
	}
	return db.DeleteTag(req)
}

func GetTagDetail(req model.Tag) (model.TagDetail, error) {
	var tagInfo model.TagDetail
	// 检查标签是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return tagInfo, err
	}
	if !result {
		return tagInfo, errors.New("标签不存在")
	}
	return db.GetTagDetail(req)
}

func GetTagList(req model.Query) (model.List, error) {
	return db.GetTagList(req)
}
