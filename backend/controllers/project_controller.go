package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// handleValidationErrorsForProject 处理数据验证错误
func handleValidationErrorsForProject(c *gin.Context, err error) {
	validationErrors := err.(validator.ValidationErrors)
	for _, err := range validationErrors {
		field := err.Field()
		switch field {
		case "ProjectName":
			utils.JSONResponse(c, http.StatusBadRequest, "项目名称不能为空，且长度在 1-20 位之间", nil)
		case "Description":
			utils.JSONResponse(c, http.StatusBadRequest, "项目描述在100字以下", nil)
		}
		break
	}
}

// AddProject 添加项目
func AddProject(c *gin.Context) {
	var requestData models.Project
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	//	验证数据
	if err := utils.GetValidator().Struct(requestData); err != nil {
		handleValidationErrorsForProject(c, err)
		return
	}
	//	数据库插入数据
	sql := "INSERT INTO project (project_name, description,logo) VALUES (?, ?,?)"
	_, err := config.DB.Exec(sql, requestData.ProjectName, requestData.Description, requestData.Logo)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库插入失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "添加成功", nil)
}

// EditProject 编辑项目信息
func EditProject(c *gin.Context) {
	var requestData models.Project
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	if err := utils.GetValidator().Struct(requestData); err != nil {
		handleValidationErrorsForProject(c, err)
		return
	}
	sql := "UPDATE project SET project_name=?,description=?,logo=? WHERE id=?"
	_, err := config.DB.Exec(sql, requestData.ProjectName, requestData.Description, requestData.Logo, requestData.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库更新失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "更新成功", nil)
}

// GetProjectList 获取项目列表
func GetProjectList(c *gin.Context) {
	var requestData struct {
		PageNum     *int   `json:"pageNum"`
		PageSize    *int   `json:"pageSize"`
		ProjectName string `json:"project_name"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 查询列表数据
	query := "SELECT * FROM project WHERE 1=1"
	args := []interface{}{}
	if requestData.ProjectName != "" {
		query += " AND project_name LIKE ?"
		args = append(args, "%"+requestData.ProjectName+"%")
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

	var projectList []models.Project = []models.Project{}
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.ProjectName, &project.Description, &project.Logo); err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据解析失败: %v", err), nil)
			return
		}
		projectList = append(projectList, project)
	}

	// 获取总记录数
	var total int
	args2 := []interface{}{}
	countQuery := "SELECT COUNT(*) FROM project WHERE 1=1"
	if requestData.ProjectName != "" {
		countQuery += " AND project_name LIKE ?"
		args2 = append(args2, "%"+requestData.ProjectName+"%")
	}
	err = config.DB.QueryRow(countQuery, args2...).Scan(&total)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询记录总数失败: %v", err), nil)
		return
	}

	// 返回查询结果
	utils.JSONResponse(c, http.StatusOK, "项目列表获取成功", gin.H{
		"total": total,
		"list":  projectList,
	})
}

// DeleteProject 删除项目
func DeleteProject(c *gin.Context) {
	var requestData models.Project
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	sql := "DELETE FROM project WHERE id=?"
	_, err := config.DB.Exec(sql, requestData.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库删除失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "删除项目成功", nil)
}

// GetProjectDetails 获取项目详情
func GetProjectDetails(c *gin.Context) {
	var requestData struct {
		ID int `json:"id"`
	}
	var project models.Project
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	sql := "SELECT * FROM project WHERE id=?"
	err := config.DB.QueryRow(sql, requestData.ID).Scan(&project.ID, &project.ProjectName, &project.Description, &project.Logo)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取项目信息成功", project)
}
