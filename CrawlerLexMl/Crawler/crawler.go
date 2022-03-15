package Crawler

import (
	"CrawlerLexMl/CSV"
	"CrawlerLexMl/Error"
	"CrawlerLexMl/Structs"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"os"
	"strconv"
	"strings"
	"time"
)

var data []Structs.CrawlerData

func DayCrawler(driver selenium.WebDriver, startTime time.Time) []Structs.CrawlerData {

	fmt.Println("Start", startTime.Format("3:4:5"))

	err := driver.Get("https://www.lexml.gov.br/busca/search?smode=advanced;f1-tipoDocumento=Doutrina::Livro")
	Error.CheckError(err)

	WaitXpath(driver, "/html/body/div[3]/table/tbody/tr[1]/td[2]")

	amount, err0 := driver.FindElement(selenium.ByXPATH, "//*[@id=\"itemCount\"]")
	Error.CheckError(err0)

	textTotal, err1 := amount.Text()
	Error.CheckError(err1)

	total, err2 := strconv.Atoi(textTotal)
	Error.CheckError(err2)

	loop := total/20

	if total % 20 != 0 {
		loop = loop + 1
	}

	fmt.Println("Total pages to crawl:", loop)

	for i := 0; i < loop; i++ {
		pageCrawler(driver)

		page, err3 := driver.FindElements(selenium.ByLinkText, "Próxima")
		Error.CheckError(err3)

		err4 := page[0].Click()
		Error.CheckError(err4)

		fmt.Println("Page", i+1, "0K")

		CSV.ExportCSV("Resultado" + strconv.Itoa(i+1), "pages/" ,data)

		if i > 0 {
			err5 := os.Remove("Files CSV/pages/" + "Resultado" + strconv.Itoa(i-1) +".csv")
			Error.CheckError(err5)
		}

	}

	err5 := driver.Close()
	Error.CheckError(err5)

	fmt.Println("Total Exe. time", time.Since(startTime).Hours())

	return data

}

func pageCrawler(driver selenium.WebDriver) {

	sourceCode, err1 := driver.PageSource()
	Error.CheckError(err1)

	time.Sleep(2 * time.Second)

	doc, err2 := htmlquery.Parse(strings.NewReader(sourceCode))
	Error.CheckError(err2)

	tables, err3 := driver.FindElements(selenium.ByXPATH, "/html/body/div[3]/table/tbody/tr[1]/td[2]/div")
	Error.CheckError(err3)

	if len(tables) != 0 {

		for i := 0; i < len(tables); i++ {

			infoType := informationCapture(doc,"Tipo  ", i)

			author := informationCapture(doc,"Autor  ", i)

			title := informationCapture(doc,"Título  ", i)

			date := informationCapture(doc,"Data  ", i)

			classification := informationCapture(doc,"Classificação  ", i)

			linkMore, err9 := tables[i].FindElements(selenium.ByLinkText,"mais")
			Error.CheckError(err9)

			if len(linkMore) > 0 {
				err10 := linkMore[0].Click()
				Error.CheckError(err10)
			}

			subject := informationCapture(doc,"Assuntos  ", i)


			singleData := Structs.CrawlerData{
				InfoType:       infoType,
				Author:         author,
				Title:          title,
				Date:           date,
				Subject:        subject,
				Classification: classification,
			}

			data = append(data, singleData)
		}
	}

}

func informationCapture(doc *html.Node, searchString string, i int) string{
	text := searchString + " não encontrado"

	totalTR, err := htmlquery.QueryAll(doc,"/html/body/div[3]/table/tbody/tr[1]/td[2]/div["+strconv.Itoa(i+1)+"]/table/tbody/tr")
	Error.CheckError(err)

	for j := 0; j < len(totalTR); j++ {

		title, err0 := htmlquery.QueryAll(totalTR[j],"td[2]")
		Error.CheckError(err0)

		for k := 0; k <len(title) ; k++ {

			if htmlquery.InnerText(title[k]) == searchString {

				textFintal, err1 := htmlquery.QueryAll(totalTR[j], "td[3]")
				Error.CheckError(err1)

				if len(textFintal)>0 {
					text = htmlquery.InnerText(textFintal[0])
				}

			}

		}

	}

	return text

}


