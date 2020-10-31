package utils

import (
	"MinXQ-server-go-dev/models"
	"crypto/sha1"
	"encoding/hex"
	"gorm.io/gorm"
	"io"
	"strconv"
)

//GetSHAEncode 加密
func GetSHAEncode(str string) string {
	w := sha1.New()
	io.WriteString(w, str)            //将str写入到w中
	bw := w.Sum(nil)                  //w.Sum(nil)将w的hash转成[]byte格式
	shastr2 := hex.EncodeToString(bw) //将 bw 转成字符串
	return shastr2
}

// 打包数据和Code,Msg, 剩下就是前端的事了
func GetReturnData(dt interface{}, msgstring string) *models.Result {
	result := new(models.Result)
	result.Data = dt
	if msgstring == "SUCCESS" {
		result.Code = 200
	} else {
		result.Code = 400
	}
	result.Msg = msgstring
	return result
}

// 分页器
// 分页数据请求分页大小为30, 45, 60三种大小
func Paginate(pages string, pagesizes string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(pages)
		if page == 0 {
			page = 1
		}
		pagesize, _ := strconv.Atoi(pagesizes)
		switch {
		case pagesize <= 30:
			pagesize = 30
		case pagesize < 60:
			pagesize = 45
		case pagesize >= 60:
			pagesize = 60
		}

		offset := (page - 1) * pagesize
		return db.Offset(offset).Limit(pagesize)
	}
}
