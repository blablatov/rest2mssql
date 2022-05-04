// Демо web-server-parser модуль rest-запросов.
// Формирует и передает данные MSSQL-модулю для копирования данных между БД.
package main

import (
	"fmt"
	"log"
	"mssqldsn"
	"mssqlinsert"
	"net/http"
	"strings"
	"time"
)

type SapData struct {
	dBaseFrom string `json:"DirectumRX.dbo.Integration_Data"`
	dBaseTo   string `json:"DirectumRXTest.dbo.Integration_Data"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// this handler is returning component path of URL.
func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// parsing of string URL.
	i := strings.Index(r.URL.Path, ":") // get index of symbol ":"
	sd := &SapData{}
	sd.dBaseFrom = strings.TrimPrefix((r.URL.Path[:i]), "/") // get slice of string before symbol ":"
	sd.dBaseTo = r.URL.Path[i+1:]                            // get slice of string after symbol ":"
	fmt.Println("\nБД донор:", sd.dBaseFrom, "\nБД реципиент:", sd.dBaseTo)

	insertIntegrTableSql := "INSERT " + sd.dBaseTo + " SELECT Id, Discriminator, Name," +
		"BusinessUnit, ItemNumber, ItemName FROM " + sd.dBaseFrom + " WHERE BusinessUnit = 65"

	cs := make(chan string) // channel of function sql-client. Канал функции sql-клиента
	// структура DSN
	dd := mssqldsn.DataDsn{
		Debug:    true,
		User:     "user",
		Password: "password",
		Port:     1433,
		Server:   "db-directum",
		Database: "DirectumRX",
	}
	// Вызов метода интерфейса, для формирования DSN и создания подключения к СУБД
	var d mssqldsn.ConDsner = dd
	db := d.SqlConDsn()
	defer db.Close()
	// Вызов метода копирования данных в горутине
	go mssqlinsert.SqlInserTrs(insertIntegrTableSql, db, cs)
	// Получение данных из канала cs горутины
	log.Println("\nРезультат запроса: ", <-cs)
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs время выполнения запроса\n", secs)
}
