package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

// Структура с методами для управления сделками
//
// ref: https://yookassa.ru/developers/api#deal
type Deals struct {
	yookassa *Client
}

// Создание новой сделки
//
// В ответ на запрос придет объект сделки в актуальном статусе.
func (d *Deals) Create(req models.DealRequest) (*models.Deal, error) {
	resp := &models.Deal{}
	err := d.yookassa.request(http.MethodPost, "deals", req, resp)
	return resp, err
}

// Запрос списка сделок
//
// В ответ на запрос вернется список сделок с учетом переданных параметров.
// В списке будет информация о сделках, созданных за последние 3 года.
// Список будет отсортирован по времени создания сделок в порядке убывания.
func (d *Deals) List(req models.ListRequest) (*models.DealListResponse, error) {
	resp := &models.DealListResponse{}
	err := d.yookassa.request(http.MethodGet, fmt.Sprintf("deals?%s", req.Query()), nil, resp)
	return resp, err
}

// Запрос информации о сделке
//
// В ответ на запрос придет объект сделки в актуальном статусе.
func (d *Deals) Info(id string) (*models.Deal, error) {
	resp := &models.Deal{}
	err := d.yookassa.request(http.MethodGet, fmt.Sprintf("deals/%s", id), nil, resp)
	return resp, err
}
