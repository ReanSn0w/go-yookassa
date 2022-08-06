package models

const (
	DealTypeSafeDeal dealType = "safe_deal"

	DealStatusOpened dealStatus = "opened"
	DealStatusClosed dealStatus = "closed"

	FeeMomentSucceeded  feeMoment = "payment_succeeded"
	FeeMomentDealClosed feeMoment = "deal_closed"
)

type Deal struct {
	ID            string     `json:"id"`                    // Идентификатор сделки
	Type          dealType   `json:"type"`                  // Тип сделки
	FeeMoment     feeMoment  `json:"fee_moment"`            // Момент перечисления вам вознаграждения платформы.
	Description   string     `json:"description,omitempty"` // Описание сделки (не более 128 символов).
	Balance       Amount     `json:"balance"`               // Баланс сделки
	PayoutBalance Amount     `json:"payout_balance"`        // Сумма вознаграждения продавца
	Status        dealStatus `json:"status"`                // Статус сделки
	CreatedAt     string     `json:"created_at"`            // Дата создания сделки (UTC, ISO8601)
	ExpiresAt     string     `json:"expires_at"`            // Время автоматического закрытия сделки.
	Metadata      Metadata   `json:"metadata,omitempty"`    // Любые дополнительные данные
	Test          bool       `json:"test"`                  // Признак тестовой сделки
}

type DealRequest struct {
	Type        dealType  `json:"type"`                  // Тип сделки
	FeeMoment   feeMoment `json:"fee_moment"`            // Момент перечисления вам вознаграждения платформы.
	Metadata    Metadata  `json:"metadata,omitempty"`    // Любые дополнительные данные
	Description string    `json:"description,omitempty"` // Описание сделки (не более 128 символов).
}

type DealListResponse struct {
	Type   string `json:"type"`
	Items  []Deal `json:"deal"`
	Cursor string `json:"next_cursor"`
}

// MARK: - Типы для констант

// Тип сделки
type dealType string

type dealStatus string

// Момент оплаты
type feeMoment string
