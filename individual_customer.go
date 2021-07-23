package fundconnext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Verified V4
type SpouseDocument struct {
	THFirstName *string `json:"thFirstName"`
	THLastName  *string `json:"thLastName"`
	ENFirstName *string `json:"enFirstName"`
	ENLastName  *string `json:"enLastName"`
}

// Verified V4
type SuitabilityForm struct {
	SuitNo1  *string `json:"suitNo1"`
	SuitNo2  *string `json:"suitNo2"`
	SuitNo3  *string `json:"suitNo3"`
	SuitNo4  *string `json:"suitNo4"`
	SuitNo5  *string `json:"suitNo5"`
	SuitNo6  *string `json:"suitNo6"`
	SuitNo7  *string `json:"suitNo7"`
	SuitNo8  *string `json:"suitNo8"`
	SuitNo9  *string `json:"suitNo9"`
	SuitNo10 *string `json:"suitNo10"`
	SuitNo11 *string `json:"suitNo11"`
	SuitNo12 *string `json:"suitNo12"`
}

// Verified V4
type Address struct {
	No          string  `json:"no"`
	Floor       *string `json:"floor"`
	Building    *string `json:"building"`
	RoomNo      *string `json:"roomNo"`
	Soi         *string `json:"soi"`
	Road        *string `json:"road"`
	Moo         *string `json:"moo"`
	SubDistrict string  `json:"subDistrict"`
	District    string  `json:"district"`
	Province    string  `json:"province"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
}

type PartialIndividualCustomerDocument struct {
	IdentificationCardType    string  `json:"identificationCardType"`
	PassportCountry           *string `json:"passportCountry"`
	CardNumber                string  `json:"cardNumber"`
	SuitabilityRiskLevel      *string `json:"suitabilityRiskLevel"`
	SuitabilityEvaluationDate *string `json:"suitabilityEvaluationDate"`
	Fatca                     *bool   `json:"fatca"`
	FatcaDeclarationDate      *string `json:"fatcaDeclarationDate"`
	CddScore                  *string `json:"cddScore"`
	CddDate                   *string `json:"cddDate"`
	ReferralPerson            *string `json:"referralPerson"`
	Approved                  bool    `json:"approved"`
}

type IndividualCustomerDocument struct {
	IdentificationCardType         string           `json:"identificationCardType"`
	PassportCountry                *string          `json:"passportCountry"`
	CardNumber                     string           `json:"cardNumber"`
	CardExpiryDate                 string           `json:"cardExpiryDate"`
	AccompanyingDocument           *string          `json:"accompanyingDocument"`
	Title                          string           `json:"title"`
	TitleOther                     *string          `json:"titleOther"`
	ENFirstName                    string           `json:"enFirstName"`
	ENLastName                     string           `json:"enLastName"`
	THFirstName                    string           `json:"thFirstName"`
	THLastName                     string           `json:"thLastName"`
	BirthDate                      string           `json:"birthDate"`
	Nationality                    string           `json:"nationality"`
	MobileNumber                   string           `json:"mobileNumber"`
	Email                          *string          `json:"email"`
	Phone                          *string          `json:"phone"`
	Fax                            *string          `json:"fax"`
	MaritalStatus                  string           `json:"maritalStatus"`
	Spouse                         *SpouseDocument  `json:"spouse"`
	OccupationId                   string           `json:"occupationId"`
	OccupationOther                *string          `json:"occupationOther"`
	BusinessTypeId                 *string          `json:"businessTypeId"`
	BusinessTypeOther              *string          `json:"businessTypeOther"`
	MonthlyIncomeLevel             string           `json:"monthlyIncomeLevel"`
	AssetValue                     *string          `json:"assetValue"`
	IncomeSource                   string           `json:"incomeSource"`
	IncomeSourceOther              *string          `json:"incomeSourceOther"`
	IdentificationDocument         Address          `json:"identificationDocument"`
	CurrentAddressSameAsFlag       *string          `json:"currentAddressSameAsFlag"`
	Current                        *Address         `json:"current"`
	CompanyName                    *string          `json:"companyName"`
	Work                           *Address         `json:"work"`
	WorkPosition                   *string          `json:"workPosition"`
	RelatedPoliticalPerson         *bool            `json:"relatedPoliticalPerson"`
	PoliticalRelatedPersonPosition *string          `json:"politicalRelatedPersonPosition"`
	CanAcceptFxRisk                bool             `json:"canAcceptFxRisk"`
	CanAcceptDerivativeInvestment  bool             `json:"canAcceptDerivativeInvestment"`
	SuitabilityRiskLevel           string           `json:"suitabilityRiskLevel"`
	SuitabilityEvaluationDate      string           `json:"suitabilityEvaluationDate"`
	Fatca                          bool             `json:"fatca"`
	FatcaDeclarationDate           string           `json:"fatcaDeclarationDate"`
	CDDScore                       *string          `json:"cddScore"`
	CDDDate                        *string          `json:"cddDate"`
	ReferralPerson                 *string          `json:"referralPerson"`
	ApplicationDate                string           `json:"applicationDate"`
	IncomeSourceCountry            string           `json:"incomeSourceCountry"`
	AcceptBy                       *string          `json:"acceptBy"`
	OpenFundConnextFormFlag        string           `json:"openFundConnextFormFlag"`
	Approved                       bool             `json:"approved"`
	VulnerableFlag                 *bool            `json:"vulnerableFlag"`
	VulnerableDetail               *string          `json:"vulnerableDetail"`
	NDIDFlag                       *bool            `json:"ndidFlag"`
	NDIDRequestId                  *string          `json:"ndidRequestId"`
	OpenChannel                    *string          `json:"openChannel"`
	SuitabilityForm                *SuitabilityForm `json:"suitabilityForm"`
	InvestorClass                  *string          `json:"investorClass"`
}

type IndividualCustomerFile struct {
	IdentificationCardType string  `json:"identificationCardType"`
	PassportCountry        *string `json:"passportCountry"`
	CardNumber             string  `json:"cardNumber"`
	Approved               bool    `json:"approved"`
}

func (f *FundConnext) CreateIndividualCustomer(identificationDoc IndividualCustomerDocument) error {
	cfg := MakeAPICallerConfig(f)
	url := "/api/customer/individual/v4"

	body, err := json.Marshal(identificationDoc)

	if err != nil {
		fmt.Println("[Func CreateIndividualCustomer] Error json.Marshal ::", err)
		return err
	}

	header := Headers{
		ContentType: "application/json",
		XAuthToken:  f.token,
	}

	//_, err = CallFCAPI(f.token, "POST", url, body, cfg)
	_, _, err = CallToFundConnext(cfg, "POST", url, header, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("[Func CreateIndividualCustomer] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

func (f *FundConnext) UpdateIndividualCustomer(identificationDoc IndividualCustomerDocument) error {
	cfg := MakeAPICallerConfig(f)
	url := "/api/customer/individual/v4"

	body, err := json.Marshal(identificationDoc)

	if err != nil {
		fmt.Println("[Func UpdateIndividualCustomer] Error json.Marshal ::", err)
		return err
	}

	header := Headers{
		ContentType: "application/json",
		XAuthToken:  f.token,
	}

	// _, err = CallFCAPI(f.token, "PUT", url, body, cfg)
	_, _, err = CallToFundConnext(cfg, "PUT", url, header, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("[Func UpdateIndividualCustomer] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

func (f *FundConnext) UpdatePartialIndividualCustomer(partialIndividualCustomerDocument PartialIndividualCustomerDocument) error {
	cfg := MakeAPICallerConfig(f)
	url := "/api/customer/individual"

	body, err := json.Marshal(partialIndividualCustomerDocument)

	if err != nil {
		fmt.Println("[Func UpdatePartialIndividualCustomer] Error json.Marshal ::", err)
		return err
	}

	header := Headers{
		ContentType: "application/json",
		XAuthToken:  f.token,
	}

	// _, err = CallFCAPI(f.token, "PATCH", url, body, cfg)
	_, _, err = CallToFundConnext(cfg, "PATCH", url, header, bytes.NewBuffer(body))

	if err != nil {
		return err
	}
	return nil
}

// TODO:
// func (f *FundConnext) UploadIndividualCustomerFile(fileType string, individualCustomerFile IndividualCustomerFile) error {
// 	cfg := MakeAPICallerConfig(f)
// 	url := fmt.Sprintf("/api/customer/individual/account/%s/upload", fileType)
// 	body, err := json.Marshal(individualCustomerFile)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = CallFCAPI(f.token, "PUT", url, body, cfg)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (f *FundConnext) UploadIndividualCustomerFile(fileType string, header Headers, body io.Reader) error {
	var err error

	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/customer/individual/%s/upload", fileType)
	header.XAuthToken = f.token
	// body, err := json.Marshal(individualCustomerFile)  /api/customer/individual/:fileType/upload

	// _, err = CallFCAPIV2(f.token, "POST", url, body, cfg, contentType)
	_, _, err = CallToFundConnext(cfg, "POST", url, header, body)

	if err != nil {
		return err
	}

	return nil
}
