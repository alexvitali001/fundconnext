package order

// ...OrderBulk
type OrderBulk struct {
	SAOrderReferenceNo                 string
	TransactionDateTime                string
	AccountID                          string
	AMCCode                            string
	UnitholderID                       string
	Filler1                            *string
	TransactionCode                    string
	FundCode                           string
	OverrideRisKProfileFlag            string
	OverrideFXRiskFlag                 string
	RedemptionType                     string
	Amount                             *float64
	Unit                               *float64
	EffectiveDate                      string
	CounterFundCode                    string
	Filler2                            *string
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
	TransactionID                      *string
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
	Filler3                            *string
	CHQBranch                          *string
	Filler4                            *string
	Filler5                            *string
	Filler6                            *string
	Filler7                            *string
	Filler8                            *string
	Filler9                            *string
	CollateralAccount                  *string
}
