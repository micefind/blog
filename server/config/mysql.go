package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 引入 MySQL 驱动包，但不直接使用，仅用于初始化
)

var DB *sql.DB // 全局变量，保存数据库连接实例

// 与 MySQL 数据库的连接
func ConnectMysql() error {
	// 数据库连接信息，格式为：username:password@protocol(address)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Mysql.User, Mysql.Password, Mysql.Host, Mysql.Port, Mysql.DbName)

	// 使用 DSN 打开数据库连接，并返回一个 *sql.DB 实例
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 设置连接池参数
	DB.SetMaxOpenConns(100)  // 最大打开连接数（
	DB.SetMaxIdleConns(10)   // 最大空闲连接数
	DB.SetConnMaxLifetime(60 * time.Minute) // 连接的最大可复用时间

	// 尝试与数据库建立实际连接，检查数据库连接是否可用
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
