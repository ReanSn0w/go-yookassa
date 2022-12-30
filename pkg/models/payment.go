package models

// Информация о платеже возвращаемая сервисом
type Payment struct {
	ID                   string                `json:"id"`
	Status               string                `json:"status"`
	Paid                 bool                  `json:"paid"`
	Amount               Amount                `json:"amount"`
	IncomeAmount         *Amount               `json:"income_amount,omitempty"`
	Description          string                `json:"description,omitempty"`
	Recipient            Recipient             `json:"recipient"`
	PaymentMethod        *PaymentMethod        `json:"payment_method,omitempty"`
	CapturedAt           string                `json:"captured_at,omitempty"`
	CreatedAt            string                `json:"created_at"`
	ExpiresAt            string                `json:"expires_at,omitempty"`
	Confirmation         *Confirmation         `json:"confirmation,omitempty"`
	Test                 bool                  `json:"test"`
	RefoundedAmount      *Amount               `json:"refounded_amount,omitempty"`
	Payed                bool                  `json:"payed"`
	Refundable           bool                  `json:"refundable"`
	ReceiptRegistration  string                `json:"receipt_registration,omitempty"`
	Metadata             Metadata              `json:"metadata,omitempty"`
	CancellationDetails  *CancellationDetails  `json:"cancellation_details,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details"`
	Transfers            []Transfer            `json:"transfers,omitempty"`
	Deal                 *DealItem             `json:"deal,omitempty"`
	MerchantCustomerID   string                `json:"merchant_customer_id,omitempty"`
}

// Структура для описания запроса создания платежа
type PaymentRequest struct {
	Amount             Amount         `json:"amount"`
	Description        string         `json:"description,omitempty"`
	Reciept            *ReceiptInfo   `json:"receipt,omitempty"`
	Recipient          *Recipient     `json:"recepient,omitempty"`
	PaymentToken       string         `json:"payment_token,omitempty"`
	PaymentMethodID    string         `json:"payment_method_id,omitempty"`
	PaymentMethodData  *PaymentMethod `json:"payment_method_data,omitempty"`
	Confirmation       *Confirmation  `json:"confirmation,omitempty"`
	SavePaymentMethod  bool           `json:"save_payment_method,omitempty"`
	ClientID           string         `json:"client_id,omitempty"`
	Metadata           *Metadata      `json:"metadata,omitempty"`
	Airline            *Airline       `json:"airline,omitempty"`
	Transfers          []Transfer     `json:"transfers,omitempty"`
	Deal               *DealItem      `json:"deal,omitempty"`
	MerchantCustomerID string         `json:"merchant_customer_id,omitempty"`
	Capture            bool           `json:"capture,omitempty"`
}

// Ответ на запрос получения списка платежей
type PaymentListResponse struct {
	Type       string    `json:"type,omitempty"`
	Items      []Payment `json:"items"`
	NextCursor string    `json:"next_cursor"`
}

// Структура для описания запроса на подтвержение проведения платежа
type CaptureRequest struct {
	Amount    *Amount    `json:"amount,omitempty"`
	Reciept   *Receipt   `json:"receipt,omitempty"`
	Airline   *Airline   `json:"airline,omitempty"`
	Transfers []Transfer `json:"transfers,omitempty"`
	Deal      *Deal      `json:"deal,omitempty"`
}

// MARK: - Дополнительные структуры для работы с платежами

type Airline struct {
	TicketNumber     string      `json:"ticket_number,omitempty"`
	BookingReference string      `json:"booking_reference,omitempty"`
	Passengers       []Passenger `json:"passengers,omitempty"`
	Legs             []Leg       `json:"legs,omitempty"`
}

type Leg struct {
	DepartureAirport   string `json:"departure_airport"`
	DestinationAirport string `json:"destination_airport"`
	DepartureDate      string `json:"departure_date"`
	CarrierCode        string `json:"carrier_code,omitempty"`
}

type Passenger struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ReceiptInfo struct {
	Customer                  *Customer                  `json:"customer,omitempty"`
	Items                     []Item                     `json:"items,omitempty"`
	TaxSystemCode             int                        `json:"tax_system_code,omitempty"`
	RecieptIndustryDetails    []RecieptIndustryDetails   `json:"receipt_industry_details,omitempty"`
	ReceiptOperationalDetails *ReceiptOperationalDetails `json:"receipt_operational_details,omitempty"`
}

type Recipient struct {
	AccountID string `json:"account_id"`
	GatewayID string `json:"gateway_id"`
}

type ReceiptOperationalDetails struct {
	OperationID string `json:"operation_id"`
	Value       string `json:"value"`
	CreatedAt   string `json:"created_at"`
}

type Transfer struct {
	AccountID         string   `json:"account_id"`
	Amount            Amount   `json:"amount"`
	Status            string   `json:"status"`
	PlatformFeeAmount *Amount  `json:"platform_fee_amount,omitempty"`
	Description       string   `json:"description,omitempty"`
	Metadata          Metadata `json:"metadata,omitempty"`
}
