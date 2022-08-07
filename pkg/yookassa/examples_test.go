package yookassa_test

import (
	"fmt"

	"github.com/ReanSn0w/go-yookassa/pkg/models"
	"github.com/ReanSn0w/go-yookassa/pkg/yookassa"
)

const (
	ycID    = ""
	ycKey   = ""
	ycToken = ""
)

// MARK: - Payments

func ExamplePayments_Create() {
	p := yookassa.NewClient(ycID, ycKey).Payments()

	item, _ := p.Create(models.PaymentRequest{
		Amount: models.Amount{
			Value:    "2.00",
			Currency: "RUB",
		},
		PaymentMethodData: &models.PaymentMethod{
			Type: "bank_card",
		},
		Confirmation: &models.Confirmation{
			Type:      "redirect",
			ReturnURl: "https://www.merchant-website.com/return_url",
		},
		Description: "Заказ №72",
	})

	fmt.Println(item.ID)
}

func ExamplePayments_List() {
	p := yookassa.NewClient(ycID, ycKey).Payments()

	list, _ := p.List(models.ListRequest{
		Limit: 10,
	})

	for _, item := range list.Items {
		fmt.Println(item.ID)
	}
}

func ExamplePayments_Info() {
	p := yookassa.NewClient(ycID, ycKey).Payments()

	item, _ := p.Info("215d8da0-000f-50be-b000-0003308c89be")

	fmt.Println(item.ID)
}

func ExamplePayments_Capture() {
	p := yookassa.NewClient(ycID, ycKey).Payments()

	item, _ := p.Capture("215d8da0-000f-50be-b000-0003308c89be", models.CaptureRequest{
		Amount: &models.Amount{
			Value:    "2.00",
			Currency: "RUB",
		},
	})

	fmt.Println(item.ID)
}

func ExamplePayments_Cancel() {
	p := yookassa.NewClient(ycID, ycKey).Payments()

	item, _ := p.Cancel("215d8da0-000f-50be-b000-0003308c89be")

	fmt.Println(item.ID)
}

// MARK: - Refounds

func ExampleRefounds_Create() {
	r := yookassa.NewClient(ycID, ycKey).Refounds()

	item, _ := r.Create(models.RefoundRequest{
		Amount: models.Amount{
			Value:    "2.00",
			Currency: "RUB",
		},
		PaymentID: "215d8da0-000f-50be-b000-0003308c89be",
	})

	fmt.Println(item.ID)
}

func ExampleRefounds_List() {
	r := yookassa.NewClient(ycID, ycKey).Refounds()

	list, _ := r.List(models.ListRequest{
		Limit: 10,
	})

	for _, item := range list.Items {
		fmt.Println(item.ID)
	}
}

func ExampleRefounds_Info() {
	r := yookassa.NewClient(ycID, ycKey).Refounds()

	item, _ := r.Info("215d8da0-000f-50be-b000-0003308c89be")

	fmt.Println(item.ID)
}

// MARK: - Reciepts

func ExampleReciepts_Create() {
	r := yookassa.NewClient(ycID, ycKey).Reciepts()

	item, _ := r.Create(models.RecieptRequest{
		Customer: models.Customer{
			FullName: "Ivanov Ivan Ivanovich",
			Email:    "email@email.ru",
			Phone:    "79211234567",
			INN:      "6321341814",
		},
		PaymentID: "24b94598-000f-5000-9000-1b68e7b15f3f",
		Type:      "payment",
		Send:      true,
		Items: []models.Item{
			{
				Description: "Наименование товара 1",
				Quantity:    "1",
				Amount: models.Amount{
					Value:    "14000.00",
					Currency: "RUB",
				},
				VATCode:             2,
				PaymentMode:         "full_payment",
				PaymentSubject:      "commodity",
				CountryOfOriginCode: "CN",
			},
		},
		Settlements: []models.Settlement{
			{
				Type: "prepayment",
				Amount: models.Amount{
					Value:    "8000.00",
					Currency: "RUB",
				},
			},
			{
				Type: "prepayment",
				Amount: models.Amount{
					Value:    "6000.00",
					Currency: "RUB",
				},
			},
		},
	})

	fmt.Sprintln(item.ID)
}

func ExampleReciepts_List() {
	r := yookassa.NewClient(ycID, ycKey).Reciepts()

	list, _ := r.List(models.ListRequest{Limit: 10})

	for _, item := range list.Items {
		fmt.Println(item.ID)
	}
}

