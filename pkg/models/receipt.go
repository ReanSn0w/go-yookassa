package models

type Receipt struct {
	ID                   string                  `json:"id"`                                    // Идентификатор чека
	Type                 string                  `json:"type"`                                  // Тип чека в онлайн кассе
	PaymentID            string                  `json:"payment_id,omitempty"`                  // Идентификатор платежа
	RefoundID            string                  `json:"refound_id,omitempty"`                  // Идентификатор возврата
	Status               string                  `json:"status"`                                // Статус доставки данных для чека
	FiscalDocumentNumber string                  `json:"fiscal_document_number,omitempty"`      // Номер фискального документа
	FiscalStorageNumber  string                  `json:"fiscal_storage_number,omitempty"`       // Номер фискального накопителя
	FiscalAttribute      string                  `json:"fiscal_attribute,omitempty"`            // Признак фискального документа
	RegisteredAt         string                  `json:"registered_at,omitempty"`               // Дата и время формирования чека
	FiscalProviderID     string                  `json:"fiscal_provider_id,omitempty"`          // Идентификатор чека в онлайн-кассе.
	Items                []Item                  `json:"items"`                                 // Список товаров в чеке
	Settlements          []Settlement            `json:"settlements,omitempty"`                 // Перечень совершенных рассчетов
	OnBehalfOf           string                  `json:"on_behalf_of"`                          // Идентификатор магазина, от имени которого нужно отправить чек.
	TaxSystemCode        int                     `json:"tax_system_code,omitempty"`             // Система налогообложения магазина
	IndustryDetails      *RecieptIndustryDetails `json:"receipt_industry_details,omitempty"`    // Отраслевой реквизит предмета расчета
	OperationalDetails   *RecieptIndustryDetails `json:"receipt_operational_details,omitempty"` // Операционный реквизит чека
}

type ReceiptRequest struct {
	Type                string                  `json:"type"`                                  // Тип чека в онлайн-кассе.
	PaymentID           string                  `json:"payment_id,omitempty"`                  // Идентификатор платежа в сервисе
	RefoundID           string                  `json:"refound_id,omitempty"`                  // Идентификатор возврата в сервисе
	Customer            Customer                `json:"customer"`                              // Информация о пользователе
	Items               []Item                  `json:"items"`                                 // Список товаров в чеке
	Send                bool                    `json:"send"`                                  // Формирование чека в онлайн-кассе сразу после создания объекта чека.
	TaxSystemCode       int                     `json:"tax_system_code,omitempty"`             // Система налогообложения магазина
	AdditionalUserProps *AdditionalUserProps    `json:"additional_user_props,omitempty"`       // Дополнительный реквизит пользователя
	IndustryDetails     *RecieptIndustryDetails `json:"receipt_industry_details,omitempty"`    // Отраслевой реквизит предмета расчета
	OperationalDetails  *RecieptIndustryDetails `json:"receipt_operational_details,omitempty"` // Операционный реквизит чека
	Settlements         []Settlement            `json:"settlements,omitempty"`                 // Перечень совершенных рассчетов
	OnBehalfOf          string                  `json:"on_behalf_of"`                          // Идентификатор магазина, от имени которого нужно отправить чек.
}

type ReceiptListResponse struct {
	Type       string    `json:"type,omitempty"` // Тип ответа возвращаемого от сервиса
	Items      []Receipt `json:"items"`          // Элементы списка
	NextCursor string    `json:"next_cursor"`    // Курсор для запроса продолжения для списка
}

// MARK: - Дополнительные структуры для создания чека

// Дополнительный реквизит пользователя
type AdditionalUserProps struct {
	Name  string `json:"name"`  // Название
	Value string `json:"value"` // Значение
}
