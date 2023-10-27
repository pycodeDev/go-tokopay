package entity

type TokoPayCreden struct {
	MERCHANTID string
	SECRETKEY  string
}

type TokoPayOrderSimple struct {
	MerchantId string
	Method     string
	Secret     string
	RefId      string
	Nominal    int
}

type TokoPayOrderAdvanced struct {
	MerchantId  string
	Code        string
	Secret      string
	RefId       string
	Amount      int
	CustName    string
	CustEmail   string
	CustPhone   string
	RedirectUrl string
	ExpiredTs   string
	Signature   string
	Items       []TokopayItemsOrderAdvanced
}

type TokopayItemsOrderAdvanced struct {
	ProductCode string
	Name        string
	Price       int
	ProductUrl  string
	ImageUrl    string
}

type TokopayGetMerchantInfo struct {
	MerchantId string
	Signature  string
}

type ReturnInfoMerchant struct {
	NamaToko  string
	Saldo     string
	SaldoHold string
}
