package lib

import (
	"errors"
	"fmt"
	"tokopay/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

const URL = "https://api.tokopay.id/v1"

type Tokopay interface {
	Order(params entity.TokoPayOrderSimple) (entity.OutputOrderSimple, error)
	OrderAdvanced(params entity.TokoPayOrderAdvanced)
	GetInfo() (entity.ReturnInfoMerchant, error)
}

type TokopayImpl struct {
	MerchantId string
	Secret     string
}

func NewTokopayImpl(param entity.TokoPayCreden) Tokopay {
	return &TokopayImpl{
		MerchantId: param.MERCHANTID,
		Secret:     param.SECRETKEY}
}

func (t TokopayImpl) Order(params entity.TokoPayOrderSimple) (entity.OutputOrderSimple, error) {
	var output entity.OutputOrderSimple
	merchant_id := t.MerchantId
	secret := t.Secret
	method := params.Method
	ref_id := params.RefId
	nominal := params.Nominal

	url := fmt.Sprintf("%v/order?merchant=%v&secret=%v&ref_id=%v&nominal=%v&metode=%v", URL, merchant_id, secret, ref_id, nominal, method)

	response, err := SendGet(url)
	if err != nil {
		return output, err
	}

	dataJson := gjson.GetMany(response, "status", "error_msg", "data.trx_id", "data.total_bayar", "data.total_diterima", "data.pay_url", "data.panduan_pembayaran", "data.nomor_va", "data.checkout_url", "data.qr_link")
	if !dataJson[0].Exists() {
		return output, errors.New("HTTP ERROR TOKOPAY !!")
	}

	status := dataJson[0].String()

	if status != "Success" {
		error_msg := dataJson[1].String()
		return output, errors.New(error_msg)
	}

	output.TrxId = dataJson[2].String()
	output.TotalBayar = dataJson[3].String()
	output.TotalDiterima = dataJson[4].String()
	output.PayUrl = dataJson[5].String()
	output.PanduanPembayaran = dataJson[6].String()

	if dataJson[7].Exists() {
		output.Data = dataJson[7].String()
	}

	if dataJson[8].Exists() {
		output.Data = dataJson[8].String()
	}

	if dataJson[9].Exists() {
		output.Data = dataJson[9].String()
	}

	return output, nil
}

func (t TokopayImpl) OrderAdvanced(params entity.TokoPayOrderAdvanced) {

}

func (t TokopayImpl) GetInfo() (entity.ReturnInfoMerchant, error) {
	var output entity.ReturnInfoMerchant
	merchant_id := t.MerchantId
	secret := t.Secret

	sign := GenerateSignatureDefault(merchant_id, secret)
	url := fmt.Sprintf("%v/merchant", URL)
	data_payload := fiber.Map{
		"merchant_id": merchant_id,
		"signature":   sign,
	}

	response, err := SendPostJson(url, data_payload)
	if err != nil {
		return output, err
	}

	dataJson := gjson.GetMany(response, "status", "error_msg", "data.nama_toko", "data.saldo_tersedia", "data.saldo_tertahan")
	if !dataJson[0].Exists() {
		return output, errors.New("HTTP ERROR TOKOPAY !!")
	}

	status := dataJson[0].String()

	if status == "0" {
		error_msg := dataJson[1].String()
		return output, errors.New(error_msg)
	}

	output.NamaToko = dataJson[2].String()
	output.Saldo = dataJson[3].String()
	output.SaldoHold = dataJson[4].String()

	return output, nil
}
