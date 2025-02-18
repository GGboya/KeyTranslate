package apiservice

import (
	"encoding/json"
	"fmt"
	"translate/model"
)

type DeepSeekService struct {
}

func NewDeepSeekService() *DeepSeekService {
	return &DeepSeekService{}
}

func (s *DeepSeekService) SendDeepSeekRequest(data string) (*model.DeepSeekResponse, error) {
	fmt.Println("把用户需要翻译内容发送到大模型")
	url := "https://api.siliconflow.cn/v1/chat/completions"

	// 构造请求体
	requestBody := map[string]interface{}{
		"model": "THUDM/glm-4-9b-chat",
		"messages": []map[string]string{
			{"role": "system", "content": "请你严格遵循以下规则：将用户的内容翻译为英文然后返回，不要输出任何多余的内容"},
			{"role": "user", "content": data},
		},
		"stream": false,
	}

	respBody, err := doRequest("POST", url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to DeepSeek: %v", err)
	}

	var deepSeekResponse model.DeepSeekResponse
	if err := json.Unmarshal(respBody, &deepSeekResponse); err != nil {
		return nil, fmt.Errorf("error parsing response JSON: %w", err)
	}

	return &deepSeekResponse, nil
}
