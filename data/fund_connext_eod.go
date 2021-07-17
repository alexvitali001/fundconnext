package data

// NAVConfirmation ...
type NAVConfirmation struct {
	AMCCode           string
	FundCode          string
	AUM               *float64
	NAV               *float64
	OfferNAV          *float64
	BidNAV            *float64
	SwitchOutNAV      *float64
	SwitchInNAV       *float64
	NAVDate           string
	SACodeForUnitLink *string
	TotalUnit         *float64
	TotalAUM          *float64
}

// UnitholderBalanceConfrimation ...
type UnitholderBalanceConfrimation struct {
	AMCCode              string
	AccountID            string
	UnitholderID         string
	FundCode             string
	UnitBalance          float64
	Amount               float64
	AvailableUnitBalance float64
	AvailableAmount      float64
	PendingUnit          float64
	PendingAmount        float64
	PledgeUnit           float64
	AverageCost          float64
	NAV                  float64
	NAVDate              string
}

// UnitholderBalanceLTFConfirmation ...
type UnitholderBalanceLTFConfirmation struct {
	AMCCode          string
	Filler1          *string
	UnitholderID     string
	FundCode         string
	Filler2          *string
	InvestmentYear   string
	AllowRedeemYear  string
	InvestmentAmount float64
	UnitInvest       float64
	AllowSellFlag    string
}

// AllottedTransactionConfrimation ...
type AllottedTransactionConfrimation struct {
	SAOrderReferenceNo                 *string
	TransactionDateTime                string
	AccountID                          string
	AMCCode                            string
	UnitholderID                       string
	Filler1                            *string
	TransactionCode                    string
	FundCode                           string
	OverrideRisKProfileFlag            string
	OverrideFXRiskFlag                 string
	RedemptionType                     *string
	Amount                             *float64
	Unit                               *float64
	EffectiveDate                      string
	Filler2                            *string
	Filler3                            *string
	PaymentType                        *string
	BankCode                           *string
	BankAccount                        *string
	ChequeNo                           *string
	ChequeDate                         *string
	ICLicense                          string
	BranchNo                           *string
	Channel                            string
	ForceEntry                         string
	LTFCondition                       *string
	ReasonToSellLTF                    *string
	RMFCapitalGainWithholdingTaxChoice *string
	RMFCapitalAmountRedeemChoice       *string
	AutoRedeemFundCode                 *string
	TransactionID                      string
	Status                             string
	AMCOrderReferenceNo                *string
	AllotmentDate                      *string
	AllottedNAV                        *float64
	AllottedAmount                     *float64
	AllotedUnit                        *float64
	Fee                                *float64
	WithholdingTax                     *float64
	VAT                                *float64
	BrokerageFee                       *float64
	WithholdingTaxForLTF               *float64
	AMCPayDate                         *string
	RegistrarTransactionFlag           *string
	SellAllUnitFlag                    *string
	SettlementBankCode                 *string
	SettlementBankAccount              *string
	RejectReason                       *string
	CHQBranch                          *string
	TaxInvoiceNo                       *string
	AMCRelatedOrderReferenceNo         *string
	ICCode                             *string
	BrokerageFeeVAT                    *float64
	ApprovalCode                       *string
	NAVDate                            *string
	CollateralAccount                  *string
	CreditCardIssuer                   *string
}
