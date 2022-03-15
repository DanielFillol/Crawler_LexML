package main

import (
	"CrawlerLexMl/CSV"
	"CrawlerLexMl/Crawler"
	"CrawlerLexMl/WebDriver"
	"time"
)

func main() {

	startTime := time.Now()

	driver := WebDriver.SeleniumWebDriver()

	data := Crawler.DayCrawler(driver, startTime)

	CSV.ExportCSV("CrawlerLexML","",  data)

}
