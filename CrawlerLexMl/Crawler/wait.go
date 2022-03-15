package Crawler

import (
	"github.com/tebeka/selenium"
)


func WaitXpath(driver selenium.WebDriver, XPATH string) {
	for {
		elem, _ := driver.FindElements(selenium.ByXPATH, XPATH)
		if len(elem) > 0 {
			break
		}
	}
}

