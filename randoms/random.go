package randoms

import (
	"math/rand"

	"github.com/pkg/errors"

	"Ninja-bot/database"
)

var randomPhrase = []string{
	"Бля, как можно быть таким лохом... ",
	"ХАХАХАХАХ да тебя наебали! ",
	"Ну хотя бы кэшбэк побольше получишь. ",
	"Вытри, сука, слезы! Тебе еще в доставку звонить! ",
	"Хотел подкрутить что б заказывал Роман, но ",
	"Дибил нахуй! ",
	"Ну носи тогда блять с собой, а не ной тут! ",
}

func Random(randomEmployee []string, rnd *database.Randoms) string {
	indPhr := rand.Intn(len(randomPhrase))
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
	return randomPhrase[indPhr] + s + randomEmployee[indEmpl]
}
