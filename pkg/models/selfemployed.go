package models

const (
	SelfEmployedStatusPending   selfEmployedStatus = "pending"
	SelfEmployedStatusConfirmed selfEmployedStatus = "confirmed"
	SelfEmployedStatusCanceled  selfEmployedStatus = "canceled"
)

type SelfEmployed struct {
	ID           string             `json:"id"`                     // Идентификатор самозанятого
	Status       selfEmployedStatus `json:"status"`                 // Статус самозанятого
	CreatedAt    string             `json:"created_at"`             // Время создания объекта
	Itn          string             `json:"itn,omitempty"`          // ИНН самозанятого
	Phone        string             `json:"phone,omitempty"`        // Телефон самозанятого
	Confirmation *Confirmation      `json:"confirmation,omitempty"` // Сценарий проверки
	Description  string             `json:"description,omitempty"`  // Описание самозанятого (не более 128 символов)
	Metadata     Metadata           `json:"metadata,omitempty"`     // Любые дополнительные данные
	Test         bool               `json:"test"`                   // Признак тестовой сделки
}

type SelfEmployedRequest struct {
	Itn          string        `json:"itn,omitempty"`          // ИНН самозанятого
	Phone        string        `json:"phone,omitempty"`        // Телефон самозанятого
	Confirmation *Confirmation `json:"confirmation,omitempty"` // Сценарий проверки
	Description  string        `json:"description,omitempty"`  // Описание самозанятого (не более 128 символов)
	Metadata     Metadata      `json:"metadata,omitempty"`     // Любые дополнительные данные
}

type selfEmployedStatus string
