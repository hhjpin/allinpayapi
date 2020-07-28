package service

import "github.com/hhjpin/allinpayapi/core"

type OrderService struct {
	base    *core.Pay
	service string
}

func NewOrderService(base *core.Pay) *OrderService {
	return &OrderService{
		base:    base,
		service: "OrderService",
	}
}

type DepositApplyReq struct {
	BizOrderNo          string                 `json:"bizOrderNo"`
	BizUserId           string                 `json:"bizUserId"`
	AccountSetNo        string                 `json:"accountSetNo"` //100001标准余额账户集 200126会员的账户余额是在托管账户集
	Amount              int                    `json:"amount"`       //单位分
	Fee                 int                    `json:"fee"`
	ValidateType        int                    `json:"validateType"` //0无验证 1短信验证码 2支付密码
	FrontUrl            string                 `json:"frontUrl"`
	BackUrl             string                 `json:"backUrl"`
	OrderExpireDatetime string                 `json:"orderExpireDatetime"` //控制订单可支付时间,订单最长时效为24小时
	PayMethod           map[string]interface{} `json:"payMethod"`
	GoodsName           string                 `json:"goodsName"`
	IndustryCode        string                 `json:"industryCode"` //1910其他行业
	IndustryName        string                 `json:"industryName"`
	Source              int                    `json:"source"`
	Summary             string                 `json:"summary"`
	ExtendInfo          string                 `json:"extendInfo"`
}

type DepositApplyRsp struct {
	PayStatus      string                 `json:"payStatus"`
	PayFailMessage string                 `json:"payFailMessage"`
	BizUserId      string                 `json:"bizUserId"`
	OrderNo        string                 `json:"orderNo"`
	BizOrderNo     string                 `json:"bizOrderNo"`
	PayCode        string                 `json:"payCode"`
	TradeNo        string                 `json:"tradeNo"`
	WeChatAPPInfo  map[string]interface{} `json:"weChatAPPInfo"`
	PayInfo        string                 `json:"payInfo"`
	ValidateType   int                    `json:"validateType"`
	ExtendInfo     string                 `json:"extendInfo"`
}

//充值申请
func (s *OrderService) DepositApply(req DepositApplyReq) (*DepositApplyRsp, error) {
	var res = &DepositApplyRsp{}
	err := s.base.CommonRequest(s.service, "depositApply", req, res)
	return res, err
}

type WithdrawApplyReq struct {
	BizOrderNo          string                 `json:"bizOrderNo"`
	BizUserId           string                 `json:"bizUserId"`
	AccountSetNo        string                 `json:"accountSetNo"`
	Amount              int                    `json:"amount"`
	Fee                 int                    `json:"fee"`
	ValidateType        int                    `json:"validateType"`
	BackUrl             string                 `json:"backUrl"`
	OrderExpireDatetime string                 `json:"orderExpireDatetime"`
	PayMethod           map[string]interface{} `json:"payMethod"`
	BankCardNo          string                 `json:"bankCardNo"`
	BankCardPro         int                    `json:"bankCardPro"` //0个人银行卡 1企业对公账户
	WithdrawType        string                 `json:"withdrawType"`
	IndustryCode        string                 `json:"industryCode"`
	IndustryName        string                 `json:"industryName"`
	Source              int                    `json:"source"`
	Summary             string                 `json:"summary"`
	ExtendInfo          string                 `json:"extendInfo"`
}

type WithdrawApplyRsp struct {
	PayStatus      string `json:"payStatus"`
	PayFailMessage string `json:"payFailMessage"`
	BizUserId      string `json:"bizUserId"`
	OrderNo        string `json:"orderNo"`
	BizOrderNo     string `json:"bizOrderNo"`
	ExtendInfo     string `json:"extendInfo"`
}

//提现申请
func (s *OrderService) WithdrawApply(req WithdrawApplyReq) (*WithdrawApplyRsp, error) {
	var res = &WithdrawApplyRsp{}
	var err error
	req.BankCardNo, err = s.base.Encrypt(req.BankCardNo)
	if err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "withdrawApply", req, res)
	return res, err
}

