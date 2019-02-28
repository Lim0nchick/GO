package main

import (
    "fmt" // пакет для форматированного ввода вывода
    "net/http" // пакет для поддержки HTTP протокола
    "strings" // пакет для работы с UTF-8 строками
    "log" // пакет для логирования
    "io/ioutil"
)

func WriteFromFile(w http.ResponseWriter, file string) {
    bs, err := ioutil.ReadFile(file)
    if err != nil {
        return
    }
    fmt.Fprintf(w, string(bs))
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //анализ аргументов,
    fmt.Println(r.Form) // ввод информации о форме на стороне сервера
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    var isParamA bool = false

    for k, v := range r.Form {
	if (k == "a") {
            fmt.Fprintf(w, strings.Join(v, ""))
            isParamA = true
        }

        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    if isParamA == false {
        WriteFromFile(w, "view.html") // отправляем данные на клиентскую сторону
    }
}

func main() {
    http.HandleFunc("/", HomeRouterHandler) // установим роутер
    err := http.ListenAndServe(":9316", nil) // задаем слушать порт
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
