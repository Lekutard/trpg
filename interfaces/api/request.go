package api

import (
	"github.com/gin-gonic/gin"
)

//GetQuery 指定された名称のクエリの値を取得
func GetQuery(c *gin.Context, name string) string {
	q := c.Request.URL.Query()
	if len(q[name]) > 0 {
		return q[name][0]
	}
	return ""
}
