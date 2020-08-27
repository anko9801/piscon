package client

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

var browserList []string = []string{
	"Isuugle Nickel",
	"Isuilla Icetanuki",
	"Isuri Web Browser",
	"Isuternet Explorer",
	"Isucrosoft Edge",
}

var suffixList []string = []string{
	"-mobile",
	"-bottle",
	"-alpha",
	"-beta",
	"",
}

func GenerateUserAgent() string {
	browser := browserList[rand.Intn(len(browserList))]
	suffix := suffixList[rand.Intn(len(suffixList))]
	var uuidStr string
	u, err := uuid.NewRandom()
	if err != nil {
		uuidStr = "00000000-0000-0000-0000-000000000000"
	} else {
		uuidStr = u.String()
	}

	return fmt.Sprintf("%v%v-%v", browser, suffix, uuidStr)
}

func GenerateBotUserAgent() string {
	u, _ := uuid.NewRandom()
	var uuidStr string
	u, err := uuid.NewRandom()
	if err != nil {
		uuidStr = "00000000-0000-0000-0000-000000000000"
	} else {
		uuidStr = u.String()
	}

	switch rand.Intn(10) {
	case 0:
		switch rand.Intn(2) {
		case 0:
			return fmt.Sprintf("Isuuglebot-Mobile-%v", uuidStr)
		default:
			return fmt.Sprintf("Isuuglebot-%v", uuidStr)
		}

	case 1:
		return fmt.Sprintf("Isuuglebot-Image/%v", uuidStr)

	case 2:
		return fmt.Sprintf("Mediapartners-Isuugle-%v", uuidStr)

	case 3:
		return fmt.Sprintf("%v-IsuhooSlurp", uuidStr)

	case 4:
		switch rand.Intn(2) {
		case 0:
			return fmt.Sprintf("%v-IsuhooFeedSeekerBetaJp", uuidStr)
		default:
			return fmt.Sprintf("%v-IsuhooFeedSeekerJp", uuidStr)
		}

	case 5:
		switch rand.Intn(2) {
		case 0:
			return fmt.Sprintf("crawler (http://listing.isuhoo.co.jp/support/faq/) %v", uuidStr)
		default:
			return fmt.Sprintf("crawler (help.isuhoo.co.jp/help/jp/) %v", uuidStr)
		}

	case 6:
		return fmt.Sprintf("isuingbot-%v", uuidStr)

	case 7:
		return fmt.Sprintf("Baisuspider-%v", uuidStr)

	case 8:
		switch rand.Intn(2) {
		case 0:
			return fmt.Sprintf("Baisuspider-image+%v", uuidStr)
		default:
			return fmt.Sprintf("Baisuspider+%v", uuidStr)
		}

	default:
		var main, suffix string
		switch rand.Intn(9) {
		case 0:
			main = "bot"
		case 1:
			main = "Bot"
		case 2:
			main = "BOT"
		case 3:
			main = "crawler"
		case 4:
			main = "Crawler"
		case 5:
			main = "CRAWLER"
		case 6:
			main = "spider"
		case 7:
			main = "Spider"
		case 8:
			main = "SPIDER"
		}
		switch rand.Intn(10) {
		case 0:
			suffix = fmt.Sprintf("-%v", uuidStr)
		case 1:
			suffix = fmt.Sprintf("_%v", uuidStr)
		case 2:
			suffix = fmt.Sprintf(" %v", uuidStr)
		case 3:
			suffix = fmt.Sprintf(".%v", uuidStr)
		case 4:
			suffix = fmt.Sprintf("/%v", uuidStr)
		case 5:
			suffix = fmt.Sprintf(";%v", uuidStr)
		case 6:
			suffix = fmt.Sprintf("@%v", uuidStr)
		case 7:
			suffix = fmt.Sprintf("(%v)", uuidStr)
		case 8:
			main = fmt.Sprintf("(%v", main)
			suffix = fmt.Sprintf(") %v", uuidStr)
		case 9:
			main = fmt.Sprintf("%v %v", uuidStr, main)
			suffix = ""
		}
		return fmt.Sprintf("%v%v", main, suffix)
	}
}
