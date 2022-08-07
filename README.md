# go-yookasssa

Пакет для работы с платежным сервисом Yookassa.

## Быстрый старт

```go
yc := NewClient("shop_id", "secret_key")

yc.Payments().Create(models.PaymentRequest{
	Amount: models.Amount{
		Value: "100.00",
		Currency: "RUB",
	},
	Capture: true,
	Confirmation: &models.Confirmation{
		Type: "redirect",
		ReturnURl: "https://www.merchant-website.com/return_url",
	},
	Description: "Заказ №1",
})
```

- [Документация](https://pkg.go.dev/github.com/ReanSn0w/go-yookassa@v0.1.3)
- MIT Лицензия
