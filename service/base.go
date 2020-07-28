package service

const (
	//平台会员id
	PlatformBizUserId = "#yunBizUserId_B2C#"

	//绑定个人会员的默认号码
	DefaultBindPhone = "88888888888"

	//会员类型
	MemberTypeCompany    = 2 //企业类型
	MemberTypeIndividual = 3 //个人类型

	//访问终端类型
	SourceMobile = 1
	SourcePC     = 2

	//身份类型 1身份证 2护照 3军官证 4回乡证 5台胞证 6警官证 7士兵证 99其他证件
	IdentityTypeIdCard = 1
	IdentityTypeOther  = 99

	//交易验证方式
	ValidateTypeNone     = 0 //无验证, 仅渠道验证,通商云不做交易验证
	ValidateTypeSms      = 1 //短信验证码, 通商云发送并验证短信验证码,有效期 3 分钟
	ValidateTypePassword = 2 //支付密码, 验证通商云支付密码

	//100001标准余额账户集 200126会员的账户余额是在托管账户集
	AccountSetNoUserTest = "200126" //通联测试环境用的
	AccountSetNoPlatform = "100001"

	//订单状态
	OrderStatusUnpaid  = 1
	OrderStatusFail    = 3
	OrderStatusSuccess = 4
	OrderStatusRefund  = 5 //交易成功,但是发生了退款
	OrderStatusClose   = 6
	OrderStatusOngoing = 99 //进行中

	//商品类型, 只有5需要录入商品参数
	GoodTypeVirtual     = 2  //虚拟
	GoodTypeEntity      = 3  //实物
	GoodTypeOffline     = 4  //线下
	GoodTypeCrossBorder = 5  //跨境
	GoodTypeMarketing   = 90 //营销活动
	GoodTypeOther       = 99 //其他
)
