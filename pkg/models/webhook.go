package models

// Структура Webhook содержит информацию о подписке на одно событие
type Webhook struct {
	ID    string `json:"id"`    // Идентификатор webhook
	Event string `json:"event"` // Событие, о котором уведомляет ЮKassa
	URL   string `json:"url"`   // URL, на который ЮKassa отправляет уведомления
}

// Структура содержит параметры для запроса на создание Webhook
type WebhookRequest struct {
	Event string `json:"event"` // Событие, о котором уведомляет ЮKassa
	URL   string `json:"url"`   // URL, на который ЮKassa отправляет уведомления
}

type WebhookListResponse struct {
	Type  string    `json:"type"`
	Items []Webhook `json:"items"`
}

type PaymentWebhook struct {
	Type    string  `json:"type"`
	Event   string  `json:"event"`
	Payment Payment `json:"object"`
}

type DealWebhook struct {
	Type  string `json:"type"`
	Event string `json:"event"`
	Deal  Deal   `json:"object"`
}
