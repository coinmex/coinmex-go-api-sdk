# coinmex-go-api-sdk
A GO SDK for CoinMex Exchange API

1.前往 CoinMex 创建 APIKey 并将信息写入 test_helper.go

```go

    /*
     用户 apiKey，需用户填写，在 https://www.coinmex.com/user  api 中获取
     */
    config.ApiKey = ""
    /*
     用户 secretKey，需用户填写，在 https://www.coinmex.com/user  api 中获取
     */
    config.SecretKey = ""
    /*
     口令，需用户填写，在 https://www.coinmex.com/user  api 中获取（创建时由用户设定）
     */
    config.Passphrase = ""
```

2.运行 spot_test.go 用例