func ExampleReciepts_Info() {
	r := yookassa.NewClient(ycID, ycKey).Reciepts()

	info, _ := r.Info("rt-2da5c87d-0384-50e8-a7f3-8d5646dd9e10")

	fmt.Sprintln(info.ID)
}

// MARK: - Payouts

func ExamplePayouts_Create() {
	p := yookassa.NewClient(ycID, ycKey).Payouts()

	item, _ := p.Create(models.PayoutRequest{
		Amount: models.Amount{
			Value:    "320.00",
			Currency: "RUB",
		},
		PayoutToken: "<Синоним банковской карты>",
		Description: "Выплата по заказу №37",
		Metadata: models.Metadata{
			"order_id": "37",
		},
		Deal: &models.DealItem{
			ID: "dl-285e5ee7-0022-5000-8000-01516a44b147",
		},
	})

	fmt.Println(item.ID)
}

func ExamplePayouts_Info() {
	p := yookassa.NewClient(ycID, ycKey).Payouts()

	item, _ := p.Info("po-285ec15d-0003-5000-a000-08d1bec7dade")

	fmt.Println(item.ID)
}

// MARK: - SelfEmployed

func ExampleSelfEmployed_Create() {
	se := yookassa.NewClient(ycID, ycKey).SelfEmployed()

	item, _ := se.Create(models.SelfEmployedRequest{
		Itn:         "123456789012",
		Description: "Курьер 001",
		Confirmation: &models.Confirmation{
			Type: "redirect",
		},
		Metadata: models.Metadata{
			"courier_id": "001",
		},
	})

	fmt.Printf(item.ID)
}

func ExampleSelfEmployed_Info() {
	se := yookassa.NewClient(ycID, ycKey).SelfEmployed()

	item, _ := se.Info("se_id")

	fmt.Printf(item.ID)
}

// MARK: - PersonalData

func ExamplePersonalData_Create() {
	pd := yookassa.NewClient(ycID, ycKey).PersonalData()

	item, _ := pd.Create(models.PersonalDataRequest{
		Type:       "sbp_payout_recipient",
		LastName:   "Иванов",
		FirstName:  "Иван",
		MiddleName: "Иванович",
		Metadata: models.Metadata{
			"recipient_id": "37",
		},
	})

	fmt.Println(item.ID)
}

func ExamplePersonalData_Info() {
	pd := yookassa.NewClient(ycID, ycKey).PersonalData()

	item, _ := pd.Info("some_pd_id")

	fmt.Println(item.ID)
}

// MARK: - Deals

func ExampleDeals_Create() {
	deals := yookassa.NewClient(ycID, ycKey).Deals()

	deal, _ := deals.Create(models.DealRequest{
		Type:      "safe_deal",
		FeeMoment: "payment_succeeded",
		Metadata: models.Metadata{
			"order_id": "37",
		},
		Description: "SAFE_DEAL 123554642-2432FF344R",
	})

	fmt.Println(deal.ID)
}

func ExampleDeals_List() {
	deals := yookassa.NewClient(ycID, ycKey).Deals()

	resp, _ := deals.List(models.ListRequest{
		Limit: 10,
	})

	for _, deal := range resp.Items {
		fmt.Println(deal.ID)
	}
}

func ExampleDeals_Info() {
	deals := yookassa.NewClient(ycID, ycKey).Deals()

	deal, _ := deals.Info("some_deal_id")

	fmt.Println(deal.ID)
}

// MARK: - Webhooks

func ExampleWebhooks_Create() {
	webhooks := yookassa.NewClient(ycID, ycKey).Webhooks(ycToken)

	wh, _ := webhooks.Create(models.WebhookRequest{
		Event: "payment.succeeded",
		URL:   "https://www.merchant-website.com/notification_url",
	})

	fmt.Println(wh.ID)
}

func ExampleWebhooks_List() {
	webhooks := yookassa.NewClient(ycID, ycKey).Webhooks(ycToken)

	list, _ := webhooks.List()

	for _, wh := range list.Items {
		fmt.Println(wh.ID)
	}
}

func ExampleWebhooks_Delete() {
	webhooks := yookassa.NewClient(ycID, ycKey).Webhooks(ycToken)
	_ = webhooks.Delete("some_webhook_id")
}

// MARK: - ClientMethods

func ExampleClient_SBPList() {
	c := yookassa.NewClient(ycID, ycKey)

	list, _ := c.SBPList()

	for _, info := range list.Items {
		fmt.Println(info.Name)
	}
}

func ExampleClient_Me() {
	c := yookassa.NewClient(ycID, ycKey)

	info, _ := c.Me()

	fmt.Println(info.AccountID)
}
