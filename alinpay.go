package allinpayapi

import (
	"github.com/hhjpin/allinpayapi/core"
	"github.com/hhjpin/allinpayapi/service"
)

type AllInPay struct {
	base *core.Pay

	MemberService   *service.MemberService   //会员服务
	OrderService    *service.OrderService    //订单服务
	MerchantService *service.MerchantService //商家服务
}

func New(config *core.Config) (*AllInPay, error) {
	base, err := core.NewPay(config)
	if err != nil {
		return nil, err
	}
	pay := &AllInPay{
		base:            base,
		MemberService:   service.NewMemberService(base),
		OrderService:    service.NewOrderService(base),
		MerchantService: service.NewMerchantService(base),
	}
	return pay, nil
}

func (a *AllInPay) Encrypt(data string) (string, error) {
	return a.base.Encrypt(data)
}

func (a *AllInPay) Decrypt(text string) (string, error) {
	return a.base.Decrypt(text)
}
