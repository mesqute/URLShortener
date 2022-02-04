package utilites

import (
	"net/http"
)

// RespondMethodNotAllowed отправляет код 405 с описанием доступных методов
func RespondMethodNotAllowed(w http.ResponseWriter, allowString string) {
	w.Header().Set("Allow", allowString)
	http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)

}

// Respond формирует и отправляет ответ
func Respond(w http.ResponseWriter, status int, data string) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	_, err := w.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}
