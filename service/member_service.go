package service

import "github.com/hhjpin/allinpayapi/core"

type MemberService struct {
	base    *core.Pay
	service string
}

func NewMemberService(base *core.Pay) *MemberService {
	return &MemberService{
		base:    base,
		service: "MemberService",
	}
}

type CreateMemberReq struct {
	BizUserId   string                 `json:"bizUserId"`  //商户系统用户标识,商户系统中唯一编号
	MemberType  int                    `json:"memberType"` //2企业会员 3个人会员
	Source      int                    `json:"source"`     //1手机 2pc
	ExtendParam map[string]interface{} `json:"extendParam"`
}

type CreateMemberRsp struct {
	BizUserId string `json:"bizUserId"` //商户系统用户标识,商户系统中唯一编号
	UserId    string `json:"userId"`    //通商云用户唯一标识
}

//创建会员
func (s *MemberService) CreateMember(req CreateMemberReq) (*CreateMemberRsp, error) {
	var res = &CreateMemberRsp{}
	if req.ExtendParam == nil {
		req.ExtendParam = map[string]interface{}{}
	}
	err := s.base.CommonRequest(s.service, "createMember", req, res)
	return res, err
}

type SendVerificationCodeReq struct {
	BizUserId            string `json:"bizUserId"`
	Phone                string `json:"phone"`
	VerificationCodeType int    `json:"verificationCodeType"` //9绑定手机 6解绑手机
}

type SendVerificationCodeRsp struct {
	BizUserId string `json:"bizUserId"`
	Phone     string `json:"phone"`
}

//发送短信验证码
func (s *MemberService) SendVerificationCode(req SendVerificationCodeReq) (*SendVerificationCodeRsp, error) {
	var res = &SendVerificationCodeRsp{}
	err := s.base.CommonRequest(s.service, "sendVerificationCode", req, res)
	return res, err
}

type BindPhoneReq struct {
	BizUserId        string `json:"bizUserId"`
	Phone            string `json:"phone"`
	VerificationCode string `json:"verificationCode"`
}

type BindPhoneRsp struct {
	BizUserId string `json:"bizUserId"`
	Phone     string `json:"phone"`
}

//绑定手机
func (s *MemberService) BindPhone(req BindPhoneReq) (*BindPhoneRsp, error) {
	var res = &BindPhoneRsp{}
	err := s.base.CommonRequest(s.service, "bindPhone", req, res)
	return res, err
}

type UnbindPhoneReq struct {
	BizUserId        string `json:"bizUserId"`
	Phone            string `json:"phone"`
	VerificationCode string `json:"verificationCode"`
}

type UnbindPhoneRsp struct {
	BizUserId string `json:"bizUserId"`
	Phone     string `json:"phone"`
}

//绑定手机
func (s *MemberService) UnbindPhone(req UnbindPhoneReq) (*UnbindPhoneRsp, error) {
	var res = &UnbindPhoneRsp{}
	err := s.base.CommonRequest(s.service, "unbindPhone", req, res)
	return res, err
}

type SignContractReq struct {
	BizUserId string `json:"bizUserId"`
	JumpUrl   string `json:"jumpUrl"`
	BackUrl   string `json:"backUrl"`
	Source    int    `json:"source"` //1手机 2PC
}

type SignContractRsp struct {
	BizUserId  string `json:"bizUserId"`
	ContractNo string `json:"contractNo"`
	Result     string `json:"result"`
}

func (s *MemberService) SignContractPage(host string, req SignContractReq) (string, error) {
	if host == "" {
		host = "https://fintech.allinpay.com/yungateway/member/signContract.html"
	}
	param, err := s.base.GetRequestParam(s.service, "signContract", req)
	return host + param, err
}

//会员电子协议签约
func (s *MemberService) SignContract(req SignContractReq) (*SignContractRsp, error) {
	var res = &SignContractRsp{}
	err := s.base.CommonRequest(s.service, "signContract", req, res)
	return res, err
}

