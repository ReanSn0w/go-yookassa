package models

type Payout struct {
	ID                  string               `json:"id"`                             // Идентификатор выплаты
	Amount              Amount               `json:"amount"`                         // Сумма выплаты
	Status              string               `json:"status"`                         // Статус
	PayoutDestination   PaymentMethod        `json:"payout_destination"`             // Платежное средство продавца
	Description         string               `json:"description.omitempty"`          // Описание транзакции
	CreatedAt           string               `json:"created_at"`                     // Время создания выплаты.
	Deal                *DealItem            `json:"deal,omitempty"`                 // Сделка, в рамках которой нужно провести выплату.
	SelfEmployed        *SelfEmployedID      `json:"self_employed,omitempty"`        // Данные самозанятого, который получит выплату.
	Reciept             *FNSReciept          `json:"reciept,omitempty"`              // Данные чека, зарегистрированного в ФНС.
	CancellationDetails *CancellationDetails `json:"cancellation_details,omitempty"` // Комментарий к статусу canceled
	Metadata            Metadata             `json:"metadata,omitempty"`             // Любые дополнительные данные, которые нужны вам для работы
	Test                bool                 `json:"test"`                           // Признак тестовой операции
}

// Структура формирует тело запроса для создания платежа
type PayoutRequest struct {
	Amount                Amount          `json:"amount"`                            // Сумма выплаты
	PayoutDestinationData PaymentMethod   `json:"payout_destination_data,omitempty"` // Данные платежного средства, на которое нужно сделать выплату
	PayoutToken           string          `json:"payout_token,omitempty"`            // Токенизированные данные для выплаты
	PayoutMethodID        string          `json:"payment_method_id"`                 // Идентификатор сохраненного способа оплаты
	Description           string          `json:"description,omitempty"`             // Описание выплаты
	Deal                  *DealItem       `json:"deal,omitempty"`                    // Сделка, в рамках которой нужно провести выплату
	SelfEmployed          *SelfEmployedID `json:"self_employed,omitempty"`           // Данные самозанятого, который получит выплату
	Reciept               *FNSReciept     `json:"reciept,omitempty"`                 // Данные чека, зарегистрированного в ФНС
	PersonalData          []string        `json:"personal_data"`                     // Персональные данные получателя выплаты
	Metadata              Metadata        `json:"metadata,omitempty"`                // Любые дополнительные данные, которые нужны вам для работы
}

// MARK: - Дополнительные модели для структуры выплат

// Данные чека, зарегистрированного в ФНС.
type FNSReciept struct {
	ServiceName  string  `json:"service_name"`             // Описание услуги
	NpdReceiptID string  `json:"npd_receipt_id,omitempty"` // Идентификатор чека в сервисе
	Url          string  `json:"url,omitempty"`            // Ссылка на зарегистрированный чек
	Amount       *Amount `json:"amount"`                   // Сумма, указанная в чеке
}

type SelfEmployedID struct {
	ID string `json:"id"`
}
