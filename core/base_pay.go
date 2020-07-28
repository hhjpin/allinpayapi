package core

import (
	"encoding/json"
	"github.com/hhjpin/allinpay"
	"github.com/hhjpin/allinpayapi/utils"
	"github.com/hhjpin/goutils/logger"
)

type Pay struct {
	pay    *allinpay.Pay
	config *Config
}

type Config struct {
	RequestUrl  string //请求地址
	PrivateData []byte //私钥数据
	PublicData  []byte //公钥数据
	Sysid       string //系统id
	IsDebug     bool   //是否打印调试信息

	//所有请求响应的回调函数
	RspCallback func(service, method string, req map[string]interface{}, rsp interface{}, isFailed bool, failReason string)
}

func (c *Config) CallRsp(service, method string, req map[string]interface{}, rsp interface{}, isFailed bool, failReason string) {
	if c.RspCallback != nil {
		c.RspCallback(service, method, req, rsp, isFailed, failReason)
	}
}

func NewPay(config *Config) (*Pay, error) {
	pay, err := allinpay.New(&allinpay.Config{
		RequestUrl:  config.RequestUrl,
		PrivateData: config.PrivateData,
		PublicData:  config.PublicData,
		Sysid:       config.Sysid,
		IsDebug:     config.IsDebug,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	obj := &Pay{
		pay:    pay,
		config: config,
	}
	return obj, nil
}

func (p *Pay) CommonRequest(service, method string, req interface{}, rsp interface{}) error {
	param := map[string]interface{}{}
	if req != nil {
		data, err := utils.ConvertStruct2Map(req)
		if err != nil {
			logger.Error(err)
			return err
		}
		param = data[0]
	}
	resp, err := p.pay.RequestAndCheckStatus(service, method, param)
	if err != nil {
		p.config.CallRsp(service, method, param, nil, true, err.Error())
		logger.Errorf("allinpay error: %s:%s, resp: %#v", service, method, resp)
		return err
	}
	if err := json.Unmarshal([]byte(resp.SignedValue), rsp); err != nil {
		p.config.CallRsp(service, method, param, nil, true, err.Error())
		logger.Errorf("allinpay json error: %s:%s, err: %s", service, method, err.Error())
		return err
	}
	p.config.CallRsp(service, method, param, rsp, false, "")
	return err
}

func (p *Pay) GetRequestParam(service, method string, req interface{}) (string, error) {
	data, err := utils.ConvertStruct2Map(req)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	param, err := p.pay.GetRequestParam(service, method, data[0])
	if err != nil {
		return "", err
	}
	return param, nil
}

func (p *Pay) Encrypt(data string) (string, error) {
	if data == "" {
		return "", nil
	}
	text, err := p.pay.GetConfig().Encrypt([]byte(data))
	return string(text), err
}

func (p *Pay) Decrypt(text string) (string, error) {
	if text == "" {
		return "", nil
	}
	data, err := p.pay.GetConfig().Decrypt([]byte(text))
	return string(data), err
}
