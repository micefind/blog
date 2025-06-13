package cate

import (
	"errors"
	"fmt"
	"server/config"
	"server/dao/mysql"
	model "server/model/article/cate"
	"strings"
)

var tableName = "article_cate"

func CreateCate(req model.Cate) error {
	sql := fmt.Sprintf("insert into %s (name,creator_id) values (?,?)", tableName)
	_, err := config.DB.Exec(sql, req.Name, req.CreatorId)
	return err
}

func UpdateCate(req model.Cate) error {
	sql := fmt.Sprintf("update %s set name=?,updater_id=? where id=?", tableName)
	_, err := config.DB.Exec(sql, req.Name, req.UpdaterId, req.Id)
	return err
}

func DeleteCate(req model.Cate) error {
	// 检查article表中是否存在关联记录
	result, err := mysql.CheckTableExist("article", "cate_id", req.Id)
	if err != nil {
		return err
	}
	if result {
		return errors.New("该分类下有文章，请先删除文章")
	}

	// 执行删除操作
	sql := fmt.Sprintf("delete from %s where id=?", tableName)
	_, err = config.DB.Exec(sql, req.Id)
	return err
}

func GetCateDetail(req model.Cate) (model.CateDetail, error) {
	// 使用LEFT JOIN从user表中获取creator和updater的信息
	sql := fmt.Sprintf(`
        SELECT 
            t.id, 
            t.name, 
            t.create_time, 
            t.creator_id, 
            uc.username AS creator,  -- 从user表查询creator名称
            t.update_time, 
            t.updater_id, 
            COALESCE(uu.username, '') AS updater   -- 从user表查询updater名称
        FROM %s t
        LEFT JOIN user uc ON t.creator_id = uc.id
        LEFT JOIN user uu ON t.updater_id = uu.id
        WHERE t.id = ?
    `, tableName)

	var cateInfo model.CateDetail
	err := config.DB.QueryRow(sql, req.Id).Scan(
		&cateInfo.Id,
		&cateInfo.Name,
		&cateInfo.CreateTime,
		&cateInfo.CreatorId,
		&cateInfo.Creator,
		&cateInfo.UpdateTime,
		&cateInfo.UpdaterId,
		&cateInfo.Updater,
	)

	if err != nil {
		return cateInfo, err
	}

	return cateInfo, nil
}

func GetCateList(req model.Query) (model.List, error) {
	var (
		args       []interface{}
		conditions []string
		result     = model.List{}
	)

	// 构建查询条件（修改：使用cate表的name字段而非不存在的username）
	if req.Name != "" {
		conditions = append(conditions, "t.name LIKE ?")
		args = append(args, "%"+req.Name+"%")
	}

	// 查询总记录数（修改：添加表别名）
	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM %s t WHERE 1=1", tableName)
	if len(conditions) > 0 {
		countSQL += " AND " + strings.Join(conditions, " AND ")
	}
	if err := config.DB.QueryRow(countSQL, args...).Scan(&result.Total); err != nil {
		return result, err
	}

	// 处理分页参数
	pageNum := 1
	pageSize := 100
	if req.PageNum != nil && *req.PageNum > 0 {
		pageNum = *req.PageNum
	}
	if req.PageSize != nil && *req.PageSize > 0 {
		pageSize = *req.PageSize
	}
	offset := (pageNum - 1) * pageSize

	// 构建分页查询SQL（修改：使用与详情一致的JOIN查询）
	querySQL := fmt.Sprintf(`
        SELECT 
            t.id, 
            t.name, 
            t.create_time, 
            t.creator_id, 
            uc.username AS creator,
            t.update_time, 
            t.updater_id, 
            COALESCE(uu.username, '') AS updater
        FROM %s t
        LEFT JOIN user uc ON t.creator_id = uc.id
        LEFT JOIN user uu ON t.updater_id = uu.id
        WHERE 1=1
    `, tableName)

	if len(conditions) > 0 {
		querySQL += " AND " + strings.Join(conditions, " AND ")
	}

	// 添加排序条件
	querySQL += " ORDER BY t.id DESC"
	// 添加分页
	querySQL += " LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	// 执行分页查询
	rows, err := config.DB.Query(querySQL, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	var cates []model.CateDetail = []model.CateDetail{}
	for rows.Next() {
		var cate model.CateDetail
		if err := rows.Scan(
			&cate.Id,
			&cate.Name,
			&cate.CreateTime,
			&cate.CreatorId,
			&cate.Creator,
			&cate.UpdateTime,
			&cate.UpdaterId,
			&cate.Updater,
		); err != nil {
			return result, err
		}
		cates = append(cates, cate)
	}

	result.List = cates
	return result, nil
}
