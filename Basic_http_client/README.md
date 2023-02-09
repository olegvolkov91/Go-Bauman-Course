## Простейший API и термины
Создадим простой API, который будет получать информацию про pizza

### Опишем простейшие GET маршруты, которые будут у данного API
* /pizza - возвращает .json со всеми пиццами
* /pizza/:id - возвращает .json по конкретно выбранной ***pizza***

### Router и Handler
* Роутер - это функционал, который принимает на вход адрес запроса и вызывает исполнителя (исполнитель ассоциирован с этим запросом).
* Исполнитель - он же handler, это функция обработчик, которая вызывается при обращении по определённому маршруту

### Шаг 1. Установка маршрутизатора
Чтобы удобно работать с маршрутизатором и не писать его с нуля, установим уже готовую библиотеку: 

```github.com/gorilla/mux```
