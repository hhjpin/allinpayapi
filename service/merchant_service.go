package service

import "github.com/hhjpin/allinpayapi/core"

type MerchantService struct {
	base    *core.Pay
	service string
}

func NewMerchantService(base *core.Pay) *MerchantService {
	return &MerchantService{
		base:    base,
		service: "MerchantService",
	}
}

type GetCheckAccountFileReq struct {
	Date     string `json:"date"`
	FileType int    `json:"fileType"` //1明细 2汇总
}

type GetCheckAccountFileRsp struct {
	Url string `json:"url"`
}

//查询订单状态
func (s *MerchantService) GetCheckAccountFile(req GetCheckAccountFileReq) (*GetCheckAccountFileRsp, error) {
	var res = &GetCheckAccountFileRsp{}
	err := s.base.CommonRequest(s.service, "getCheckAccountFile", req, res)
	return res, err
}

type QueryReserveFundBalanceRsp struct {
	AccountNo   string `json:"account_no"`
	AccountName string `json:"account_name"`
	Balance     int    `json:"balance"`
	DefClr      int    `json:"def_clr"`
}

//查询订单状态
func (s *MerchantService) QueryReserveFundBalance() (*QueryReserveFundBalanceRsp, error) {
	var res = &QueryReserveFundBalanceRsp{}
	err := s.base.CommonRequest(s.service, "queryReserveFundBalance", nil, res)
	return res, err
}

type QueryMerchantBalanceReq struct {
	AccountSetNo string `json:"accountSetNo"`
}

type QueryMerchantBalanceRsp struct {
	AllAmount    int `json:"allAmount"`
	FreezeAmount int `json:"freezeAmount"`
}

//查询订单状态
func (s *MerchantService) QueryMerchantBalance(req QueryMerchantBalanceReq) (*QueryMerchantBalanceRsp, error) {
	var res = &QueryMerchantBalanceRsp{}
	err := s.base.CommonRequest(s.service, "queryMerchantBalance", req, res)
	return res, err
}

type QueryBankBalanceReq struct {
	AcctOrgType int    `json:"acctOrgType"`
	AcctNo      string `json:"acctNo"`
	AcctName    string `json:"acctName"`
}

type QueryBankBalanceRsp struct {
	AcctOrgType int    `json:"acctOrgType"`
	AcctNo      string `json:"acctNo"`
	AcctName    string `json:"acctName"`
	Balance     int    `json:"balance"`
}

//查询订单状态
func (s *MerchantService) QueryBankBalance(req QueryBankBalanceReq) (*QueryBankBalanceRsp, error) {
	var res = &QueryBankBalanceRsp{}
	err := s.base.CommonRequest(s.service, "queryBankBalance", req, res)
	return res, err
}
