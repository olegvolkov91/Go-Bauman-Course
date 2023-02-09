## BASIC PIZZA REST API

### Шаг 1. Инициализация
Инициализируем модуль с ранее изученной командой 
```
  go mod init [название_модуля]
```
### Шаг 2. Базовый скелет приложения
```
package main

import (
  "log"
  "net/http"
)

var (
  port string = "8080"
)

func main() {
  log.Println("Starting pizza rest api ...")
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

