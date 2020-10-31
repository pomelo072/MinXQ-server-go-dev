package models

// aliyun响应数据
type Shield struct {
	Data      ShieldData
	RequestID string
}

type ShieldData struct {
	Elements []Element
}

type Element struct {
	TaskID  string
	Results []ALResult
}

type ALResult struct {
	Details    []Detail
	Suggestion string
	Label      string
	Rate       float32
}

type Detail struct {
	Contexts []Contx
	SubLabel string
}

type Contx struct {
	Context string
}
