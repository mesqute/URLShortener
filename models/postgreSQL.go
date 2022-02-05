package models

import (
	u "URLShortener/utilites"
	"database/sql"
	_ "github.com/lib/pq"
)

type postgreSQL struct {
	db *sql.DB
}

// Get возвращает полный URL, соответствующий заданной ссылке
func (p *postgreSQL) Get(link string) (URL string, err error) {
	row := p.db.QueryRow("SELECT url FROM urls WHERE link = ?", link)
	err = row.Scan(&URL)
	return
}

// Insert добавляет в БД заданный URL и возвращает сгенерированную ссылку
func (p *postgreSQL) Insert(URL string) (link string, err error) {
	var _link string

	// проверка, есть ли добавляемый ULR в БД
	row := p.db.QueryRow("SELECT * FROM urls WHERE url = $1", URL)
	checkError := row.Scan(&_link)
	if checkError == nil { // если такой URL уже существует в БД
		link = _link // возвращает сгенерированную ранее ссылку
		return
	}

	link = p.findFreeToken()
	_, err = p.db.Exec("INSERT INTO urls (url, link) VALUES ($1, $2)", URL, link)
	return
}

// рекурсивный поиск свободного токена
func (p *postgreSQL) findFreeToken() (token string) {
	for {
		link := u.GenerateToken(10)
		row := p.db.QueryRow("SELECT * FROM urls WHERE link = $1", link)
		err := row.Scan()
		if err != nil {
			token = link
			return
		}
	}
}
