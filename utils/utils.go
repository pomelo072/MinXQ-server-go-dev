package utils

import (
	"MinXQ-server-go-dev/models"
	"crypto/sha1"
	"encoding/hex"
	"io"
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
	result.Code = 200
	result.Msg = msgstring
	return result
}
