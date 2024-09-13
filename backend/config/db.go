package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // 引入 MySQL 驱动包，但不直接使用，仅用于初始化
)

var DB *sql.DB // 全局变量，保存数据库连接实例

// ConnectDatabase 负责建立与 MySQL 数据库的连接
func ConnectDatabase() {
	var err error // 用于存储错误信息

	// 数据库连接信息，格式为：username:password@protocol(address)/dbname
	// 此处为本地 MySQL 数据库连接信息，用户名为 'root'，密码为 '123456'，数据库为 'blog_db'
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog_db"

	// 使用 DSN 打开数据库连接，并返回一个 *sql.DB 实例
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		// 如果打开连接时发生错误，记录错误日志并终止程序
		log.Fatalf("Error opening database: %s\n", err)
	}

	// 尝试与数据库建立实际连接，检查数据库连接是否可用
	err = DB.Ping()
	if err != nil {
		// 如果无法连接到数据库，记录错误日志并终止程序
		log.Fatalf("Error connecting to the database: %s\n", err)
	}
}
