## Модульность

### package main 
Пакет main это главный пакет, у него должна быть функция main, которая является точкой запуска программы

### Шаг 1. Инициализация go.mod
* Для начала инициализируем модуль, с помощью следующей команды
```
  go mod init [название_модуля]
``` 
### Шаг 2. Cоздадим отдельную директорию для пакета
* Создадим директорию calculator, после чего создадим в ней файл calculator.go
* Внутри этого файла укажем что это новый пакет и дадим ему имя  
```
  package calculator
```
* Теперь этот пакет доступен в модуле, однако в нём нет никаких данных 
* Создадим в пакете calculator функцию Add
```
  func Add(a, b int) int {...}
```
***ВАЖНО*** - работая с пакетами нужно помнить, если мы хотим иметь доступ к какой-либо сущности, её полям,
методу или переменным другого пакета, они должны быть указаны с большой буквы

### Шаг 3. Импортируем ранее созданный пакет в main.go
* Теперь в файле main.go можно обратиться к методу Add пакета calculator
```
    package main
    
    import (
      "[module_name]/calculator/calculator.go"
      "fmt"  
    )
    
    func main() {
      res := calculator.Add(5, 10)
      fmt.Println(res) // <- Результат: 15
    }
```
***
## Работа с функцией init()
Данная функция вызывается единожды при импортировании пакета
В момент инициализации пакета происходит следующее:

* Сначала компилятор смотрит на содержимое пакета
* Затем на пути импорта, если что-то импортируется, компилятор уходит туда
* Затем компилятор инициализирует переменные уровня пакета
* Затем запускается функция init для данного пакета
* Повторяет данную процедуру для всех пакетов проекта
* После чего вызывается функция main()

Добавив к предыдущему примеру функцию init, при выполнении кода сначала будет распечатан текст из функции init, а затем выполнится функция main
```
  func init() {
    fmt.Println("Hello from init")
  } 
```
***Заметка*** - все импорты (вне зависимости, стандартные или пользовательские) сортируются по алфавиту