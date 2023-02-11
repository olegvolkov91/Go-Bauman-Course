## Select в Go
***
### Использование Select
Select - это инструмент, позволяющий выбирать из множества канальных операций (чтение/запись) для множества каналов.

* Если из 10 каналов что-то пришло в один - select выбирает его
* Если из 10 каналов что-то пришло сразу в два и более - select выбирает случайный
* На практике select чаще всего используется для того, чтобы предпринимать какие-то действия, пока в каналы еще не пришли данные
```
func process(ch chan string) {
  time.Sleep(10500 * time.Millisecond)
  ch <- "process successful"
}

func main() {
  ch := make(chan string)
  go process(ch)
  for {
    time.Sleep(1000 * time.Millisecond)
    select {
    case v := <-ch:
      fmt.Println("received value: ", v)
      return
    default:
      fmt.Println("no value received")
    }
  }
}
```
***
### Select как инструмент защиты от deadlock
Select - это инструмент, позволяющий выбирать из множества канальных операций (чтение/запись) для множества каналов.
```
func main() {
  var ch chan string
  select {
  case v := <-ch:
    fmt.Println("received value", v)
  default:
    fmt.Println("default case executed")
  }
}
```