type SetRealNameReq struct {
	BizUserId    string `json:"bizUserId"`
	IsAuth       bool   `json:"isAuth"`
	Name         string `json:"name"`
	IdentityType int    `json:"identityType"` //1身份证 2护照 3军官证 4回乡证 5台胞证 6警官证 7士兵证 99其他证件
	IdentityNo   string `json:"identityNo"`
}

type SetRealNameRsp struct {
	BizUserId    string `json:"bizUserId"`
	Name         string `json:"name"`
	IdentityType int    `json:"identityType"`
	IdentityNo   string `json:"identityNo"`
}

//个人实名认证
func (s *MemberService) SetRealName(req SetRealNameReq) (*SetRealNameRsp, error) {
	var res = &SetRealNameRsp{}
	var err error
	if req.IdentityNo, err = s.base.Encrypt(req.IdentityNo); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "setRealName", req, res)
	if err != nil {
		return res, err
	}
	if res.IdentityNo, err = s.base.Decrypt(res.IdentityNo); err != nil {
		return res, err
	}
	return res, err
}

type CompanyBasicInfo struct {
	CompanyName      string `json:"companyName"`
	CompanyAddress   string `json:"companyAddress"`
	AuthType         int    `json:"authType"` //1三证 2一证
	UniCredit        string `json:"uniCredit"`
	BusinessLicense  string `json:"businessLicense"`
	OrganizationCode string `json:"organizationCode"`
	TaxRegister      string `json:"taxRegister"`
	ExpLicense       string `json:"expLicense"`
	Telephone        string `json:"telephone"`
	LegalName        string `json:"legalName"`
	IdentityType     int    `json:"identityType"` //1身份证
	LegalIds         string `json:"legalIds"`
	LegalPhone       string `json:"legalPhone"`
	AccountNo        string `json:"accountNo"`
	ParentBankName   string `json:"parentBankName"`
	BankCityNo       string `json:"bankCityNo"`
	BankName         string `json:"bankName"`
	UnionBank        string `json:"unionBank"`
	Province         string `json:"province"`
	City             string `json:"city"`
}

type SetCompanyInfoReq struct {
	BizUserId        string           `json:"bizUserId"`
	BackUrl          string           `json:"backUrl"`
	CompanyBasicInfo CompanyBasicInfo `json:"companyBasicInfo"`
	//CompanyExtendInfo map[string]interface{} `json:"companyExtendInfo"` //不需传
	IsAuth bool `json:"isAuth"`
}

type SetCompanyInfoRsp struct {
	BizUserId  string `json:"bizUserId"`
	Result     int    `json:"result"`
	FailReason string `json:"failReason"`
	Remark     string `json:"remark"`
}

//设置企业信息
func (s *MemberService) SetCompanyInfo(req SetCompanyInfoReq) (*SetCompanyInfoRsp, error) {
	var res = &SetCompanyInfoRsp{}
	var err error
	if req.CompanyBasicInfo.AccountNo, err = s.base.Encrypt(req.CompanyBasicInfo.AccountNo); err != nil {
		return res, err
	}
	if req.CompanyBasicInfo.LegalIds, err = s.base.Encrypt(req.CompanyBasicInfo.LegalIds); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "setCompanyInfo", req, res)
	return res, err
}

type VerifyResultReq struct {
	BizUserId  string `json:"bizUserId"`
	Result     int    `json:"result"`
	CheckTime  string `json:"checkTime"`
	Remark     string `json:"remark"`
	FailReason string `json:"failReason"`
}

type VerifyResultRsp struct {
	BizUserId string `json:"bizUserId"`
	//todo 文档没有
}

//企业信息审核结果通知
func (s *MemberService) VerifyResult(req VerifyResultReq) (*VerifyResultRsp, error) {
	var res = &VerifyResultRsp{}
	err := s.base.CommonRequest(s.service, "verifyResult", req, res)
	return res, err
}

type GetMemberInfoReq struct {
	BizUserId string `json:"bizUserId"`
}