type ConsumeApplyReq struct {
	PayerId             string                 `json:"payerId"`
	RecieverId          string                 `json:"recieverId"`
	BizOrderNo          string                 `json:"bizOrderNo"`
	Amount              int                    `json:"amount"`
	Fee                 int                    `json:"fee"`
	ValidateType        int                    `json:"validateType"`
	SplitRule           []SplitRuleItem        `json:"splitRule"`
	FrontUrl            string                 `json:"frontUrl"`
	BackUrl             string                 `json:"backUrl"`
	OrderExpireDatetime string                 `json:"orderExpireDatetime"`
	PayMethod           map[string]interface{} `json:"payMethod"`
	GoodsType           int                    `json:"goodsType"`
	BizGoodsNo          string                 `json:"bizGoodsNo"`
	GoodsName           string                 `json:"goodsName"`
	GoodsDesc           string                 `json:"goodsDesc"`
	IndustryCode        string                 `json:"industryCode"`
	IndustryName        string                 `json:"industryName"`
	Source              int                    `json:"source"`
	Summary             string                 `json:"summary"`
	ExtendInfo          string                 `json:"extendInfo"`
}

type ConsumeApplyRsp struct {
	PayStatus      string                 `json:"payStatus"`
	PayFailMessage string                 `json:"payFailMessage"`
	BizUserId      string                 `json:"bizUserId"`
	OrderNo        string                 `json:"orderNo"`
	BizOrderNo     string                 `json:"bizOrderNo"`
	TradeNo        string                 `json:"tradeNo"`
	PayCode        string                 `json:"payCode"`
	ExtendInfo     string                 `json:"extendInfo"`
	WeChatAPPInfo  map[string]interface{} `json:"weChatAPPInfo"`
	PayInfo        string                 `json:"payInfo"`
	ValidateType   int                    `json:"validateType"`
}

//消费申请
func (s *OrderService) ConsumeApply(req ConsumeApplyReq) (*ConsumeApplyRsp, error) {
	var res = &ConsumeApplyRsp{}
	err := s.base.CommonRequest(s.service, "consumeApply", req, res)
	return res, err
}

type RecieverItem struct {
	BizUserId string `json:"bizUserId"`
	Amount    int    `json:"amount"`
}

type AgentCollectApplyReq struct {
	BizOrderNo          string                 `json:"bizOrderNo"`
	PayerId             string                 `json:"payerId"`
	RecieverList        []RecieverItem         `json:"recieverList"`
	GoodsType           int                    `json:"goodsType"`
	BizGoodsNo          string                 `json:"bizGoodsNo"`
	TradeCode           string                 `json:"tradeCode"`
	Amount              int                    `json:"amount"`
	Fee                 int                    `json:"fee"`
	ValidateType        int                    `json:"validateType"`
	FrontUrl            string                 `json:"frontUrl"`
	BackUrl             string                 `json:"backUrl"`
	OrderExpireDatetime string                 `json:"orderExpireDatetime"`
	PayMethod           map[string]interface{} `json:"payMethod"`
	GoodsName           string                 `json:"goodsName"`
	GoodsDesc           string                 `json:"goodsDesc"`
	IndustryCode        string                 `json:"industryCode"`
	IndustryName        string                 `json:"industryName"`
	Source              int                    `json:"source"`
	Summary             string                 `json:"summary"`
	ExtendInfo          string                 `json:"extendInfo"`
}

type AgentCollectApplyRsp struct {
	PayStatus      string                 `json:"payStatus"`
	PayFailMessage string                 `json:"payFailMessage"`
	BizUserId      string                 `json:"bizUserId"`
	OrderNo        string                 `json:"orderNo"`
	BizOrderNo     string                 `json:"bizOrderNo"`
	TradeNo        string                 `json:"tradeNo"`
	PayCode        string                 `json:"payCode"`
	ExtendInfo     string                 `json:"extendInfo"`
	WeChatAPPInfo  map[string]interface{} `json:"weChatAPPInfo"`
	PayInfo        string                 `json:"payInfo"`
	ValidateType   int                    `json:"validateType"`
}

