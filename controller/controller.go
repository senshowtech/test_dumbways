package controller

import (
	Model "dumbways/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/streadway/amqp"
)

const jwtSecret = "secret"
const url = "amqp://guest:guest@localhost:5672/"

var transaction []Model.Transaction
var balance []Model.Balance

func ConsumeTransaction(ctx *fiber.Ctx) {

	conn, err := amqp.Dial(url)

	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"Transaction",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			var jsonData = []byte(d.Body)
			var data Model.Transaction

			var err = json.Unmarshal(jsonData, &data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			transaction = append(transaction, Model.Transaction(data))

		}
	}()

	<-forever
}

func PublishTransaction(ctx *fiber.Ctx) {

	var body Model.Transaction
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("There is nothing in Body")
		panic(err)
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)

	x := Model.Transaction{
		Id:     id,
		Price:  body.Price,
		Status: body.Status,
	}

	data, _ := json.Marshal(x)

	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Transaction",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}

	err = ch.Publish(
		"",
		"Transaction",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	ctx.Send("Successfully Published Message to Queue")
}

func ConsumeBalance(ctx *fiber.Ctx) {

	conn, err := amqp.Dial(url)

	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"Balance",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			var jsonData = []byte(d.Body)
			var data Model.Balance

			var err = json.Unmarshal(jsonData, &data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			balance = append(balance, Model.Balance(data))

		}
	}()

	<-forever
}

func PublishBalance(ctx *fiber.Ctx) {

	var body Model.Balance
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("There is nothing in Body")
		panic(err)
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)

	x := Model.Balance{
		Id:     id,
		Wallet: body.Wallet,
		Status: body.Status,
	}

	data, _ := json.Marshal(x)

	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Balance",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}

	err = ch.Publish(
		"",
		"Balance",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	ctx.Send("Successfully Published Message to Queue")
}

func Login(ctx *fiber.Ctx) {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	if body.Email != "febrisena@gmail.com" || body.Password != "password123" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // a week

	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}{
			Id:    1,
			Email: "febrisena@gmail.com",
		},
	})
}

func Transaction(ctx *fiber.Ctx) {

	var data_akhir []Model.Transaction
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)

	if len(transaction) == 0 {
		transaction = append(transaction, Model.Transaction{Price: 0, Status: "Belum ada pembayaran"})
	}

	for i := range transaction {
		if transaction[i].Id == id {
			data_akhir = append(data_akhir, transaction[i])
		}
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": struct {
			Id string `json:"id"`
		}{
			Id: id,
		},
		"transaction": data_akhir,
	})
}

func Balance(ctx *fiber.Ctx) {

	var data_balance []Model.Balance
	var data_transaksi []Model.Transaction

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)

	if len(balance) == 0 {
		data_balance = append(balance, Model.Balance{Wallet: 0, Status: "Belum ada Top Up"})
	}

	for i := range balance {
		if balance[i].Id == id {
			data_balance = append(data_balance, balance[i])
		}
	}

	if len(transaction) != 0 {
		for i := range transaction {
			if transaction[i].Id == id {
				data_transaksi = append(data_transaksi, transaction[i])
			}
		}
		jumlah_transaksi := transaction[len(data_transaksi)-1]
		jumlah_balance := balance[len(data_balance)-1]
		data_balance = append(data_balance, Model.Balance{Wallet: jumlah_balance.Wallet - jumlah_transaksi.Price, Status: "success"})
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": struct {
			Id string `json:"id"`
		}{
			Id: id,
		},
		"balance": data_balance,
	})
}
