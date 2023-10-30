# Tokopay Client for Golang

This library is the abstraction of tokopay API for access from applications written with Golang.

## Instalasi

```bash
go get github.com/pycodeDev/go-tokopay
```

## How To Use
Get Your Merchant ID and Secret Key at [Tokopay Dashboard](https://dash.tokopay.id/pengaturan/secret-key).

### Get Info Merchant
prepare your struct to use this func.
```golang
type ParamInfoMerchant struct {
    MerchantId string
    Secret string
    Idr bool
}
```

call the func

```golang
    response, err := tokopay.GetInfoMerchant(ParamInfoMerchant)
    if err != nil
    {
        fmt.Println(err)
    }
```

get the response

```golang
    response.MerchantName //name of merchant
    response.Balance // balance of merchant available to pull
    response.BalanceHold // balance of merchant hol
```
> Note:
> Idr (true) if you want get your balance in indonesian rupiahs format
<!-- 
```js
const tokovoucher = require('tokovoucher');
const client = new tokovoucher("YOUR MERCHANT ID","YOUR SECRET");
```

### Get Info Merchant
```js
let saldo = await client.cekSaldo();
```

### Simple Order 

```js
let trx = await client.transaksi(refId, kodeProduk, tujuan, serverId);
```

> Note:
> RefID adalah kode transaksi unik kamu yang di generate secara acak -->

### License

[MIT](https://github.com/aripadrian/tokovoucher/blob/master/LICENSE)

### Author

[Muhammad Anang Ramadhan](mailto:muhammadanangr@gmail.com)