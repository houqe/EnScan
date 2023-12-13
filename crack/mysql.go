package crack

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func mySql(cancel context.CancelFunc, ip, user, passwd string, port, timeout int) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds", user, passwd, ip, port, "mysql", timeout)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //设置日志级别为silent
	})
	if err == nil {
		end(ip, user, passwd, port, "MySql")
		cancel()
	}

}
