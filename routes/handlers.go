package routes

import (
	m "URLShortener/models"
	u "URLShortener/utilites"
	"io"
	"log"
	"net/http"
)

// mainHandler обрабатывает все запросы по адресу "/"
func mainHandler(w http.ResponseWriter, r *http.Request) {
	db := m.GetDB()

	// обработка метода POST.
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Плохой запрос", http.StatusBadRequest)
			return
		}

		url := string(b)
		if url == "" {
			http.Error(w, "URL не может быть пустым", http.StatusBadRequest)
			return
		}

		link, err := db.Insert(url)
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			log.Print(err)
			return
		}

		err = u.Respond(w, http.StatusOK, link)
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			log.Print(err)
			return
		}
		return
	}

	// обработка метода GET.
	if r.Method == http.MethodGet {
		link := r.URL.Query().Get("link")

		url, err := db.Get(link)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		err = u.Respond(w, http.StatusOK, url)
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			log.Print(err)
			return
		}

		return
	}

	//Если другой метод, то формирует ответное сообщение об ошибке
	//с описанием разрешенных методов в заголовке.
	allowString := http.MethodPost + ", " + http.MethodGet
	u.RespondMethodNotAllowed(w, allowString)
}