type IndividualMemberInfo struct {
	Name              string `json:"name"`
	UserState         int    `json:"userState"`
	UserId            string `json:"userId"`
	Country           string `json:"country"`
	Province          string `json:"province"`
	Area              string `json:"area"`
	Address           string `json:"address"`
	Phone             string `json:"phone"`
	IdentityCardNo    string `json:"identityCardNo"`
	IsPhoneChecked    bool   `json:"isPhoneChecked"`
	RegisterTime      string `json:"registerTime"`
	RegisterIp        string `json:"registerIp"`
	PayFailAmount     int    `json:"payFailAmount"`
	IsIdentityChecked bool   `json:"isIdentityChecked"`
	RealNameTime      string `json:"realNameTime"`
	Remark            string `json:"remark"`
	Source            int    `json:"source"`
	IsSetPayPwd       bool   `json:"isSetPayPwd"`
	IsSignContract    bool   `json:"isSignContract"`
	AcctOrgType       int    `json:"acctOrgType"`
	SubAcctNo         string `json:"subAcctNo"`
}

type GetMemberInfoForIndividualRsp struct {
	BizUserId  string               `json:"bizUserId"`
	MemberType int                  `json:"memberType"`
	MemberInfo IndividualMemberInfo `json:"memberInfo"`
}

//获取会员信息(个人)
func (s *MemberService) GetMemberInfoForIndividual(req GetMemberInfoReq) (*GetMemberInfoForIndividualRsp, error) {
	var res = &GetMemberInfoForIndividualRsp{}
	err := s.base.CommonRequest(s.service, "getMemberInfo", req, res)
	if err != nil {
		return res, err
	}
	if res.MemberInfo.IdentityCardNo, err = s.base.Decrypt(res.MemberInfo.IdentityCardNo); err != nil {
		return res, err
	}
	return res, err
}

type CompanyMemberInfo struct {
	CompanyName      string `json:"companyName"`
	CompanyAddress   string `json:"companyAddress"`
	AuthType         int    `json:"authType"`
	BusinessLicense  string `json:"businessLicense"`
	OrganizationCode string `json:"organizationCode"`
	UniCredit        string `json:"uniCredit"`
	TaxRegister      string `json:"taxRegister"`
	ExpLicense       string `json:"expLicense"`
	Telephone        string `json:"telephone"`
	Phone            string `json:"phone"`
	LegalName        string `json:"legalName"`
	IdentityType     int    `json:"identityType"`
	LegalIds         string `json:"legalIds"`
	LegalPhone       string `json:"legalPhone"`
	AccountNo        string `json:"accountNo"`
	ParentBankName   string `json:"parentBankName"`
	BankCityNo       string `json:"bankCityNo"`
	BankName         string `json:"bankName"`
	UnionBank        string `json:"unionBank"`
	Province         string `json:"province"`
	City             string `json:"city"`
	IsSignContract   bool   `json:"isSignContract"`
	Status           int    `json:"status"`
	CheckTime        string `json:"checkTime"`
	Remark           string `json:"remark"`
	FailReason       string `json:"failReason"`
	AcctOrgType      int    `json:"acctOrgType"`
	SubAcctNo        string `json:"subAcctNo"`
}

type GetMemberInfoForCompanyRsp struct {
	BizUserId  string            `json:"bizUserId"`
	MemberType int               `json:"memberType"`
	MemberInfo CompanyMemberInfo `json:"memberInfo"`
}

//获取会员信息(企业)
func (s *MemberService) GetMemberInfoForCompany(req GetMemberInfoReq) (*GetMemberInfoForCompanyRsp, error) {
	var res = &GetMemberInfoForCompanyRsp{}
	err := s.base.CommonRequest(s.service, "getMemberInfo", req, res)
	if err != nil {
		return res, err
	}
	if res.MemberInfo.LegalIds, err = s.base.Decrypt(res.MemberInfo.LegalIds); err != nil {
		return res, err
	}
	if res.MemberInfo.AccountNo, err = s.base.Decrypt(res.MemberInfo.AccountNo); err != nil {
		return res, err
	}
	return res, err
}

type GetBankCardBinReq struct {
	CardNo string `json:"cardNo"`
}

