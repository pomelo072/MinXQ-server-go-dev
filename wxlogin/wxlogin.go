package wxlogin

import (
	"MinXQ-server-go-dev/config"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// 微信接口响应数据
type WxloginRsp struct {
	OpenID     string `json:"OpenID"`
	SessionKey string `json:"SessionKey"`
	UnionID    string `json:"UnionID"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func WXLogin(code string) (*WxloginRsp, error) {
	// url组合完整请求
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url = fmt.Sprintf(url, config.Sysconfig.AppID, config.Sysconfig.SecretKeyWx, code)
	// 向微信发起请求
	rsp, err := http.Get(url)
	// 最好没错, 用完了就关了吧
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	// json解包塞进结构体, 最好也没错
	wxrsp := WxloginRsp{}
	decoder := jsoniter.NewDecoder(rsp.Body)
	if err := decoder.Decode(&wxrsp); err != nil {
		return nil, err
	}
	// 最好微信接口调用成功 (别错
	if wxrsp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxrsp.ErrCode, wxrsp.ErrMsg))
	}
	return &wxrsp, nil
}
