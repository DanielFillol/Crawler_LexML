# Crawler_LexML

Projeto que visa capturar informações do site [LexML](https://www.lexml.gov.br). O objetivo é simplesmente obter as informações de forma a possibilitar o consumo das informações de maneira ágil.

Para iniciar a captura de qualquer informação do site basta alterar a URL no parâmetro **driver.Get("")** na função **DayCrawler** no arquivo **crawler.go**

Ao final são gerados um arquivo CSV com todos os tados pesquisados.
 
## Dependências

Para o esse crawler somos dependentes do projeto [selenium](https://github.com/tebeka/selenium#readme), sendo necessário a pré-instalação do [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/) e do [selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/)
Também somos dependentes do [html package](https://pkg.go.dev/golang.org/x/net/html) e do [htmlquery] (https://github.com/antchfx/htmlquery).


## Run
Basta dar o start no selenium server e iniciar o app em GO: 

```java -jar selenium-server-standalone.jar```

```go run main.go```

## Aviso legal (Disclaimer)

O projeto não mantém nem atualiza os dados do LexML, não altera nem transforma os dados durante a coleta. O autor do projeto não assume qualquer responsabilidade sobre o seu uso e resultados.

## Considerações éticas

A) Idealmente o sistema LexML deveria disponibilizar uma API ou no mínimo um web service para facilitar o acesso, como isso não existe projetos como esse são necessários. O LexML não proíbe expressamente o uso de crawlers/scraapers, você pode conferir acessando o robots.txt do sistema. Isso porém não quer dizer que seu IP não será bloqueado caso você execute muitas requisições.

B) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.


C) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
