package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"github.com/hatjwe/soar-client/log"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// BlockIP 结构体
type BlockIP struct {
	AttackIP    string `json:"attack_ip"`
	AttackLevel string `json:"attack_level"`
}
type Soar struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	//
	Methon string
	// AuthInfo is for authentication
	Body    string
	BlockIp []BlockIP
	Header  map[string]string
}

// New creates a new harbor API HTTP client.
func New() *Soar {
	cli := new(Soar)
	return cli
}

func (soar *Soar) SetBody(BlockedIps BlockIP) {

	soar.BlockIp = append(soar.BlockIp, BlockedIps)
}
func (soar *Soar) ConventJson() (string, error) {
	if soar.BlockIp == nil {
		log.Logger.Error("未设置封禁ip数据转换失败")
	}
	jsonData, err := json.Marshal(soar.BlockIp)
	if err != nil {
		log.Logger.Error("转换为 JSON 失败:", zap.Error(err))

		return "", err
	}
	return string(jsonData), err

}

func (soar *Soar) HeaderSet(key, value string) {

	soar.Header[key] = value

}
func (soar *Soar) SentHttps(method, url string, body *strings.Reader) (string, error) {

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	if soar.Header != nil {
		for key, value := range soar.Header {
			request.Header.Set(key, value)
		}

	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	bodystr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	log.Logger.Info("响应包信息", zap.String("body", string(bodystr)), zap.String("url", url))
	return string(bodystr), nil
}
