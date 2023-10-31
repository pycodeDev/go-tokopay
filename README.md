# go-tokopay Client for Golang

This library is the abstraction of go-tokopay API for access from applications written with Golang.

## Instalasi

```bash
go get github.com/pycodeDev/go-tokopay
```

## How To Use
Get Your Merchant ID and Secret Key at [go-tokopay Dashboard](https://dash.go-tokopay.id/pengaturan/secret-key).

### Get Info Merchant
prepare your struct to use this func.
```golang
type ParamInfoMerchant struct {
    MerchantId string
    Secret string
    Idr bool
}
```
> Note:
> Idr (true) if you want get your balance in indonesian rupiahs format

call the func

```golang
import go_tokopay "github.com/pycodeDev/go-tokopay"

func main() {
    response, err := go_tokopay.GetInfoMerchant(ParamInfoMerchant)
    if err != nil
    {
        fmt.Println(err)
    }

    //response
    response.MerchantName //name of merchant
    response.Balance // balance of merchant available to pull
    response.BalanceHold // balance of merchant hold
}
```

### Order Simple
prepare your struct to use this func.
```golang
type ParamOrderSimple struct {
    MerchantId string
    Secret string
    Method string
    RefId string
    Nominal int
    Idr bool
}
```
> Note:
> Idr (true) if you want get your balance in indonesian rupiahs format

call the func

```golang
import go_tokopay "github.com/pycodeDev/go-tokopay"

func main() {
    response, err := go_tokopay.Order(ParamOrderSimple)
    if err != nil
    {
        fmt.Println(err)
    }

    response.Data // like nomor_va, checkout_url, qr_link
    response.PanduanPembayaran // guide for pay
    response.PayUrl // route after pay the payment
    response.Status // status transaction
    response.TrxId // trx id if your transaction problem, send to go-tokopay cs
}
```

### License

[MIT]

### Author

[Muhammad Anang Ramadhan](mailto:muhammadanangr@gmail.com)
