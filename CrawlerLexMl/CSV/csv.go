package CSV

import (
	"CrawlerLexMl/Error"
	"CrawlerLexMl/Structs"
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
)

func ExportCSV(nameCSV string, path string, resultado []Structs.CrawlerData) {

	var empData [][]string

	empData = append(empData, []string{"Tipo", "Autor", "Título", "Data","Assuntos", "Classificação"})


	for i := 0; i < len(resultado); i++ {
		final := []string{
			strings.Trim(resultado[i].InfoType, " "),
			strings.Trim(resultado[i].Author, " "),
			strings.Trim(resultado[i].Title, " "),
			strings.Trim(resultado[i].Date, " "),
			strings.Trim(resultado[i].Subject, " "),
			strings.Trim(resultado[i].Classification, " "),
		}
		empData = append(empData, final)
	}

	csvFile, _ := createFile("Files CSV/"+ path +nameCSV + ".csv")
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}

func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

