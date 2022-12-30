package models

// MARK: - Метод для оплаты
// Набор структур для описания метода оплаты

// Структура для описания метода, используемого для платежа
type PaymentMethod struct {
	Type             string            `json:"type"` // Тип метода оплаты
	ID               string            `json:"id"`
	Saved            bool              `json:"saved"`
	Title            string            `json:"title,omitempty"`
	Login            string            `json:"login,omitempty"`
	Card             *Card             `json:"card,omitempty"`
	PayerBankDetails *PayerBankDetails `json:"payer_bank_details,omitempty"`
	PaymentPurpose   string            `json:"payment_purpose,omitempty"`
	VatData          *VATData          `json:"vat_data,omitempty"`
}

// Стркутура для описания данных карты
type Card struct {
	Last4         string `json:"last4"`
	ExpiryYear    string `json:"expiry_year"`
	ExpityMonth   string `json:"expiry_month"`
	CardType      string `json:"card_type"`
	First6        string `json:"first6,omitempty"`
	IssuerCountry string `json:"issuer_country,omitempty"`
	IssuerName    string `json:"issuer_name,omitempty"`
	Source        string `json:"source,omitempty"`
}

// Структура для описания данных о банке
type PayerBankDetails struct {
	FullName   string `json:"full_name"`
	ShortName  string `json:"short_name"`
	Address    string `json:"address"`
	Inn        string `json:"inn"`
	BankName   string `json:"bank_name"`
	BankBranch string `json:"bank_branch"`
	BankBik    string `json:"bank_nik"`
	Account    string `json:"account"`
	Kpp        string `json:"kpp,omitempty"`
}

// Структура для рассчета налога
type VATData struct {
	Type   string  `json:"type"`
	Amount *Amount `json:"amount,omitempty"`
	Rate   string  `json:"rate,omitempty"`
}

type CancellationDetails struct {
	Party  string `json:"party"`
	Reason string `json:"reason"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type AuthorizationDetails struct {
	Rrn          string       `json:"rrn,omitempty"`
	AuthCode     string       `json:"auth_code,omitempty"`
	ThreeDSecure ThreeDSecure `json:"three_d_secure"`
}

type ThreeDSecure struct {
	Applied bool `json:"applied"`
}

type Confirmation struct {
	Type              string `json:"type,omitempty"`
	ConfirmationToken string `json:"confirmation_token,omitempty"`
	ConfirmationURL   string `json:"confirmation_url,omitempty"`
	ConfirmationData  string `json:"confirmation_data,omitempty"`
	Enforce           bool   `json:"enforce,omitempty"`
	ReturnURl         string `json:"return_url,omitempty"`
}

// Структура для описания информации о пользователе
type Customer struct {
	FullName string `json:"full_name,omitempty"`
	INN      string `json:"inn,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// Любые дополнительные данные, которые нужны вам для работы
// (например, ваш внутренний идентификатор заказа).
//
// Ограничения:
// - максимум 16 ключей,
// - имя ключа не больше 32 символов,
// - значение ключа не больше 512 символов
type Metadata map[string]string

// MARK: - Item
// Структура для описания элементов в чеке

// Структура для описания элемента в чеке
type Item struct {
	Description              string            `json:"description"`
	Quantity                 string            `json:"quantity"`
	Measure                  string            `json:"measure,omitempty"`
	MarkQuantity             *MarkQuantity     `json:"mark_quantity,omitempty"`
	Amount                   Amount            `json:"amount"`
	VATCode                  float64           `json:"vat_code"`
	PaymentSubject           string            `json:"payment_subject,omitempty"`
	PaymentMode              string            `json:"payment_mode,omitempty"`
	CountryOfOriginCode      string            `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber string            `json:"customs_declaration_number,omitempty"`
	Exsize                   string            `json:"exsize,omitempty"`
	ProductCode              string            `json:"product_code,omitempty"`
	MarkCodeInfo             *MarkCodeInfo     `json:"mark_code_info,omitempty"`
	MarkMode                 string            `json:"mark_mode,omitempty"`
	IndustryDetails          []IndustryDetails `json:"payment_subject_industry_details,omitempty"`
}

type MarkQuantity struct {
	Numerator   int `json:"numerator,omitempty"`
	Denominator int `json:"denominator,omitempty"`
}

type MarkCodeInfo struct {
	MarkCodeRaw string `json:"mark_code_raw,omitempty"`
	Unknown     string `json:"unknown,omitempty"`
	Ean8        string `json:"ean_8,omitempty"`
	Ean13       string `json:"ean_13,omitempty"`
	Itf14       string `json:"itf_14,omitempty"`
	Gs10        string `json:"gs_10,omitempty"`
	Gs1m        string `json:"gs_1m,omitempty"`
	Short       string `json:"short,omitempty"`
	Fur         string `json:"fur,omitempty"`
	Egais20     string `json:"egais_20,omitempty"`
	Egais30     string `json:"egais_30,omitempty"`
}

type IndustryDetails struct {
	FederalID      string `json:"federal_id"`
	DocumentData   string `json:"document_date"`
	DocumentNumber string `json:"document_number"`
	Value          string `json:"value"`
}

// Структура для описания сделки
type DealItem struct {
	ID          string       `json:"id"`
	Settlements []Settlement `json:"settlements"`
}

type RecieptIndustryDetails struct {
	FederalID      string `json:"federal_id"`
	DocumentDate   string `json:"document_date"`
	DocumentNumber string `json:"document_number"`
	Value          string `json:"value"`
}

type Settlement struct {
	Type   string `json:"type"`
	Amount Amount `json:"amount"`
}