//托管代收申请(标准版)
func (s *OrderService) AgentCollectApply(req AgentCollectApplyReq) (*AgentCollectApplyRsp, error) {
	var res = &AgentCollectApplyRsp{}
	err := s.base.CommonRequest(s.service, "agentCollectApply", req, res)
	return res, err
}

type CollectPayItem struct {
	BizOrderNo string `json:"bizOrderNo"`
	Amount     int    `json:"amount"`
}

type SplitRuleItem struct {
	BizUserId     string          `json:"bizUserId"`
	AccountSetNo  string          `json:"accountSetNo"`
	Amount        int             `json:"amount"`
	Fee           int             `json:"fee"`
	Remark        string          `json:"remark"`
	SplitRuleList []SplitRuleItem `json:"splitRuleList"`
}

type SignalAgentPayReq struct {
	BizOrderNo     string           `json:"bizOrderNo"`
	CollectPayList []CollectPayItem `json:"collectPayList"`
	BizUserId      string           `json:"bizUserId"`
	AccountSetNo   string           `json:"accountSetNo"`
	BackUrl        string           `json:"backUrl"`
	Amount         int              `json:"amount"`
	Fee            int              `json:"fee"`
	SplitRuleList  []SplitRuleItem  `json:"splitRuleList"`
	GoodsType      int              `json:"goodsType"`
	BizGoodsNo     string           `json:"bizGoodsNo"`
	TradeCode      string           `json:"tradeCode"`
	Summary        string           `json:"summary"`
	ExtendInfo     string           `json:"extendInfo"`
}

type SignalAgentPayRsp struct {
	PayStatus      string `json:"payStatus"`
	PayFailMessage string `json:"payFailMessage"`
	OrderNo        string `json:"orderNo"`
	BizOrderNo     string `json:"bizOrderNo"`
	PayWhereabouts int    `json:"payWhereabouts"`
	ExtendInfo     string `json:"extendInfo"`
}

//单笔托管代付(标准版)
func (s *OrderService) SignalAgentPay(req SignalAgentPayReq) (*SignalAgentPayRsp, error) {
	var res = &SignalAgentPayRsp{}
	err := s.base.CommonRequest(s.service, "signalAgentPay", req, res)
	return res, err
}

type BatchPayItem struct {
	BizOrderNo     string           `json:"bizOrderNo"`
	CollectPayList []CollectPayItem `json:"collectPayList"`
	BizUserId      string           `json:"bizUserId"`
	AccountSetNo   string           `json:"accountSetNo"`
	BackUrl        string           `json:"backUrl"`
	Amount         int              `json:"amount"`
	Fee            int              `json:"fee"`
	SplitRuleList  []SplitRuleItem  `json:"splitRuleList"`
	Summary        string           `json:"summary"`
	ExtendInfo     string           `json:"extendInfo"`
}

type BatchAgentPayReq struct {
	BizBatchNo   string         `json:"bizBatchNo"`
	BatchPayList []BatchPayItem `json:"batchPayList"`
	GoodsType    int            `json:"goodsType"`
	BizGoodsNo   string         `json:"bizGoodsNo"`
	TradeCode    string         `json:"tradeCode"`
}

type BatchAgentPayRsp struct {
	BizBatchNo string `json:"bizBatchNo"`
}

//批量托管代付(标准版)
func (s *OrderService) BatchAgentPay(req BatchAgentPayReq) (*BatchAgentPayRsp, error) {
	var res = &BatchAgentPayRsp{}
	err := s.base.CommonRequest(s.service, "batchAgentPay", req, res)
	return res, err
}

type PayReq struct {
	BizUserId        string `json:"bizUserId"`
	BizOrderNo       string `json:"bizOrderNo"`
	TradeNo          string `json:"tradeNo"`          //(后台+短信验证码确认)
	JumpUrl          string `json:"jumpUrl"`          //(前台+密码验证版)
	VerificationCode string `json:"verificationCode"` //(前台+短信验证码确认)
	ConsumerIp       string `json:"consumerIp"`
}

type PayRsp struct {
	PayStatus      string `json:"payStatus"`
	PayFailMessage string `json:"payFailMessage"`
	BizUserId      string `json:"bizUserId"`
	BizOrderNo     string `json:"bizOrderNo"`
}