type GetBankCardBinRsp struct {
	CardBinInfo struct {
		CardBin       string `json:"cardBin"`
		CardType      int    `json:"cardType"` //1借记卡 2信用卡
		BankCode      string `json:"bankCode"` //发卡行代码
		BankName      string `json:"bankName"`
		CardName      string `json:"cardName"`
		CardLenth     int    `json:"cardLenth"`
		CardState     int    `json:"cardState"`     //状态(1:有效;0:无效)
		CardTypeLabel string `json:"cardTypeLabel"` //卡种名称
	} `json:"cardBinInfo"`
}

//查询卡 bin
func (s *MemberService) GetBankCardBin(req GetBankCardBinReq) (*GetBankCardBinRsp, error) {
	var res = &GetBankCardBinRsp{}
	var err error
	if req.CardNo, err = s.base.Encrypt(req.CardNo); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "getBankCardBin", req, res)
	if err != nil {
		return res, err
	}
	return res, err
}

type ApplyBindBankCardReq struct {
	BizUserId    string `json:"bizUserId"`
	CardNo       string `json:"cardNo"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	CardCheck    int    `json:"cardCheck"` //绑卡方式 7：收银宝快捷支付签约
	IdentityType int    `json:"identityType"`
	IdentityNo   string `json:"identityNo"`
	Validate     string `json:"validate"`
	Cvv2         string `json:"cvv2"`
	IsSafeCard   bool   `json:"isSafeCard"`
	UnionBank    string `json:"unionBank"`
}

type ApplyBindBankCardRsp struct {
	BizUserId string `json:"bizUserId"`
	TranceNum string `json:"tranceNum"`
	TransDate string `json:"transDate"`
	BankName  string `json:"bankName"`
	BankCode  string `json:"bankCode"`
	CardType  int    `json:"cardType"`
}

func (s *MemberService) ApplyBindBankCard(req ApplyBindBankCardReq) (*ApplyBindBankCardRsp, error) {
	var res = &ApplyBindBankCardRsp{}
	var err error
	if req.CardNo, err = s.base.Encrypt(req.CardNo); err != nil {
		return res, err
	}
	if req.IdentityNo, err = s.base.Encrypt(req.IdentityNo); err != nil {
		return res, err
	}
	if req.Validate, err = s.base.Encrypt(req.Validate); err != nil {
		return res, err
	}
	if req.Cvv2, err = s.base.Encrypt(req.Cvv2); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "applyBindBankCard", req, res)
	return res, err
}

type BindBankCardReq struct {
	BizUserId        string `json:"bizUserId"`
	TranceNum        string `json:"tranceNum"`
	TransDate        string `json:"transDate"`
	Phone            string `json:"phone"`
	Validate         string `json:"validate"`
	Cvv2             string `json:"cvv2"`
	VerificationCode string `json:"verificationCode"`
}

type BindBankCardRsp struct {
	BizUserId string `json:"bizUserId"`
	TranceNum string `json:"tranceNum"`
	TransDate string `json:"transDate"`
}

func (s *MemberService) BindBankCard(req BindBankCardReq) (*BindBankCardRsp, error) {
	var res = &BindBankCardRsp{}
	var err error
	if req.Validate, err = s.base.Encrypt(req.Validate); err != nil {
		return res, err
	}
	if req.Cvv2, err = s.base.Encrypt(req.Cvv2); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "bindBankCard", req, res)
	return res, err
}

type UnbindBankCardReq struct {
	BizUserId string `json:"bizUserId"`
	CardNo    string `json:"cardNo"`
}

type UnbindBankCardRsp struct {
	BizUserId string `json:"bizUserId"`
	CardNo    string `json:"cardNo"`
}

func (s *MemberService) UnbindBankCard(req UnbindBankCardReq) (*UnbindBankCardRsp, error) {
	var res = &UnbindBankCardRsp{}
	var err error
	if req.CardNo, err = s.base.Encrypt(req.CardNo); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "unbindBankCard", req, res)
	if err != nil {
		return res, err
	}
	if res.CardNo, err = s.base.Decrypt(res.CardNo); err != nil {
		return res, err
	}
	return res, err
}

type LockMemberReq struct {
	BizUserId string `json:"bizUserId"`
}

type LockMemberRsp struct {
	BizUserId string `json:"bizUserId"`
}

//锁定会员
func (s *MemberService) LockMember(req LockMemberReq) (*LockMemberRsp, error) {
	var res = &LockMemberRsp{}
	err := s.base.CommonRequest(s.service, "lockMember", req, res)
	return res, err
}

type UnlockMemberReq struct {
	BizUserId string `json:"bizUserId"`
}

type UnlockMemberRsp struct {
	BizUserId string `json:"bizUserId"`
}

//解锁会员
func (s *MemberService) UnlockMember(req UnlockMemberReq) (*UnlockMemberRsp, error) {
	var res = &UnlockMemberRsp{}
	err := s.base.CommonRequest(s.service, "unlockMember", req, res)
	return res, err
}

type ApplyBindAcctReq struct {
	BizUserId     string `json:"bizUserId"`
	OperationType string `json:"operationType"` //set-绑定
	AcctType      string `json:"acctType"`      //支付账户类型
	Acct          string `json:"acct"`          //支付账户用户标识
}

type ApplyBindAcctRsp struct {
	BizUserId string `json:"bizUserId"`
	Result    string `json:"result"`
}

//会员绑定支付账户用户标识
func (s *MemberService) ApplyBindAcct(req ApplyBindAcctReq) (*ApplyBindAcctRsp, error) {
	var res = &ApplyBindAcctRsp{}
	err := s.base.CommonRequest(s.service, "applyBindAcct", req, res)
	return res, err
}

type BankCardChangeBindPhoneReq struct {
	BizUserId    string `json:"bizUserId"`
	CardNo       string `json:"cardNo"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	CardCheck    int    `json:"cardCheck"`
	IdentityType int    `json:"identityType"`
	IdentityNo   string `json:"identityNo"`
	Validate     string `json:"validate"`
	Cvv2         string `json:"cvv2"`
}

