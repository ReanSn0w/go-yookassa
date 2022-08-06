package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

// Структура содержит методы для работы с выплатами
//
// ref: https://yookassa.ru/developers/api#payout
type Payouts struct {
	yookassa *Client
}

// Создание новой выплаты
//
// В ответ на запрос придет объект выплаты в актуальном статусе.
func (p *Payouts) Create(req models.PayoutRequest) (*models.Payout, error) {
	resp := &models.Payout{}
	err := p.yookassa.request(http.MethodPost, "payouts", req, resp)
	return resp, err
}

// Запрос информации о выплате по ее ID
//
// В ответ на запрос придет объект выплаты в актуальном статусе.
func (p *Payouts) Info(id string) (*models.Payout, error) {
	resp := &models.Payout{}
	err := p.yookassa.request(http.MethodGet, fmt.Sprintf("payouts/%s", id), nil, resp)
	return resp, err
}
