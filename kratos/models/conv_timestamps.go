package models

import (
	"reflect"
	"time"
)

func ConvTimeStamps(object interface{}) {
	val := reflect.ValueOf(object)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		panic("object must be a pointer to a slice")
	}
	slice := val.Elem()
	if slice.IsNil() {
		return
	}
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i)
		if elem.Kind() == reflect.Struct {
			field := elem.FieldByName("DeletedAt")
			if field.IsValid() && field.CanSet() {
				// 将字段的值转换为 time.Time 类型，然后格式化
				tDeletedAt := TimeParse(field.String()).Format("2006-01-02 15:04:05")
				field.SetString(tDeletedAt)
			}
			field = elem.FieldByName("CreatedAt")
			if field.IsValid() && field.CanSet() {
				// 将字段的值转换为 time.Time 类型，然后格式化
				tCreatedAt := TimeParse(field.String()).Format("2006-01-02 15:04:05")
				field.SetString(tCreatedAt)
			}
			field = elem.FieldByName("UpdatedAt")
			if field.IsValid() && field.CanSet() {
				// 将字段的值转换为 time.Time 类型，然后格式化
				tUpdatedAt := TimeParse(field.String()).Format("2006-01-02 15:04:05")
				field.SetString(tUpdatedAt)
			}
		}
	}
}

func TimeParse(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic(err)
	}
	return t
}
