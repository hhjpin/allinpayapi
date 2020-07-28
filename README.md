# allinpayapi 

通联支付API库



# Usage

获取golang包

```sh
go get -u github.com/hhjpin/allinpayapi
```

使用例子

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/hhjpin/allinpayapi"
    "github.com/hhjpin/allinpayapi/core"
    "github.com/hhjpin/allinpayapi/pem"
    "github.com/hhjpin/allinpayapi/service"
)

func main() {
    //初始化支付实例
    pay, err := allinpayapi.New(&core.Config{
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
    
    //创建会员
    res, err := pay.MemberService.CreateMember(service.CreateMemberReq{
        BizUserId:   "test_user_01",
        MemberType:  service.MemberTypeCompany,
        Source:      service.SourceMobile,
        ExtendParam: map[string]interface{}{},
    })
    if err != nil {
        panic(err)
    }
    fmt.Println("res:", res)
}
```

更多的例子请参考 `allinpay_test.go`


# License

MIT