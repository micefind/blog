package tag

import (
	"fmt"
	"server/config"
	model "server/model/article"
	"strings"
)

var tableName = "article"

func CreateArticle(req model.Article) error {
	// 插入文章
	sql := fmt.Sprintf("insert into %s (cate_id, tag_id, title, intro, content, cover_img, status, creator_id) values (?,?,?,?,?,?,?,?)", tableName)
	_, err := config.DB.Exec(sql, req.CateId, req.TagId, req.Title, req.Intro, req.Content, req.CoverImg, req.Status, req.CreatorId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateArticle(req model.Article) error {
	sql := fmt.Sprintf("update %s set cate_id=?, tag_id=?, title=?, intro=?, content=?, cover_img=?, status=?, updater_id=? where id=?", tableName)
	_, err := config.DB.Exec(sql, req.CateId, req.TagId, req.Title, req.Intro, req.Content, req.CoverImg, req.Status, req.UpdaterId, req.Id)
	return err
}

func DeleteArticle(req model.Article) error {
	// 执行删除操作
	sql := fmt.Sprintf("update %s set is_deleted=1 where id=?", tableName)
	_, err := config.DB.Exec(sql, req.Id)
	return err
}

func RecoverArticle(req model.Article) error {
	// 执行恢复操作
	sql := fmt.Sprintf("update %s set is_deleted=0 where id=?", tableName)
	_, err := config.DB.Exec(sql, req.Id)
	return err
}

func GetArticleDetail(req model.Article) (model.ArticleDetail, error) {
	// 使用LEFT JOIN从user表中获取creator和updater的信息
	sql := fmt.Sprintf(`
        SELECT 
            t.id, 
            t.cate_id, 
            t.tag_id, 
            t.title, 
            t.intro, 
            t.content, 
			t.cover_img, 
			t.views,
            t.status, 
            t.creator_id, 
            t.updater_id, 
            t.create_time, 
            t.update_time,
            t.is_deleted,
            uc.username AS creator,  -- 从user表查询creator名称
            COALESCE(uu.username, '') AS updater,   -- 从user表查询updater名称
			cn.name AS cate_name,	-- 从cate表查询cate_name
			tn.name AS tag_name		-- 从tag表查询tag_name
        FROM %s t
        LEFT JOIN user uc ON t.creator_id = uc.id
        LEFT JOIN user uu ON t.updater_id = uu.id
		LEFT JOIN article_cate cn ON t.cate_id = cn.id
		LEFT JOIN article_tag tn ON t.tag_id = tn.id
        WHERE t.id = ?
    `, tableName)

	var tagInfo model.ArticleDetail
	err := config.DB.QueryRow(sql, req.Id).Scan(
		&tagInfo.Id,
		&tagInfo.CateId,
		&tagInfo.TagId,
		&tagInfo.Title,
		&tagInfo.Intro,
		&tagInfo.Content,
		&tagInfo.CoverImg,
		&tagInfo.Views,
		&tagInfo.Status,
		&tagInfo.CreatorId,
		&tagInfo.UpdaterId,
		&tagInfo.CreateTime,
		&tagInfo.UpdateTime,
		&tagInfo.IsDeleted,
		&tagInfo.Creator,
		&tagInfo.Updater,
		&tagInfo.CateName,
		&tagInfo.TagName,
	)

	if err != nil {
		return tagInfo, err
	}

	// 浏览量加1
	sql = fmt.Sprintf("UPDATE %s SET views = views + 1 WHERE id = ?", tableName)
	_, err = config.DB.Exec(sql, req.Id)
	if err != nil {
		return tagInfo, err
	}

	return tagInfo, nil
}

func GetArticleList(req model.Query) (model.List, error) {
	var (
		args       []interface{}
		conditions []string
		result     = model.List{}
	)

	// 构建查询条件 - 同时在title和intro中搜索keyword
	if req.Keyword != "" {
		// 使用SQL的OR操作符实现多字段搜索
		conditions = append(conditions, "(t.title LIKE ? OR t.intro LIKE ?)")
		// 注意：每个占位符?都需要对应一个参数
		args = append(args, "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.CateId != nil {
		conditions = append(conditions, "t.cate_id = ?")
		args = append(args, *req.CateId)
	}
	if req.TagId != nil {
		conditions = append(conditions, "t.tag_id = ?")
		args = append(args, *req.TagId)
	}
	if req.Status != nil {
		conditions = append(conditions, "t.status = ?")
		args = append(args, *req.Status)
	} else {
		// 默认显示已经发布的文章
		conditions = append(conditions, "t.status = 1")
	}
	if req.IsDeleted != nil {
		conditions = append(conditions, "t.is_deleted = ?")
		args = append(args, *req.IsDeleted)
	} else {
		// 默认不显示已删除的记录
		conditions = append(conditions, "t.is_deleted = 0")
	}

	// 查询总记录数
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
		if *req.PageSize > 100 {
			pageSize = 100
		} else {
			pageSize = *req.PageSize
		}
	}
	offset := (pageNum - 1) * pageSize

	// 构建分页查询SQL（与详情接口保持一致的JOIN查询）
	querySQL := fmt.Sprintf(`
        SELECT 
            t.id, 
            t.cate_id, 
            t.tag_id, 
            t.title, 
            t.intro, 
            t.cover_img, 
            t.views,
            t.status, 
            t.create_time, 
            t.creator_id, 
            t.update_time, 
            t.updater_id,
            t.is_deleted,
            uc.username AS creator,
            COALESCE(uu.username, '') AS updater,
            cn.name AS cate_name,
            tn.name AS tag_name
        FROM %s t
        LEFT JOIN user uc ON t.creator_id = uc.id
        LEFT JOIN user uu ON t.updater_id = uu.id
        LEFT JOIN article_cate cn ON t.cate_id = cn.id
        LEFT JOIN article_tag tn ON t.tag_id = tn.id
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

	var articles []model.ArticleDetail = []model.ArticleDetail{}
	for rows.Next() {
		var article model.ArticleDetail
		if err := rows.Scan(
			&article.Id,
			&article.CateId,
			&article.TagId,
			&article.Title,
			&article.Intro,
			&article.CoverImg,
			&article.Views,
			&article.Status,
			&article.CreateTime,
			&article.CreatorId,
			&article.UpdateTime,
			&article.UpdaterId,
			&article.IsDeleted,
			&article.Creator,
			&article.Updater,
			&article.CateName,
			&article.TagName,
		); err != nil {
			return result, err
		}
		articles = append(articles, article)
	}

	result.List = articles
	return result, nil
}
