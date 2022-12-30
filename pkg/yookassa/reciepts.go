package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

// Структура содержит методы для работы с чеками
//
// ref: https://yookassa.ru/developers/api#refund
type Reciepts struct {
	yookassa *Client
}

// Создание нового чека
//
// В ответ на запрос придет объект чека в актуальном статусе.
func (r *Reciepts) Create(req models.ReceiptRequest) (*models.Receipt, error) {
	resp := &models.Receipt{}
	err := r.yookassa.request(http.MethodPost, "receipts", req, resp)
	return resp, err
}

// Получение списока чеков
//
// В ответ на запрос вернется список чеков с учетом переданных параметров.
// В списке будет информация о чеках, созданных за последние 3 года.
// Список будет отсортирован по времени создания чеков в порядке убывания.
func (r *Reciepts) List(req models.ListRequest) (*models.ReceiptListResponse, error) {
	resp := &models.ReceiptListResponse{}
	err := r.yookassa.request(http.MethodGet, fmt.Sprintf("reciepts?%s", req.Query()), nil, resp)
	return resp, err
}

// Запрос информации о чеке
//
// В ответ на запрос придет объект чека в актуальном статусе.
func (r *Reciepts) Info(id string) (*models.Receipt, error) {
	resp := &models.Receipt{}
	err := r.yookassa.request(http.MethodGet, fmt.Sprintf("reciepts/%s", id), nil, resp)
	return resp, err
}
