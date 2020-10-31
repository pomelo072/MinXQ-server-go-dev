package myhandler

import (
	"MinXQ-server-go-dev/config"
	"MinXQ-server-go-dev/models"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	jsoniter "github.com/json-iterator/go"
)

// 调用Aliyun内容安全API
// 审核范围包括广告, 侮辱, 政治, 恐怖, 色情, 血腥, 灌水...
func UseShield(str string) string {
	client, err := sdk.NewClientWithAccessKey("cn-shanghai", config.Sysconfig.ALYAccessKeyID, config.Sysconfig.ALYAccessKeySecret)
	if err != nil {
		return "server error"
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "imageaudit.cn-shanghai.aliyuncs.com"
	request.Version = "2019-12-30"
	request.ApiName = "ScanText"
	request.QueryParams["RegionId"] = "cn-shanghai"
	request.QueryParams["Tasks.1.Content"] = str
	request.QueryParams["Labels.1.Label"] = "ad"
	request.QueryParams["Labels.2.Label"] = "spam"
	request.QueryParams["Labels.3.Label"] = "politics"
	request.QueryParams["Labels.4.Label"] = "abuse"
	request.QueryParams["Labels.5.Label"] = "terrorism"
	request.QueryParams["Labels.6.Label"] = "porn"
	request.QueryParams["Labels.7.Label"] = "flood"
	request.QueryParams["Labels.8.Label"] = "contraband"

	response, er := client.ProcessCommonRequest(request)
	if er != nil {
		return "shield request error"
	}
	r := response.GetHttpContentBytes()
	var result = new(models.Shield)
	err = jsoniter.Unmarshal(r, result)
	return result.Data.Elements[0].Results[0].Suggestion
}
