package models

type Refound struct {
	ID                  string               `json:"id"`                             // Идентификатор возврата
	PaymentID           string               `json:"payment_id"`                     // Идентификатор платежа
	Status              string               `json:"status"`                         // Статус возврата
	CancellationDetails *CancellationDetails `json:"cancellation_details,omitempty"` // Комментарий к возврату
	CreatedAt           string               `json:"created_at"`                     // Время создания возврата
	Amount              Amount               `json:"amount"`                         // Сумма средств возвращенных пользователю
	Description         string               `json:"description,omitempty"`          // Описание для возврата
	Sources             []Source             `json:"sources,omitempty"`              // Данные об источниках возврата
	Deal                *DealItem            `json:"deal,omitempty"`                 // Данные о сделке в составе которой происходит возврат
}

type RefoundRequest struct {
	PaymentID   string    `json:"payment_id"`            // Идентификатор платежа
	Amount      Amount    `json:"amount"`                // Сумма средств возвращенных пользователю
	Description string    `json:"description,omitempty"` // Описание для возврата
	Sources     []Source  `json:"sources,omitempty"`     // Данные об источниках возврата
	Deal        *DealItem `json:"deal,omitempty"`        // Данные о сделке в составе которой происходит возврат
}

type RefoundListResponse struct {
	Type       string    `json:"type,omitempty"`
	Items      []Refound `json:"items"`
	NextCursor string    `json:"next_cursor"`
}

// MARK: - Дополнительные структуры для описания возврата

// Структура описывает источники, из которых берутся деньги для возврата
type Source struct {
	AccountID        string  `json:"account_id"`                   // Идентификатор аккаунта
	Amount           Amount  `json:"amount"`                       // Сумма возврата
	PaymentFeeAmount *Amount `json:"payment_fee_amount,omitempty"` // Коммисия, которую вы удержали при оплате
}
