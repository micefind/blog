package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// handleValidationErrorsForArticle 处理数据验证错误
func handleValidationErrorsForArticle(c *gin.Context, err error) {
	validationErrors := err.(validator.ValidationErrors)
	for _, err := range validationErrors {
		field := err.Field()
		switch field {
		case "Title":
			utils.JSONResponse(c, http.StatusBadRequest, "文章名称不能为空，且长度在 1-50 位之间", nil)
		case "Status":
			utils.JSONResponse(c, http.StatusBadRequest, "文章状态设置错误", nil)
		}
		break
	}
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var requestData models.Article
	if err := c.ShouldBind(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	//	验证数据
	if err := utils.GetValidator().Struct(requestData); err != nil {
		handleValidationErrorsForArticle(c, err)
		return
	}
	//设置默认数据
	requestData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	views := 0
	// 设置创建人id
	if userID, ok := c.Get("userID"); ok {
		requestData.CreatorID = userID.(int)
	}
	//	数据库插入数据
	sql := "INSERT INTO article (title,cover_image, intro,keywords,content,creator_id,create_time,status,views) VALUES (?,?,?,?,?,?,?,?,?)"
	_, err := config.DB.Exec(sql, requestData.Title, requestData.CoverImage, requestData.Intro, requestData.Keywords, requestData.Content, requestData.CreatorID, requestData.CreateTime, requestData.Status, views)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库插入失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "添加成功", nil)
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var requestData models.Article
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	if err := utils.GetValidator().Struct(requestData); err != nil {
		handleValidationErrorsForArticle(c, err)
		return
	}
	sql := "UPDATE article SET title=?,cover_image=?,intro=?,keywords=?,content=?,status=? WHERE id=?"
	_, err := config.DB.Exec(sql, requestData.Title, requestData.CoverImage, requestData.Intro, requestData.Keywords, requestData.Content, requestData.Status, requestData.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库更新失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "更新成功", nil)
}

// GetArticleList 获取文章列表
func GetArticleList(c *gin.Context) {
	var requestData struct {
		PageNum  *int   `json:"pageNum"`
		PageSize *int   `json:"pageSize"`
		Keyword  string `json:"keyword"`
		Status   string `json:"status"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 查询列表数据
	query := "SELECT article.id,article.title,article.intro,article.cover_image,article.keywords,article.views,article.creator_id,article.create_time,article.status,user.username FROM article JOIN user ON article.creator_id = user.id WHERE 1=1"
	args := []interface{}{}
	if requestData.Keyword != "" {
		query += " AND (title LIKE ? OR intro LIKE ? OR keywords LIKE ?)"
		args = append(args, "%"+requestData.Keyword+"%")
		args = append(args, "%"+requestData.Keyword+"%")
		args = append(args, "%"+requestData.Keyword+"%")
	}
	if requestData.Status != "" {
		query += " AND article.status = ?"
		args = append(args, requestData.Status)
	}
	query += " ORDER BY id DESC" // 添加按照 id 降序排列
	if requestData.PageNum != nil && requestData.PageSize != nil {
		offset := (*requestData.PageNum - 1) * *requestData.PageSize
		query += " LIMIT ? OFFSET ?"
		args = append(args, *requestData.PageSize, offset)
	}
	rows, err := config.DB.Query(query, args...)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询列表失败: %v", err), nil)
		return
	}
	defer rows.Close()

	type articleItem struct {
		ID         int    `json:"id"`
		Title      string `json:"title" validate:"required,min=1,max=20"`
		Intro      string `json:"intro"`
		CoverImage string `json:"cover_image"`
		Keywords   string `json:"keywords"`
		Views      int    `json:"views"`
		CreatorID  int    `json:"creator_id"`
		CreateTime string `json:"create_time"`
		Status     string `json:"status" validate:"required,oneof=0 1 2"`
		Creator    string `json:"creator"`
	}

	var articleList []articleItem = []articleItem{}
	for rows.Next() {
		var article articleItem
		if err := rows.Scan(&article.ID, &article.Title, &article.Intro, &article.CoverImage, &article.Keywords, &article.Views, &article.CreatorID, &article.CreateTime, &article.Status, &article.Creator); err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据解析失败: %v", err), nil)
			return
		}
		articleList = append(articleList, article)
	}

	// 获取总记录数
	var total int
	args2 := []interface{}{}
	countQuery := "SELECT COUNT(*) FROM article WHERE 1=1"
	if requestData.Keyword != "" {
		countQuery += " AND (title LIKE ? OR intro LIKE ? OR keywords LIKE ?)"
		args2 = append(args2, "%"+requestData.Keyword+"%")
		args2 = append(args2, "%"+requestData.Keyword+"%")
		args2 = append(args2, "%"+requestData.Keyword+"%")
	}
	if requestData.Status != "" {
		countQuery += " AND article.status = ?"
		args2 = append(args2, requestData.Status)
	}
	err = config.DB.QueryRow(countQuery, args2...).Scan(&total)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询记录总数失败: %v", err), nil)
		return
	}

	// 返回查询结果
	utils.JSONResponse(c, http.StatusOK, "项目列表获取成功", gin.H{
		"total": total,
		"list":  articleList,
	})
}

// DeleteArticle 删除项目
func DeleteArticle(c *gin.Context) {
	var requestData models.Article
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	sql := "DELETE FROM article WHERE id=?"
	_, err := config.DB.Exec(sql, requestData.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库删除失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "删除文章成功", nil)
}

// GetArticleDetails 获取文章详情并使 views 字段自增一
func GetArticleDetails(c *gin.Context) {
	var requestData struct {
		ID int `json:"id"`
	}
	var article models.Article
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	// 开启事务
	tx, err := config.DB.Begin()
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("开启事务失败: %v", err), nil)
		return
	}
	// 查询文章信息
	sql := "SELECT * FROM article WHERE id=?"
	err = tx.QueryRow(sql, requestData.ID).Scan(&article.ID, &article.Title, &article.CoverImage, &article.Intro, &article.Keywords, &article.Content, &article.Views, &article.CreatorID, &article.CreateTime, &article.Status)
	if err != nil {
		// 回滚事务
		tx.Rollback()
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询失败: %v", err), nil)
		return
	}
	// 更新 views 字段
	updateSql := "UPDATE article SET views = views + 1 WHERE id =?"
	_, err = tx.Exec(updateSql, requestData.ID)
	if err != nil {
		// 回滚事务
		tx.Rollback()
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("更新 views 字段失败: %v", err), nil)
		return
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		// 回滚事务
		tx.Rollback()
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("提交事务失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取文章信息成功", article)
}
