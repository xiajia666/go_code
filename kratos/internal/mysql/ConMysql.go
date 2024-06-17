package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBTeacherInfo *gorm.DB
var err error

func init() {
	dsn := "root:985211@tcp(127.0.0.1)/lt_teacher_info?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:985211@tcp(118.25.191.214:3306)/Information"
	DBTeacherInfo, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DBTeacherInfo == nil && err != nil {
		panic("failed to connect mysql")
	}
}
