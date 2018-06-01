package controller

import (
	"time"
	"html/template"
)

// 格式化时间
func DateFormat(date time.Time) string {
	return date.Format("2006-01-02")
}

// 格式化html(输出原生html)
func Unescaped(x string) interface{} {
	return template.HTML(x)
}
