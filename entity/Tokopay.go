package entity

type TokopayCreden struct {
	MERCHANTID string
	SECRETKEY  string
}

type TokopayOrderSimple struct {
	Method  string
	RefId   string
	Nominal int
}

type TokopayOrderAdvanced struct {
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
