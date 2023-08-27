package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// вызвается при внутренних ошибках сервера
func (app *application) serverError(w http.ResponseWriter, err error){
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// вызывается при проблемах с пользовательским запросом
func (app *application) clientError(w http.ResponseWriter, status int){
	http.Error(w, http.StatusText(status), status)
}

// отдельная оболочка для 404 ошибки
func (app *application) notFound (w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

