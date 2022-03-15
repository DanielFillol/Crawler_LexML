package WebDriver

import (
	"CrawlerLexMl/Error"
	"github.com/tebeka/selenium"
)

func SeleniumWebDriver() selenium.WebDriver {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})

	driver, err1 := selenium.NewRemote(caps, "")
	Error.CheckError(err1)

	return driver
}
