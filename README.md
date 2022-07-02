# Crawler_LexML

Projeto que visa capturar informações do site [LexML](https://www.lexml.gov.br). O objetivo é simplesmente obter as informações de forma a possibilitar o consumo das informações de maneira ágil.

Para iniciar a captura de qualquer informação do site basta alterar a URL no parâmetro **driver.Get("")** na função **DayCrawler** no arquivo **crawler.go**

Ao final são gerados um arquivo CSV com todos os tados pesquisados.
 
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

- Sometimes chromedriver need a clearnce in security options on MacOS.
- Don't forget to previus install Java.
