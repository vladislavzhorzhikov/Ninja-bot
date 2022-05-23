package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	"github.com/Syfaro/telegram-bot-api"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"Ninja-bot/boobs"
	"Ninja-bot/config"
	"Ninja-bot/database"
	"Ninja-bot/randoms"
)

const (
	dbDriverName = "postgres"
)

var chatId = map[int64]struct{}{
	-1001269173682: {}, //JS NINJA sente
}

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
				if checkEmployee(randomEmployee, update) {
					randomEmployee = append(randomEmployee, update.Message.From.UserName)
				}
			case "stop":
				if randomEmployee != nil {
					reply = randoms.Random(randomEmployee, rnd)
				}
				isRandomEmployees = false
				randomEmployee = nil
			}
		}

		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот что б вы не ебались с рандом оргом"
		case "menu":
			reply = "/random - зарандомить лоха на заказ еды \n/delivery - даже блять с едой определиться не можете...рандом доставки\n/boobs-сиськи!"
		case "random":
			if _, ok := chatId[update.Message.Chat.ID]; ok {
				reply = "Не хочешь ебаться с рандом оргом? Кто хочет кушать тыкает на /eat. Зарандомимся как мужики тут. Для завершения кликай /stop"
				isRandomEmployees = true
			}
		case "delivery":
			reply = randoms.RandomDelivery()
		case "boobs":
			photo := tgbotapi.NewPhotoShare(update.Message.Chat.ID, boobs.RandomBoobs())
			_, err := bot.Send(photo)
			for err != nil {
				photo := tgbotapi.NewPhotoShare(update.Message.Chat.ID, boobs.RandomBoobs())
				_, err = bot.Send(photo)
			}
		}

		if reply != "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}
	}
}

func checkEmployee(randomEmployee []string, update tgbotapi.Update) bool {
	for _, empl := range randomEmployee {
		if empl == update.Message.From.UserName {
			return false
		}
	}
	return true
}
