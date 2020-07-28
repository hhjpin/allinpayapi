package allinpayapi

import (
	"encoding/json"
	"fmt"
	"github.com/hhjpin/allinpayapi/core"
	"github.com/hhjpin/allinpayapi/pem"
	"github.com/hhjpin/allinpayapi/service"
	"testing"
	"time"
)

var (
	testPay  *AllInPay
	backUrl  = "https://xxx.com/api/v1/backurl"  //填写一个可以访问的url
	frontUrl = "https://xxx.com/api/v1/fronturl" //填写一个可以访问的url
)

func init() {
	var err error
	testPay, err = New(&core.Config{
		RequestUrl:  "http://116.228.64.55:6900/service/soa",
		PrivateData: pem.PrivateData,
		PublicData:  pem.PublicData,
		Sysid:       "1902271423530473681",
		IsDebug:     true,
		RspCallback: func(service, method string, req map[string]interface{}, rsp interface{}, isFailed bool, failReason string) {
			reqStr, _ := json.Marshal(req)
			rspStr, _ := json.Marshal(rsp)
			fmt.Printf("+++RspCallback: %s:%s, %s, %s, %t, %s\n", service, method, reqStr, rspStr, isFailed, failReason)
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestMemberServiceCreateMember(t *testing.T) {
	res, err := testPay.MemberService.CreateMember(service.CreateMemberReq{
		BizUserId:   "test_user_71",
		MemberType:  service.MemberTypeCompany,
		Source:      service.SourceMobile,
		ExtendParam: map[string]interface{}{},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceSendVerificationCode(t *testing.T) {
	res, err := testPay.MemberService.SendVerificationCode(service.SendVerificationCodeReq{
		BizUserId:            "test_user_68",
		Phone:                "13760648472",
		VerificationCodeType: 6,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceBindPhone(t *testing.T) {
	res, err := testPay.MemberService.BindPhone(service.BindPhoneReq{
		BizUserId:        "test_user_68",
		Phone:            "13760648472",
		VerificationCode: "11111",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceUnbindPhone(t *testing.T) {
	res, err := testPay.MemberService.UnbindPhone(service.UnbindPhoneReq{
		BizUserId:        "test_user_68",
		Phone:            "13760648472",
		VerificationCode: "888888",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceSignContractPage(t *testing.T) {
	url, err := testPay.MemberService.SignContractPage(
		"http://116.228.64.55:6900/yungateway/member/signContract.html",
		service.SignContractReq{
			BizUserId: "test_user_68",
			JumpUrl:   frontUrl,
			BackUrl:   backUrl,
			Source:    service.SourceMobile,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestMemberServiceSignContract(t *testing.T) {
	res, err := testPay.MemberService.SignContract(service.SignContractReq{
		BizUserId: "test_user_99",
		JumpUrl:   frontUrl,
		BackUrl:   backUrl,
		Source:    service.SourceMobile,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceSetRealName(t *testing.T) {
	//要设置真实信息，后面绑卡和支付要校验真实信息
	res, err := testPay.MemberService.SetRealName(service.SetRealNameReq{
		BizUserId:    "test_user_68",
		IsAuth:       true,
		Name:         "张三",
		IdentityType: service.IdentityTypeIdCard,
		IdentityNo:   "xxxx",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceSetCompanyInfo(t *testing.T) {
	res, err := testPay.MemberService.SetCompanyInfo(service.SetCompanyInfoReq{
		BizUserId: "test_user_71",
		IsAuth:    false,
		BackUrl:   backUrl,
		CompanyBasicInfo: service.CompanyBasicInfo{
			CompanyName:      "测试公司",
			CompanyAddress:   "测试公司地址",
			AuthType:         2,
			UniCredit:        "https://www.baidu.com",
			BusinessLicense:  "",
			OrganizationCode: "",
			TaxRegister:      "",
			ExpLicense:       "",
			Telephone:        "13711111111",
			LegalName:        "小林",
			IdentityType:     service.IdentityTypeIdCard,
			LegalIds:         "111111111111111111",
			LegalPhone:       "13711111111",
			AccountNo:        "22222222222222222",
			ParentBankName:   "工商银行",
			BankCityNo:       "",
			BankName:         "中国工商银行股份有限公司北京樱桃园支行",
			UnionBank:        "666666666666",
			Province:         "",
			City:             "",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceGetMemberInfoForIndividual(t *testing.T) {
	res, err := testPay.MemberService.GetMemberInfoForIndividual(service.GetMemberInfoReq{
		BizUserId: "test_user_68",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceGetMemberInfoForCompany(t *testing.T) {
	res, err := testPay.MemberService.GetMemberInfoForCompany(service.GetMemberInfoReq{
		BizUserId: "test_user_99",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceGetBankCardBin(t *testing.T) {
	res, err := testPay.MemberService.GetBankCardBin(service.GetBankCardBinReq{
		CardNo: "xxxx",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceApplyBindBankCard(t *testing.T) {
	//要绑真实银行卡和信息
	res, err := testPay.MemberService.ApplyBindBankCard(service.ApplyBindBankCardReq{
		BizUserId:    "test_user_68",
		CardNo:       "xx",
		Phone:        "xx",
		Name:         "xx",
		CardCheck:    7,
		IdentityType: service.IdentityTypeIdCard,
		IdentityNo:   "xx",
		Validate:     "",
		Cvv2:         "",
		IsSafeCard:   false,
		UnionBank:    "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceBindBankCard(t *testing.T) {
	res, err := testPay.MemberService.BindBankCard(service.BindBankCardReq{
		BizUserId:        "test_user_68",
		TranceNum:        "101362201437",
		TransDate:        "",
		Phone:            "xx",
		Validate:         "",
		Cvv2:             "",
		VerificationCode: "530405",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceUnbindBankCard(t *testing.T) {
	res, err := testPay.MemberService.UnbindBankCard(service.UnbindBankCardReq{
		BizUserId: "test_user_68",
		CardNo:    "xxxx",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceLockMember(t *testing.T) {
	res, err := testPay.MemberService.LockMember(service.LockMemberReq{
		BizUserId: "test_user_66",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceUnlockMember(t *testing.T) {
	res, err := testPay.MemberService.UnlockMember(service.UnlockMemberReq{
		BizUserId: "test_user_66",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMemberServiceApplyBindAcct(t *testing.T) {
	res, err := testPay.MemberService.ApplyBindAcct(service.ApplyBindAcctReq{
		BizUserId:     "test_user_68",
		OperationType: "set",
		AcctType:      "weChatMiniProgram",
		Acct:          "xxx",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceDepositApply(t *testing.T) {
	bank, _ := testPay.Encrypt("xxxx") //要写真实密码
	res, err := testPay.OrderService.DepositApply(service.DepositApplyReq{
		BizOrderNo:          "test_user_68_order15",
		BizUserId:           "test_user_68",
		AccountSetNo:        service.AccountSetNoUserTest,
		Amount:              1,
		Fee:                 0,
		ValidateType:        service.ValidateTypeSms,
		FrontUrl:            frontUrl,
		BackUrl:             backUrl,
		OrderExpireDatetime: time.Now().Add(time.Hour * 12).Format("2006-01-02 15:04:05"),
		PayMethod: map[string]interface{}{
			/*"GATEWAY_VSP": map[string]interface{}{
				"amount":  1,
				"gateid":  "0103", //农业银行
				"paytype": "B2C",
			},*/
			"QUICKPAY_VSP": map[string]interface{}{
				"amount":     1,
				"bankCardNo": bank,
			},
		},
		GoodsName:    "消费申请商品",
		IndustryCode: "1910",
		IndustryName: "其他",
		Source:       service.SourceMobile,
		Summary:      "测试测试",
		ExtendInfo:   "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceWithdrawApply(t *testing.T) {
	res, err := testPay.OrderService.WithdrawApply(service.WithdrawApplyReq{
		BizOrderNo:          "test_user_68_order2",
		BizUserId:           "test_user_68",
		AccountSetNo:        service.AccountSetNoUserTest,
		Amount:              1,
		Fee:                 0,
		ValidateType:        service.ValidateTypeNone,
		BackUrl:             backUrl,
		OrderExpireDatetime: time.Now().Add(time.Hour).Format("2006-01-02 15:04:05"),
		PayMethod:           map[string]interface{}{},
		BankCardNo:          "xxxx",
		BankCardPro:         0,
		WithdrawType:        "D0",
		IndustryCode:        "1910",
		IndustryName:        "其他",
		Source:              service.SourceMobile,
		Summary:             "测试测试提现申请",
		ExtendInfo:          "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceConsumeApply(t *testing.T) {
	res, err := testPay.OrderService.ConsumeApply(service.ConsumeApplyReq{
		PayerId:             "test_user_68",
		RecieverId:          service.PlatformBizUserId,
		BizOrderNo:          "test_user_68_order3",
		Amount:              1,
		Fee:                 0,
		ValidateType:        service.ValidateTypeNone,
		SplitRule:           []service.SplitRuleItem{},
		FrontUrl:            frontUrl,
		BackUrl:             backUrl,
		OrderExpireDatetime: time.Now().Add(time.Hour * 12).Format("2006-01-02 15:04:05"),
		PayMethod: map[string]interface{}{
			"GATEWAY_VSP": map[string]interface{}{
				"amount":  1,
				"gateid":  "0103", //农业银行
				"paytype": "B2C",
			},
		},
		GoodsType:    service.GoodTypeEntity,
		BizGoodsNo:   "test_user_68_good1",
		GoodsName:    "测试消费申请",
		GoodsDesc:    "测试测试消费申请描述",
		IndustryCode: "1910",
		IndustryName: "其他",
		Source:       service.SourceMobile,
		Summary:      "测试测试消费申请",
		ExtendInfo:   "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceConsumeApplyForMp(t *testing.T) {
	res, err := testPay.OrderService.ConsumeApply(service.ConsumeApplyReq{
		PayerId:             "test_user_68",
		RecieverId:          service.PlatformBizUserId,
		BizOrderNo:          "test_user_68_order17",
		Amount:              1,
		Fee:                 0,
		ValidateType:        service.ValidateTypeNone,
		SplitRule:           []service.SplitRuleItem{},
		FrontUrl:            frontUrl,
		BackUrl:             backUrl,
		OrderExpireDatetime: time.Now().Add(time.Hour * 12).Format("2006-01-02 15:04:05"),
		PayMethod: map[string]interface{}{
			"WECHATPAY_MINIPROGRAM": map[string]interface{}{
				"amount":   1,
				"limitPay": "no_credit",
				"acct":     "otA9K5B_wIGr2UkdP8mLZWbbyuVU", //openid
			},
		},
		GoodsType:    service.GoodTypeEntity,
		BizGoodsNo:   "test_user_68_good1",
		GoodsName:    "测试消费申请",
		GoodsDesc:    "测试测试消费申请描述",
		IndustryCode: "1910",
		IndustryName: "其他",
		Source:       service.SourceMobile,
		Summary:      "测试测试消费申请",
		ExtendInfo:   "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceAgentCollectApply(t *testing.T) {
	res, err := testPay.OrderService.AgentCollectApply(service.AgentCollectApplyReq{
		BizOrderNo:          "test_user_68_order7",
		PayerId:             "test_user_68",
		RecieverList:        []service.RecieverItem{{BizUserId: "test_user_68", Amount: 1}},
		GoodsType:           service.GoodTypeEntity,
		BizGoodsNo:          "test_user_68_good2",
		TradeCode:           "1001",
		Amount:              1,
		Fee:                 0,
		ValidateType:        service.ValidateTypeSms,
		FrontUrl:            frontUrl,
		BackUrl:             backUrl,
		OrderExpireDatetime: time.Now().Add(time.Hour * 12).Format("2006-01-02 15:04:05"),
		PayMethod: map[string]interface{}{
			"GATEWAY_VSP": map[string]interface{}{
				"amount":  1,
				"gateid":  "0103", //农业银行
				"paytype": "B2C",
			},
		},
		GoodsName:    "测试消费申请",
		GoodsDesc:    "测试测试消费申请描述",
		IndustryCode: "1910",
		IndustryName: "其他",
		Source:       service.SourceMobile,
		Summary:      "测试测试消费申请",
		ExtendInfo:   "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceSignalAgentPay(t *testing.T) {
	res, err := testPay.OrderService.SignalAgentPay(service.SignalAgentPayReq{
		BizOrderNo:     "test_user_68_order5",
		CollectPayList: []service.CollectPayItem{{BizOrderNo: "test_user_68_order4", Amount: 1}},
		BizUserId:      "test_user_68",
		AccountSetNo:   service.AccountSetNoUserTest,
		BackUrl:        backUrl,
		Amount:         1,
		Fee:            0,
		SplitRuleList:  []service.SplitRuleItem{},
		GoodsType:      service.GoodTypeEntity,
		BizGoodsNo:     "test_user_68_good1",
		TradeCode:      "2001",
		Summary:        "测试测试单笔托管代付",
		ExtendInfo:     "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceBatchAgentPay(t *testing.T) {
	res, err := testPay.OrderService.BatchAgentPay(service.BatchAgentPayReq{
		BizBatchNo: "test_user_68",
		BatchPayList: []service.BatchPayItem{
			{
				BizOrderNo:     "test_user_68_order8",
				CollectPayList: []service.CollectPayItem{{BizOrderNo: "test_user_68_order7", Amount: 1}},
				BizUserId:      "test_user_68",
				AccountSetNo:   service.AccountSetNoUserTest,
				BackUrl:        backUrl,
				Amount:         1,
				Fee:            0,
				SplitRuleList:  []service.SplitRuleItem{},
				Summary:        "测试测试批量托管代付",
				ExtendInfo:     "",
			},
		},
		GoodsType:  service.GoodTypeEntity,
		BizGoodsNo: "test_user_68_good2",
		TradeCode:  "2001",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

//后台+验证码
func TestOrderServicePay(t *testing.T) {
	res, err := testPay.OrderService.Pay(service.PayReq{
		BizUserId:        "test_user_68",
		BizOrderNo:       "test_user_68_order15",
		TradeNo:          "",
		JumpUrl:          frontUrl,
		VerificationCode: "415301",
		ConsumerIp:       "119.123.132.108",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

//前台+验证码
func TestOrderServiceGetPayCodePage(t *testing.T) {
	url, err := testPay.OrderService.GetPayCodePage(
		"http://116.228.64.55:6900/yungateway/frontTrans.do",
		service.PayReq{
			BizUserId:        "test_user_68",
			BizOrderNo:       "test_user_68_order7",
			TradeNo:          "",
			JumpUrl:          "",
			VerificationCode: "11111",
			ConsumerIp:       "119.123.132.108",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

//前台+密码
func TestOrderServiceGetPayPasswordPage(t *testing.T) {
	url, err := testPay.OrderService.GetPayPasswordPage(
		"http://116.228.64.55:6900/yungateway/pwd/payOrder.html",
		service.PayReq{
			BizUserId:        "test_user_68",
			BizOrderNo:       "test_user_68_order13",
			TradeNo:          "",
			JumpUrl:          "",
			VerificationCode: "111111",
			ConsumerIp:       "119.123.132.108",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestOrderServiceEntryGoods(t *testing.T) {
	res, err := testPay.OrderService.EntryGoods(service.EntryGoodsReq{
		BizUserId:   "test_user_68",
		GoodsType:   service.GoodTypeEntity,
		BizGoodsNo:  "test_user_68_good2",
		GoodsName:   "沙发2",
		GoodsDetail: "三人沙发2",
		GoodsParams: map[string]interface{}{},
		ShowUrl:     "https://baidu.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceFreezeMoney(t *testing.T) {
	res, err := testPay.OrderService.FreezeMoney(service.FreezeMoneyReq{
		BizUserId:    "test_user_68",
		BizFreezenNo: "test_user_68_order9",
		AccountSetNo: service.AccountSetNoUserTest,
		Amount:       1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceUnfreezeMoney(t *testing.T) {
	res, err := testPay.OrderService.UnfreezeMoney(service.UnfreezeMoneyReq{
		BizUserId:    "test_user_68",
		BizFreezenNo: "test_user_68_order9",
		AccountSetNo: service.AccountSetNoUserTest,
		Amount:       1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceApplicationTransfer(t *testing.T) {
	res, err := testPay.OrderService.ApplicationTransfer(service.ApplicationTransferReq{
		BizTransferNo:      "test_user_68_order20",
		SourceAccountSetNo: service.AccountSetNoPlatform,
		TargetBizUserId:    "test_user_68",
		TargetAccountSetNo: service.AccountSetNoUserTest,
		Amount:             100000,
		ExtendInfo:         "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceRefund(t *testing.T) {
	res, err := testPay.OrderService.Refund(service.RefundReq{
		BizOrderNo:    "test_user_68_order10",
		OriBizOrderNo: "test_user_68_order7",
		BizUserId:     "test_user_68",
		RefundType:    "D0",
		RefundList:    []service.RefundItem{},
		BackUrl:       backUrl,
		Amount:        1,
		CouponAmount:  0,
		FeeAmount:     0,
		ExtendInfo:    "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceQueryBalance(t *testing.T) {
	res, err := testPay.OrderService.QueryBalance(service.QueryBalanceReq{
		BizUserId:    "test_user_68",
		AccountSetNo: service.AccountSetNoUserTest,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestOrderServiceGetOrderDetail(t *testing.T) {
	res, err := testPay.OrderService.GetOrderDetail(service.GetOrderDetailReq{
		BizOrderNo: "test_user_68_order7",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestMerchantServiceGetCheckAccountFile(t *testing.T) {
	res, err := testPay.MerchantService.GetCheckAccountFile(service.GetCheckAccountFileReq{
		Date:     time.Now().AddDate(0, 0, -1).Format("20060102"),
		FileType: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestQueryReserveFundBalance(t *testing.T) {
	res, err := testPay.MerchantService.QueryReserveFundBalance()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestQueryMerchantBalance(t *testing.T) {
	res, err := testPay.MerchantService.QueryMerchantBalance(service.QueryMerchantBalanceReq{
		AccountSetNo: service.AccountSetNoPlatform,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
