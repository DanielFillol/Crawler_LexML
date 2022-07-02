# Crawler_LexML
Project for crawlling information on [LexML](https://www.lexml.gov.br)
CSV file is generate with collected data.
 
## Dependencies
- [Selenium](https://github.com/tebeka/selenium#readme)
- [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/)
- [Selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/)
- [html package](https://pkg.go.dev/golang.org/x/net/html)
- [htmlquery](https://github.com/antchfx/htmlquery)

## Run
```brew install chrome driver ``` (not needed if you alredy have chrome driver)

```java -jar selenium-server-standalone.jar```

```go run main.go```

- To config a new search go in **crawler.go** file, function **DayCrawler** and alter the URL parameter on **driver.Get("")**


## Notes
- Sometimes chromedriver need a clearnce in security options on MacOS.
- Don't forget to previus install Java.
