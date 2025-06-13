package mysql

import (
	"fmt"
	"server/config"
)

// 检查指定表的字段值是否唯一
// * tableName 表名
// * fieldName 字段名
// * fieldValue 字段值
func CheckTableExist(tableName string, fieldName string, fieldValue interface{}) (bool, error) {
	// 构建查询语句
	sql := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = ? LIMIT 1)", tableName, fieldName)
	// 执行查询
	var exists bool
	if err := config.DB.QueryRow(sql, fieldValue).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
