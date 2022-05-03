//Демо MSSQL-модуль для копирования данных из одной таблицы в другую за транзакцию.

package mssqlinsert

import (
	"database/sql"
	"log"
)

func SqlInserTrs(insertIntegrTableSql string, db *sql.DB, cs chan string) {
	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
	// Откат действий будет проигнорирован, если позже выполнение транзакции будет зафиксировано.

	stmt, err := tx.Prepare(insertIntegrTableSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use
	// Операторы занимают ресурсы сервера и должны быть закрыты, после выполнения.

	if _, err := stmt.Exec("open source"); err != nil {
		log.Fatal(err)
	}
	// Завершение транзакции
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	cs <- "Значение записано в БД" // передача результата в main-программу
}