//确认支付 (分前台和后台)
func (s *OrderService) Pay(req PayReq) (*PayRsp, error) {
	var res = &PayRsp{}
	err := s.base.CommonRequest(s.service, "pay", req, res)
	return res, err
}

func (s *OrderService) GetPayCodePage(host string, req PayReq) (string, error) {
	if host == "" {
		host = "https://fintech.allinpay.com/yungateway/frontTrans.do"
	}
	param, err := s.base.GetRequestParam(s.service, "pay", req)
	return host + param, err
}

func (s *OrderService) GetPayPasswordPage(host string, req PayReq) (string, error) {
	if host == "" {
		host = "https://fintech.allinpay.com/yungateway/pwd/payOrder.html"
	}
	param, err := s.base.GetRequestParam(s.service, "pay", req)
	return host + param, err
}

type GoodsDetail struct {
	GoodsId      string `json:"goods_id"`
	WxpayGoodsId string `json:"wxpay_goods_id"`
	GoodsName    string `json:"goods_name"`
	Quantity     int    `json:"quantity"`
	Price        int    `json:"price"`
}

type GoodsParams struct {
	CostPrice   string        `json:"cost_price"`
	ReceiptId   string        `json:"receipt_id"`
	GoodsDetail []GoodsDetail `json:"goods_detail"`
}

type EntryGoodsReq struct {
	BizUserId   string      `json:"bizUserId"`
	GoodsType   int         `json:"goodsType"`
	BizGoodsNo  string      `json:"bizGoodsNo"`
	GoodsName   string      `json:"goodsName"`
	GoodsDetail string      `json:"goodsDetail"`
	GoodsParams interface{} `json:"goodsParams"` //具体使用 GoodsParams 类型
	ShowUrl     string      `json:"showUrl"`
}

type EntryGoodsRsp struct {
	GoodsNo    string `json:"goodsNo"`
	BizGoodsNo string `json:"bizGoodsNo"`
}

//商品录入
func (s *OrderService) EntryGoods(req EntryGoodsReq) (*EntryGoodsRsp, error) {
	var res = &EntryGoodsRsp{}
	err := s.base.CommonRequest(s.service, "entryGoods", req, res)
	return res, err
}

type FreezeMoneyReq struct {
	BizUserId    string `json:"bizUserId"`
	BizFreezenNo string `json:"bizFreezenNo"`
	AccountSetNo string `json:"accountSetNo"`
	Amount       int    `json:"amount"`
}

type FreezeMoneyRsp struct {
	BizFreezenNo string `json:"bizFreezenNo"`
	Amount       int    `json:"amount"`
}

//冻结金额
func (s *OrderService) FreezeMoney(req FreezeMoneyReq) (*FreezeMoneyRsp, error) {
	var res = &FreezeMoneyRsp{}
	err := s.base.CommonRequest(s.service, "freezeMoney", req, res)
	return res, err
}

type UnfreezeMoneyReq struct {
	BizUserId    string `json:"bizUserId"`
	BizFreezenNo string `json:"bizFreezenNo"`
	AccountSetNo string `json:"accountSetNo"`
	Amount       int    `json:"amount"`
}

type UnfreezeMoneyRsp struct {
	BizFreezenNo string `json:"bizFreezenNo"`
	Amount       int    `json:"amount"`
}

//解冻金额
func (s *OrderService) UnfreezeMoney(req UnfreezeMoneyReq) (*UnfreezeMoneyRsp, error) {
	var res = &UnfreezeMoneyRsp{}
	err := s.base.CommonRequest(s.service, "unfreezeMoney", req, res)
	return res, err
}

//todo 订单结果通知

type RefundItem struct {
	AccountSetNo string `json:"accountSetNo"`
	BizUserId    string `json:"bizUserId"`
	Amount       int    `json:"amount"`
}

type RefundReq struct {
	BizOrderNo    string       `json:"bizOrderNo"`
	OriBizOrderNo string       `json:"oriBizOrderNo"`
	BizUserId     string       `json:"bizUserId"`
	RefundType    string       `json:"refundType"`
	RefundList    []RefundItem `json:"refundList"`
	BackUrl       string       `json:"backUrl"`
	Amount        int          `json:"amount"`
	CouponAmount  int          `json:"couponAmount"`
	FeeAmount     int          `json:"feeAmount"`
	ExtendInfo    string       `json:"extendInfo"`
}

