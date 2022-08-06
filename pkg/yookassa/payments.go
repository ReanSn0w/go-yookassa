package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

// Структура содержит методы для работы с платежами
//
// ref: https://yookassa.ru/developers/api#payment
type Payments struct {
	yookassa *Client
}

// Создание объекта платежа
//
// В ответ на запрос придет объект платежа в актуальном статусе.
func (p *Payments) Create(req models.PaymentRequest) (*models.Payment, error) {
	resp := &models.Payment{}
	err := p.yookassa.request(http.MethodPost, "payments", req, resp)
	return resp, err
}

// Запрос списка платежей
//
// В ответ на запрос вернется список платежей с учетом переданных параметров.
// В списке будет информация о платежах, созданных за последние 3 года.
// Список будет отсортирован по времени создания платежей в порядке убывания.
func (p *Payments) List(req models.ListRequest) (*models.PaymentListResponse, error) {
	resp := &models.PaymentListResponse{}
	err := p.yookassa.request(http.MethodGet, fmt.Sprintf("payments?%v", req.Query()), nil, resp)
	return resp, err
}

// Запрос для получения информации о платеже
//
// В ответ на запрос придет объект платежа в актуальном статусе.
func (p *Payments) Info(id string) (*models.Payment, error) {
	resp := &models.Payment{}
	err := p.yookassa.request(http.MethodGet, fmt.Sprintf("payments/%s", id), nil, resp)
	return resp, err
}

// Запрос для подтверждения платежа
//
// В ответ на запрос придет объект платежа в актуальном статусе.
func (p *Payments) Capture(id string, req models.CaptureRequest) (*models.Payment, error) {
	resp := &models.Payment{}
	err := p.yookassa.request(http.MethodPost, fmt.Sprintf("payments/%s/capture", id), req, resp)
	return resp, err
}

// Запрос отмены платежа
//
// В ответ на запрос придет объект платежа в актуальном статусе.
func (p *Payments) Cancel(id string) (*models.Payment, error) {
	resp := &models.Payment{}
	err := p.yookassa.request(http.MethodPost, fmt.Sprintf("payments/%s/cancel", id), nil, resp)
	return resp, err
}
