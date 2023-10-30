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

func OrderSimple(param entity.ParamOrderSimple) (entity.OutputOrderSimple, error) {
	var d entity.OutputOrderSimple

	merchant_id := param.MerchantId
	secret := param.Secret
	idr := param.Idr
	method := param.Method
	ref_id := param.RefId
	nominal := param.Nominal

	order, err := lib.NewTokopayImpl(entity.TokoPayCreden{
		MERCHANTID: merchant_id,
		SECRETKEY:  secret,
	}).Order(entity.TokoPayOrderSimple{
		Method:  method,
		RefId:   ref_id,
		Nominal: nominal,
	})
	if err != nil {
		return d, err
	}

	if idr {
		d.TotalBayar = lib.ToIDR(lib.StringToInt(order.TotalBayar))
		d.TotalDiterima = lib.ToIDR(lib.StringToInt(order.TotalDiterima))
	} else {
		d.TotalBayar = order.TotalBayar
		d.TotalDiterima = order.TotalDiterima
	}

	d.Data = order.Data
	d.PanduanPembayaran = order.PanduanPembayaran
	d.PayUrl = order.PayUrl
	d.Status = order.Status
	d.TrxId = order.TrxId

	return d, nil
}
