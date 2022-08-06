package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

// Структура предоставляет доступ к методам управления персональными данными
//
// ref: https://yookassa.ru/developers/api#personal_data
type PersonalData struct {
	yookassa *Client
}

// Создание персональных данных
//
// Используйте этот запрос, чтобы создать в ЮKassa объект персональных данных.
// В запросе необходимо передать фамилию, имя, отчество пользователя и указать,
// с какой целью эти данные будут использоваться.
//
// В ответ на запрос придет объект персональных данных в актуальном статусе.
func (pd *PersonalData) Create(req models.PersonalDataRequest) (*models.PersonalData, error) {
	resp := &models.PersonalData{}
	err := pd.yookassa.request(http.MethodPost, "personal_data", req, resp)
	return resp, err
}

// Информация о персональных данных
//
// С помощью этого запроса вы можете получить информацию
// о текущем статусе объекта персональных данных по его уникальному идентификатору.
//
// В ответ на запрос придет объект персональных данных в актуальном статусе.
func (pd *PersonalData) Info(id string) (*models.PersonalData, error) {
	resp := &models.PersonalData{}
	err := pd.yookassa.request(http.MethodGet, fmt.Sprintf("personal_data/%s", id), nil, resp)
	return resp, err
}
