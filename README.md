# Tokopay Client for Golang

This library is the abstraction of tokopay API for access from applications written with Golang.

## Instalasi

```bash
go get github.com/pycodeDev/go-tokopay
```

## Pemakaian
Get Your Merchant ID and Secret Key at [Tokopay Dashboard](https://dash.tokopay.id/pengaturan/secret-key).
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