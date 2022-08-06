package models

const (
	PersonalDataStatusWaiting  personalDataStatus = "waiting_for_operation"
	PersonalDataStatusActive   personalDataStatus = "active"
	PersonalDataStatusCanceled personalDataStatus = "canceled"
)

type PersonalInfo struct {
	AccountID            string          `json:"acoount_id"`
	Status               string          `json:"status"`
	Test                 bool            `json:"test"`
	FiscalizationEnabled bool            `json:"fiscalization_enabled,omitempty"`
	PaymentMethods       []PaymentMethod `json:"payment_methods,omitempty"`
	Itn                  string          `json:"itn,omitempty"`
	PayoutMethods        []string        `json:"payout_methods,omitempty"`
	Name                 string          `json:"name,omitempty"`
	PayoutBalance        *Amount         `json:"payout_balance,omitempty"`
}

type PersonalData struct {
	ID                  string               `json:"id"`                             // Идентификатор персональных данных
	Type                string               `json:"type"`                           // Тип персональных данных
	Status              personalDataStatus   `json:"status"`                         // Статус персональных данных
	CancellationDetails *CancellationDetails `json:"cancellation_details,omitempty"` // Комментарий к статусу canceled
	CreatedAt           string               `json:"created_at"`                     // Время создания персональных данных
	ExpiresAt           string               `json:"expires_at,omitempty"`           // Срок жизни объекта персональных данных
	Metadata            Metadata             `json:"metadata,omitempty"`             // Произвольные дополнительные данные
}

// Форма для запроса на сохранение персональных данных в сервисе
type PersonalDataRequest struct {
	Type       string   `json:"type"`                  // Тип персональных данных
	FirstName  string   `json:"first_name"`            // Имя пользователя
	MiddleName string   `json:"middle_name,omitempty"` // Отчество пользователя
	LastName   string   `json:"last_name"`             // Фамилия пользователя
	Metadata   Metadata `json:"metadata,omitempty"`    // Произвольные дополнительные данные
}

// Статус персональных данных
//
// возможные состояния описаны в константах
type personalDataStatus string
