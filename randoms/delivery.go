package randoms

import "math/rand"

var randomDeliverySlice = []string{
	"33 - https://yesdostavka.ru/",
	"Еда домой - https://xn--58-6kcpbd5axrb.xn--p1ai/",
	"Блинбери - https://xn--90aamkcop0a.xn--p1ai/",
	"Пхали-хинкали - https://phali-hinkali.ru/",
	"Xan Doner - https://eda.yandex.ru/penza/r/xan_doner",
}

var randomDeliveryPhrase = []string{
	"Мужчины, предлагаю ахуительно отобедать. ",
	"Я бы там повариху отжарил. ",
	"Нет что б домашней еды с собой взять, он блять деньги тратит. ",
	"Чисто для пидоров еда. Заказывай! ",
	"Вкуснее материнского молока! ",
	"За такую хавку надо сосать...Очень вкусно! ",
}

func RandomDelivery() string {
	indPhr := rand.Intn(len(randomDeliveryPhrase))
	indDlv := rand.Intn(len(randomDeliverySlice))
	return randomDeliveryPhrase[indPhr] + randomDeliverySlice[indDlv]
}
