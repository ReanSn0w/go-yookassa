package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

// Структура содержит методы для работы с возвратами
//
// ref: https://yookassa.ru/developers/api#refund
type Refounds struct {
	yookassa *Client
}

// Создание возврата
//
// В ответ на запрос придет объект возврата в актуальном статусе.
func (r *Refounds) Create(req models.RefoundRequest) (*models.Refound, error) {
	resp := &models.Refound{}
	err := r.yookassa.request(http.MethodPost, "refunds", req, resp)
	return resp, err
}

// Запрос списка возвратов
//
// В ответ на запрос вернется список возвратов с учетом переданных параметров.
// В списке будет информация о возвратах, созданных за последние 3 года.
// Список будет отсортирован по времени создания возвратов в порядке убывания.
func (r *Refounds) List(req models.ListRequest) (*models.RefoundListResponse, error) {
	resp := &models.RefoundListResponse{}
	err := r.yookassa.request(http.MethodGet, fmt.Sprintf("refounds?%s", req.Query()), nil, resp)
	return resp, err
}

// Запрос информации о возврате
//
// В ответ на запрос придет объект возврата в актуальном статусе.
func (r *Refounds) Info(id string) (*models.Refound, error) {
	resp := &models.Refound{}
	err := r.yookassa.request(http.MethodGet, fmt.Sprintf("refounds/%s", id), nil, resp)
	return resp, err
}
