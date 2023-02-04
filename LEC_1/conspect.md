***Важно*** - .json не гарантирует соблюдения упорядоченности при выдаче ключей!

### Шаг 1. Как читают файл .json
* Для начала надо создать файловый дексриптор
```
    jsonFile, err := os.Open("users.json")
``` 
* Сразу же обрабатываем ошибки!
```
    if err != nil {...}
```
* Не забываем закрывать файл после прочтения
```
    defer jsonFile.Close()
```
* Затем нам нужно из файл-дескриптора забрать данные и куда-то их поместить!
```
    json.Unmarshall(byteArr, &куда_записываем)
```
### Шаг 2. Теперь более конкретно
В Go существует 2 способа работы с .json файлами это:
* структуризованая сериализация/десериализация
* не структуризованная сериализация/десериализация

### Шаг 2.1 Cтруктуризация
***Сериализация*** - процесс конвертации объектов в последовательность байтов
***Десериализация*** - процесс конвертации последовательности байтов в объекты