type RefundRsp struct {
	PayStatus      string `json:"payStatus"`
	PayFailMessage string `json:"payFailMessage"`
	OrderNo        string `json:"orderNo"`
	BizOrderNo     string `json:"bizOrderNo"`
	Amount         int    `json:"amount"`
	CouponAmount   int    `json:"couponAmount"`
	FeeAmount      int    `json:"feeAmount"`
	ExtendInfo     string `json:"extendInfo"`
}

//退款申请
func (s *OrderService) Refund(req RefundReq) (*RefundRsp, error) {
	var res = &RefundRsp{}
	err := s.base.CommonRequest(s.service, "refund", req, res)
	return res, err
}

type ApplicationTransferReq struct {
	BizTransferNo      string `json:"bizTransferNo"`
	SourceAccountSetNo string `json:"sourceAccountSetNo"`
	TargetBizUserId    string `json:"targetBizUserId"`
	TargetAccountSetNo string `json:"targetAccountSetNo"`
	Amount             int    `json:"amount"`
	ExtendInfo         string `json:"extendInfo"`
}

type ApplicationTransferRsp struct {
	TransferNo    string `json:"transferNo"`
	BizTransferNo string `json:"bizTransferNo"`
	Amount        int    `json:"amount"`
	ExtendInfo    string `json:"extendInfo"`
}

//平台转账
func (s *OrderService) ApplicationTransfer(req ApplicationTransferReq) (*ApplicationTransferRsp, error) {
	var res = &ApplicationTransferRsp{}
	err := s.base.CommonRequest(s.service, "applicationTransfer", req, res)
	return res, err
}

type QueryBalanceReq struct {
	BizUserId    string `json:"bizUserId"`
	AccountSetNo string `json:"accountSetNo"`
}

type QueryBalanceRsp struct {
	AllAmount     int `json:"allAmount"`
	FreezenAmount int `json:"freezenAmount"`
}

//查询余额
func (s *OrderService) QueryBalance(req QueryBalanceReq) (*QueryBalanceRsp, error) {
	var res = &QueryBalanceRsp{}
	err := s.base.CommonRequest(s.service, "queryBalance", req, res)
	return res, err
}

type GetOrderDetailReq struct {
	BizOrderNo string `json:"bizOrderNo"`
}

type GetOrderDetailRsp struct {
	OrderNo                string `json:"orderNo"`
	BizOrderNo             string `json:"bizOrderNo"`
	OriOrderNo             string `json:"oriOrderNo"`
	OriBizOrderNo          string `json:"oriBizOrderNo"`
	OrderStatus            int    `json:"orderStatus"`
	ErrorMessage           string `json:"errorMessage"`
	Amount                 int    `json:"amount"`
	PayDatetime            string `json:"payDatetime"`
	BuyerBizUserId         string `json:"buyerBizUserId"`
	RefundWhereabouts      int    `json:"refundWhereabouts"`
	PayWhereabouts         int    `json:"payWhereabouts"`
	Acct                   string `json:"acct"`
	Accttype               string `json:"accttype"`
	Termno                 string `json:"termno"`
	Cusid                  string `json:"cusid"`
	PayInterfaceOutTradeNo string `json:"payInterfaceOutTradeNo"`
	Termrefnum             string `json:"termrefnum"`
	ChannelFee             string `json:"channelFee"`
	ChannelPaytime         string `json:"channelPaytime"`
	PayInterfacetrxcode    string `json:"payInterfacetrxcode"`
	Traceno                string `json:"traceno"`
	ExtendInfo             string `json:"extendInfo"`
}

//查询订单状态
func (s *OrderService) GetOrderDetail(req GetOrderDetailReq) (*GetOrderDetailRsp, error) {
	var res = &GetOrderDetailRsp{}
	err := s.base.CommonRequest(s.service, "getOrderDetail", req, res)
	return res, err
}
