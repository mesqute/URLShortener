package main

import (
	m "URLShortener/models"
	r "URLShortener/routes"
	"flag"
	"log"
	"net/http"
)

func main() {
	// считывание параметра запуска
	base := flag.String("db", "in-memory", "Выбор БД: in-memory или postgresql")
	flag.Parse()

	// устанавливает выбранную БД
	m.SetDB(*base)

	// конфигурация сервера
	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Routes(),
	}

	// запуск сервера
	log.Printf("Запуск сервера на %s, активная БД: %s", server.Addr, *base)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
