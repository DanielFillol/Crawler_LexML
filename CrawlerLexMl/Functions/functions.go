package Functions

import (
	"CrawlerLexMl/Error"
	"strings"
	"time"
)

const magico = 15

func DayAmount(initDate string, endDate string) int {
	var days int
	date1 := strings.Replace(initDate, "/", "-", -1)
	date2 := strings.Replace(endDate, "/", "-", -1)

	date3, err0 := time.Parse("02-01-2006", date1)
	Error.CheckError(err0)
	date4, err1 := time.Parse("02-01-2006", date2)
	Error.CheckError(err1)

	date5 := date4.Sub(date3).Hours()

	if date5 == 0 {
		days = 1
	} else {
		final := date5 / 24
		days = int(final)
	}

	return days
}

func MagicNumber(days int) int {
	var magic int
	if days < magico {
		magic = days
	} else {
		magic = magico
	}
	return magic
}

func LoopAmount(days int, magicNumber int) int {
	loop := days / magicNumber
	remainder := days % magicNumber

	if days > 1 {
		if remainder != 0 {
			loop = loop + 1
		}
	}

	return loop
}

func AddData(searchDate string, i int) string {
	date1 := strings.Replace(searchDate, "/", "-", -1)
	date2, err0 := time.Parse("02-01-2006", date1)
	Error.CheckError(err0)
	date3 := date2.AddDate(0, 0, i)
	date4 := date3.Format("02-01-2006")
	return date4
}
