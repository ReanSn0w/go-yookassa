package models

// Список банков участвующих в системе быстрых платежей
type SBPList struct {
	Type  string    `json:"type"`
	Items []SBPItem `json:"items"` // Список банков участников отсортированный по уменьшению идентификатора
}

// Элемент списка банков участников системы быстрых платежей
type SBPItem struct {
	ID   string `json:"bank_id"` // Идентификатор банка
	Name string `json:"name"`    // Название банка
}
