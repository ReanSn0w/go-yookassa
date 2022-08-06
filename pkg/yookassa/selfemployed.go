package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

// Структура предоставляет доступ к методам управления самозанятым
//
// ref: https://yookassa.ru/developers/api#self_employed
type SelfEmployed struct {
	yookassa *Client
}

// Создание нового самозанятого
//
// Используйте этот запрос, чтобы создать в ЮKassa объект самозанятого.
// В запросе необходимо передать ИНН или телефон самозанятого для идентификации в сервисе Мой налог,
// сценарий подтверждения пользователем заявки ЮMoney на получение прав для регистрации чеков и описание самозанятого.
//
// Идентификатор созданного объекта самозанятого необходимо использовать в запросе на проведение выплаты.
//
// В ответ на запрос придет объект самозанятого в актуальном статусе.
func (se *SelfEmployed) Create(req models.SelfEmployedRequest) (*models.SelfEmployed, error) {
	resp := &models.SelfEmployed{}
	err := se.yookassa.request(http.MethodPost, "self_employed", req, resp)
	return resp, err
}

// Информация о самозанятом
//
// С помощью этого запроса вы можете получить информацию о текущем статусе самозанятого
// по его уникальному идентификатору.
//
// В ответ на запрос придет объект самозанятого в актуальном статусе.
func (se *SelfEmployed) Info(id string) (*models.SelfEmployed, error) {
	resp := &models.SelfEmployed{}
	err := se.yookassa.request(http.MethodGet, fmt.Sprintf("self_employed/%s", id), nil, resp)
	return resp, err
}
