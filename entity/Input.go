package entity

type ParamInfoMerchant struct {
	MerchantId string
	Secret     string
	Idr        bool
}

type OutputInfoMerchant struct {
	MerchantName string
	Balance      string
	BalanceHold  string
}
