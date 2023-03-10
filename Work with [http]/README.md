## Стандартный WebServer

### Шаг 1. Инициализация go mod
```
go mod init [название_модуля]
```

### Шаг 2. Где найти стандартные шаблоны?
**Полезная ссылка**: [структурирование проекта](https://github.com/golang-standards/project-layout)

### Шаг 3. Создадим входную точку в приложение
Стандартный шаблон входной точки
```
cmd/<app_name>/main.go
```

### Шаг 3. Инициализация ядра сервера
Стандартным шаблоном диктуется следующая схема
```
internal/app/<app_name>/<app_name>.go
```

### Шаг 4. Важный пункт про конфигурацию
**Правило**: в Go принято, что:
* конфигурация всегда хранится в сторонних файлах (.toml, .env)
* в Go проектах ВСЕГДА присутствуют дефолтные настройки (исключение - бд стараются не дефолтить)
