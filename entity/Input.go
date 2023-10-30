package entity

type ParamInfoMerchant struct {
	MerchantId string
	Secret     string
	Idr        bool
}

type ParamOrderSimple struct {
	MerchantId string
	Secret     string
	Method     string
	RefId      string
	Nominal    int
	Idr        bool
}

type OutputInfoMerchant struct {
	MerchantName string
	Balance      string
	BalanceHold  string
}

type OutputOrderSimple struct {
	Status            string
	TotalBayar        string
	TotalDiterima     string
	PayUrl            string
	PanduanPembayaran string
	TrxId             string
	Data              string
}
