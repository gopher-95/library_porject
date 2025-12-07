package api

import "net/http"

//создали маршрутизатор
var Mux = http.NewServeMux()

// функция для регистрации обработчиков
func Init() {
	Mux.HandleFunc("/", HelloHandler)
}
