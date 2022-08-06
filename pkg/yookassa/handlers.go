package yookassa

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
)

var (
	yookassaSubnets = []string{
		"185.71.76.0/27",
		"185.71.77.0/27",
		"77.75.153.0/25",
		"77.75.156.11/32",
		"77.75.156.35/32",
		"77.75.154.128/25",
		"2a02:5180::/32",
	}

	subnets = []*net.IPNet{}
)

func init() {
	for _, item := range yookassaSubnets {
		_, subnet, err := net.ParseCIDR(item)
		if err != nil {
			log.Fatalln(err)
		}

		subnets = append(subnets, subnet)
	}
}

func NewPaymentWebhookHandler(checkIP bool, webhook func(*models.PaymentWebhook)) http.HandlerFunc {
	return baseWebhook(checkIP, func(r io.Reader) {
		decoder := json.NewDecoder(r)

		obj := &models.PaymentWebhook{}
		err := decoder.Decode(obj)
		if err != nil {
			log.Println(err)
			return
		}

		webhook(obj)
	})
}

func NewDealWebhookHandler(checkIP bool, webhook func(*models.DealWebhook)) http.HandlerFunc {
	return baseWebhook(checkIP, func(r io.Reader) {
		decoder := json.NewDecoder(r)

		obj := &models.DealWebhook{}
		err := decoder.Decode(obj)
		if err != nil {
			log.Println(err)
			return
		}

		webhook(obj)
	})
}

func baseWebhook(checkIP bool, webhook func(io.Reader)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверка IP адреса на вхождение в пул сервиса
		if checkIP {
			ipString := strings.Split(r.RemoteAddr, ":")[0]
			ip := net.ParseIP(ipString)
			err := CheckIP(ip)
			if err != nil {
				log.Println(err)
				return
			}
		}

		webhook(r.Body)
		r.Body.Close()
		w.WriteHeader(http.StatusOK)
	}
}

// Функция проверяет IP адрес на вхождение в пул адресов сервиса
func CheckIP(ip net.IP) error {
	if ip == nil {
		return errors.New("ip is nil")
	}

	checked := false

	for _, subnet := range subnets {
		if subnet.Contains(ip) {
			checked = true
			break
		}
	}

	if !checked {
		return errors.New("yookassa webhook ignored by ip")
	}

	return nil
}