type BankCardChangeBindPhoneRsp struct {
	BizUserId string `json:"bizUserId"`
	TranceNum string `json:"tranceNum"`
	BankName  string `json:"bankName"`
	BankCode  string `json:"bankCode"`
	CardType  int    `json:"cardType"`
}

//修改绑定手机(银行卡验证)
func (s *MemberService) BankCardChangeBindPhone(req BankCardChangeBindPhoneReq) (*BankCardChangeBindPhoneRsp, error) {
	var res = &BankCardChangeBindPhoneRsp{}
	var err error
	if req.CardNo, err = s.base.Encrypt(req.CardNo); err != nil {
		return res, err
	}
	if req.IdentityNo, err = s.base.Encrypt(req.IdentityNo); err != nil {
		return res, err
	}
	if req.Validate, err = s.base.Encrypt(req.Validate); err != nil {
		return res, err
	}
	if req.Cvv2, err = s.base.Encrypt(req.Cvv2); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "bankCardChangeBindPhone", req, res)
	return res, err
}

type VerifyBankCardChangeBindPhoneReq struct {
	BizUserId        string `json:"bizUserId"`
	TranceNum        string `json:"tranceNum"`
	Phone            string `json:"phone"`
	VerificationCode string `json:"verificationCode"`
	Validate         string `json:"validate"`
	Cvv2             string `json:"cvv2"`
}

type VerifyBankCardChangeBindPhoneRsp struct {
	BizUserId string `json:"bizUserId"`
	TranceNum string `json:"tranceNum"`
}

//修改绑定手机(银行卡验证)
func (s *MemberService) VerifyBankCardChangeBindPhone(req VerifyBankCardChangeBindPhoneReq) (*VerifyBankCardChangeBindPhoneRsp, error) {
	var res = &VerifyBankCardChangeBindPhoneRsp{}
	var err error
	if req.Validate, err = s.base.Encrypt(req.Validate); err != nil {
		return res, err
	}
	if req.Cvv2, err = s.base.Encrypt(req.Cvv2); err != nil {
		return res, err
	}
	err = s.base.CommonRequest(s.service, "verifyBankCardChangeBindPhone", req, res)
	return res, err
}
