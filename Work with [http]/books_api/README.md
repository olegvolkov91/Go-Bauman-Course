## Примитивный WebServer

В данном примере рассмотрим реализацию примитивного веб-сервера, в котором будет реализован базовый CRUD без работы с базой данных

### Шаг 1. Используем библиотеку gorilla/mux

Для более удобной и быстрой реализации сервера и работы с роутером используем библиотеку **gorilla/mux** для этого установим зависимость с помощью команды:

```
  go get github.com/gorilla/mux
```

После чего в файле main.go создадим новый экземпляр роутера

### Шаг 2. Используем переменные окружения

Для более удобной работы с переменными окружения используем библиотеку **joho/godotenv** для этого установим зависимость с помощью команды:

```
  go get github.com/joho/godotenv
```

После чего в файле main.go в функции init загрузим .env файл

### Шаг 3. Опишем API маршруты и добавим обработчики

Для начала создадим папку utils с файлом builder.go, для того, чтобы улчшить читаемость main.go

В файле builder.go реализуем две вспомогательные функции

### Шаг 4. Опишем модель данных

Создадим папку models, а в ней файл, в котором опишем структуры и тд

```
type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}
```

### Шаг 5. Завершим описание обработчиков для маршрутов

Создадим папку handlers, в которой будут реализованы обработчики для каждого из роутов, которые были описаны в шаге 3

Пример:

```
func GetBookById(writer http.ResponseWriter, request *http.Request) {
  initHeaders(writer)
  id, err := strconv.Atoi(mux.Vars(request)["id"])
  if err != nil {
    log.Println("error while parsing happened:", err)
    writer.WriteHeader(400)
    msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
    json.NewEncoder(writer).Encode(msg)
    return
  }
  book, ok := models.FindBookById(id)
  log.Println("Get book with id:", id)
  if !ok {
    writer.WriteHeader(404)
    msg := models.Message{Message: "book with that ID does not exists in database"}
    json.NewEncoder(writer).Encode(msg)
  } else {
    writer.WriteHeader(200)
    json.NewEncoder(writer).Encode(book)
  }
}
```

где функция **_initHeaders_** это часть пакета в которой описаны хедеры

```
func initHeaders(writer http.ResponseWriter) {
  writer.Header().Set("Content-Type", "application/json")
}
```
