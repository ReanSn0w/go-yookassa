package yookassa

import (
	"fmt"
	"net/http"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

type Webhooks struct {
	yookassa *Client
	token    string
}

// Создание webhook
//
// Запрос позволяет подписаться на уведомления об основных событиях (например, переход платежа в статус succeeded).
//
// Если вы хотите получать уведомления о нескольких событиях,
// вам нужно для каждого из них создать свой webhook.
// Для каждого OAuth-токена нужно создавать свой набор webhook.
func (w *Webhooks) Create(req models.WebhookRequest) (*models.Webhook, error) {
	resp := &models.Webhook{}
	err := w.yookassa.requestByToken(w.token, http.MethodPost, "webhooks", req, resp)
	return resp, err
}

// Список созданных webhook
//
// Запрос позволяет узнать, какие webhook есть для переданного OAuth-токена.
//
// В ответ на запрос придет актуальный список объектов webhook.
func (w *Webhooks) List() (*models.WebhookListResponse, error) {
	resp := &models.WebhookListResponse{}
	err := w.yookassa.requestByToken(w.token, http.MethodGet, "webhooks", nil, resp)
	return resp, err
}

// Удаление webhook
//
// Запрос позволяет отписаться от уведомлений о событии для переданного OAuth-токена.
// Чтобы удалить webhook, вам нужно передать в запросе его идентификатор.
func (w *Webhooks) Delete(id string) error {
	return w.yookassa.requestByToken(w.token, http.MethodDelete, fmt.Sprintf("webhooks/%s", id), nil, nil)
}
