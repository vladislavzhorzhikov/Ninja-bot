package boobs

import (
	"fmt"
	"math/rand"
	"strconv"
)

var boobs = map[int]string{
	64: `https://pelotok.net/wp-content/uploads/2020/05/krasivye-devushki-%s.jpg`,
	71: `https://pelotok.net/wp-content/uploads/2020/04/%s-1.jpg`,
	90: "https://pelotok.net/wp-content/uploads/2020/03/%s-2.jpg",
	30: "https://pelotok.net/wp-content/uploads/2020/02/%s-1.jpg",
	29: "https://pelotok.net/wp-content/uploads/2020/01/golye-na-plyazhe-%s.jpg",
}

var mapIndexBoobs = map[int]int{
	0: 64,
	1: 71,
	2: 90,
	3: 30,
	4: 29,
}

func RandomBoobs() string {
	indLink := rand.Intn(len(mapIndexBoobs))
	maxRand := mapIndexBoobs[indLink]
	link := boobs[maxRand]
	indPhoto := strconv.Itoa(rand.Intn(maxRand))
	indPhoto = addZeroForRepo(indLink, indPhoto)
	return fmt.Sprintf(link, indPhoto)
}

func addZeroForRepo(indLink int, indPhoto string) string {
	if indLink == 2 || indLink == 3 {
		switch len(indPhoto) {
		case 1:
			return "00" + indPhoto
		case 2:
			return "0" + indPhoto
		}
	}
	if indLink == 4 {
		if len(indPhoto) == 1 {
			return "0" + indPhoto
		}
	}
	return indPhoto
}
