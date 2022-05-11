package main

import (
	"Ninja-bot/config"
	"Ninja-bot/database"
	"Ninja-bot/randoms"
	"database/sql"
	"github.com/Syfaro/telegram-bot-api"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"time"
)

const dbDriverName = "postgres"

func main() {
	rand.Seed(time.Now().Unix())

	conf, err := config.Get()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	db, err := sql.Open(dbDriverName, conf.GetDatabaseDSN())
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	rnd := database.NewRandoms(db, conf.GetRandomTableName())

	bot, err := tgbotapi.NewBotAPI(conf.GetTGToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	var isRandomEmployees bool
	var randomEmployee []string

	for update := range updates {
		reply := ""
		if update.Message == nil {
			continue
		}

		if isRandomEmployees {
			switch update.Message.Command() {
			case "eat":
				randomEmployee = append(randomEmployee, update.Message.From.UserName)
			case "stop":
				if randomEmployee != nil {
					reply = random(randomEmployee, rnd)
				}
				isRandomEmployees = false
				randomEmployee = nil
			}
		}

		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот что б вы не ебались с рандом оргом"
		case "menu":
			reply = "/random - зарандомить лоха на заказ еды \n/delivery - даже блять с едой определиться не можете...рандом доставки"
		case "random":
			reply = "Не хочешь ебаться с рандом оргом? Кто хочет кушать тыкает на /eat. Зарандомимся как мужики тут. Для завершения кликай /stop"
			isRandomEmployees = true
		case "delivery":
			reply = randomDelivery()
		}

		if reply != "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}
	}
}

func random(randomEmployee []string, rnd *database.Randoms) string {
	indPhr := rand.Intn(len(randoms.RandomPhrase))
	s := "Заказывает @"
	indEmpl := rand.Intn(len(randomEmployee))
	randCount, err := rnd.IsPrevious(randomEmployee[indEmpl])
	if err != nil {
		if err.Error() != database.NotMatching {
			errors.WithStack(err)
		}
		rnd.AddUserName(randomEmployee[indEmpl])
	}

	if randCount >= 3 {
		randomEmployee = append(randomEmployee[:indEmpl], randomEmployee[indEmpl+1:]...)
		if len(randomEmployee) == 0 {
			return ""
		}
		indEmpl = rand.Intn(len(randomEmployee))
	}

	rnd.UpCount(randomEmployee[indEmpl], randCount+1)
	return randoms.RandomPhrase[indPhr] + s + randomEmployee[indEmpl]
}

func randomDelivery() string {
	indPhr := rand.Intn(len(randoms.RandomDeliveryPhrase))
	indDlv := rand.Intn(len(randoms.RandomDelivery))
	return randoms.RandomDeliveryPhrase[indPhr] + randoms.RandomDelivery[indDlv]
}
