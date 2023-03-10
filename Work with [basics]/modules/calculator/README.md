## Тестирование пакетов

Файлы с модульными тестами принято называть следующим образом:
* <script_name>_test.go
* <package_name>_test.go

### Название тестовых функций
В Go комьюнити, когда тестируем функции, методы, структуры, интерфейсы и т.д,
принято создавать отдельную функцию для каждого юнита и именовать её следующим образом 
Test[имя_функции]

Для запуска тестов используется следующая команда
```
  go test -v
```
У этой команды может быть множество флагов, детальнее смотреть [тут](https://pkg.go.dev/cmd/go/internal/test)

Флаг -cover показывает % покрытия кода тестами
```
  go test -сover -v
```
***Заметка*** - 75-80% покрытия обычно более чем достаточно

Для удобства работы с тестами, проведения различных сравнений и тд
можно использовать модуль [assert](github.com/stretchr/testify/assert
), для этого необходимо выполнить команду 
```
  go get github.com/stretchr/testify/assert
```

***ВАЖНО*** - всё что начинается со слова Test, будет запущено командой 
```
  go test -v
```
В Go принято, что создаётся один модуль с тестами на весь пакет (вне зависимости от количества модулей в нём)