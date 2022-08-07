package yookassa

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

const (
	apiURL = "https://api.yookassa.ru/v3/"
)

func NewClient(shopID, key string) *Client {
	return &Client{
		shopID: shopID,
		key:    key,
		wc: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func NewClientWithHttp(shopID, key string, client *http.Client) *Client {
	return &Client{
		shopID: shopID,
		key:    key,
		wc:     client,
	}
}

type Client struct {
	wc      *http.Client
	shopID  string
	key     string
	idemKey *string
}

func (c *Client) WithIdempotenceKey(key string) *Client {
	return &Client{
		wc:      c.wc,
		shopID:  c.shopID,
		key:     c.key,
		idemKey: &key,
	}
}

func (c *Client) Payments() *Payments {
	return &Payments{yookassa: c}
}

func (c *Client) Payouts() *Payouts {
	return &Payouts{yookassa: c}
}

func (c *Client) Reciepts() *Reciepts {
	return &Reciepts{yookassa: c}
}

func (c *Client) Refounds() *Refounds {
	return &Refounds{yookassa: c}
}

func (c *Client) SelfEmployed() *SelfEmployed {
	return &SelfEmployed{yookassa: c}
}

func (c *Client) PersonalData() *PersonalData {
	return &PersonalData{yookassa: c}
}

func (c *Client) Deals() *Deals {
	return &Deals{yookassa: c}
}

func (c *Client) Webhooks(token string) *Webhooks {
	return &Webhooks{yookassa: c, token: token}
}

// Метод для получения актуального списка банков участвующих
// в системе быстрых платежей
//
// ref: https://yookassa.ru/developers/api#sbp_banks
func (c *Client) SBPList() (*models.SBPList, error) {
	resp := &models.SBPList{}
	err := c.request("GET", "sbp_banks", nil, resp)
	return resp, err
}

// Метод для получения информации о своем магазине или шлюзе
// ref: https://yookassa.ru/developers/api#me
func (c *Client) Me() (*models.PersonalInfo, error) {
	resp := &models.PersonalInfo{}
	err := c.request("GET", "me", nil, resp)
	return resp, err
}

// Информация о настройках магазина для сплитирования платежей
//
// Метод работает аналогично методу ME, за исключением того что в запрос передается
// идентификатор магазина продавца
// ref: https://yookassa.ru/developers/api#me
func (c *Client) SplitInfo(id string) (*models.PersonalInfo, error) {
	obj := map[string]string{"on_behalf_of": id}
	resp := &models.PersonalInfo{}
	err := c.request("GET", "me", obj, resp)
	return resp, err
}

func (c *Client) request(method string, relativeURL string, reqBody interface{}, obj interface{}) error {
	buffer := new(bytes.Buffer)
	if reqBody != nil {
		encoder := json.NewEncoder(buffer)
		err := encoder.Encode(reqBody)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, apiURL+relativeURL, buffer)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.shopID, c.key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", c.getIdempotentKey())

	resp, err := c.wc.Do(req)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(obj)
}

func (c *Client) requestByToken(token string, method string, relativeURL string, reqBody interface{}, obj interface{}) error {
	buffer := new(bytes.Buffer)
	if reqBody != nil {
		encoder := json.NewEncoder(buffer)
		err := encoder.Encode(reqBody)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, apiURL+relativeURL, buffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", c.getIdempotentKey())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.wc.Do(req)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(obj)
}

func (c *Client) getIdempotentKey() string {
	if c.idemKey != nil {
		return *c.idemKey
	}

	hash := md5.Sum([]byte(time.Now().String()))
	return string(hash[:8])
}
