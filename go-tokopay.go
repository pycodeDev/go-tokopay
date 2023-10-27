package main

import (
	"go-tokopay/entity"
	"go-tokopay/lib"
)

func GetInfoMerchant(param entity.ParamInfoMerchant) (entity.OutputInfoMerchant, error) {
	var d entity.OutputInfoMerchant

	merchant_id := param.MerchantId
	secret := param.Secret
	idr := param.Idr

	merchant, err := lib.NewTokopayImpl(entity.TokoPayCreden{
		MERCHANTID: merchant_id,
		SECRETKEY:  secret,
	}).GetInfo()
	if err != nil {
		return d, err
	}

	if idr {
		d.Balance = lib.ToIDR(lib.StringToInt(merchant.Saldo))
		d.BalanceHold = lib.ToIDR(lib.StringToInt(merchant.SaldoHold))
	} else {
		d.Balance = merchant.Saldo
		d.BalanceHold = merchant.SaldoHold
	}

	d.MerchantName = merchant.NamaToko

	return d, nil
}
