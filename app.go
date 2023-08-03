package main

import (
	"database/sql"
	"warungdana/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB()

	// list api
	//api for test number 2
	e.GET("/soal_nomor_2", handlers.Number2(db))
	//api for test number 3
	e.GET("/soal_nomor_3", handlers.Number3(db))
	//api for test number 4
	e.GET("/soal_nomor_4", handlers.Number4(db))
	//api for test number 5
	e.GET("/soal_nomor_5", handlers.Number5(db))
	//api for test number 6
	e.POST("/soal_nomor_6", handlers.Number6(db))
	//api for test number 7
	e.GET("/soal_nomor_7", handlers.Number7(db))
	//api for test number 8
	e.GET("/soal_nomor_8", handlers.Number8(db))
	//api for test number 9a and 9b
	e.GET("/soal_nomor_9a_9b", handlers.Number9(db))
	//api for test number 9c
	e.GET("/soal_nomor_9c", handlers.Number9c(db))
	//api for test number 9d
	e.GET("/soal_nomor_9d", handlers.Number9d(db))
	//api for test number 10
	e.GET("/soal_nomor_10", handlers.Number10a(db))

	e.Logger.Fatal(e.Start(":8080"))
}

func initDB() *sql.DB {
	//db, err := sql.Open("sqlite3", filepath)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/warung_dana")
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}
