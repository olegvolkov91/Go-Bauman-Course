## Waitgroup в Go
***
Waitgroup - это ещё один инструмент для работы с горутинами.

***БАНАЛЬНЫЙ ПРИМЕР*** - waitgroup может использоваться как счётчик горутин,
вместо того, чтобы создавать потенциально несколько горутин и каналов к ним
и дожидаться пока в канал будет записано true, можно воспользоваться waitgroup

* wg.Wait() - блокирует основной поток
* wg.Add() - добавляет элемент счётчика (WaitGroup++)
* wg.Done() - вычитает элемент счётчика (WaitGroup--)

Таким образом когда WaitGroup == 0 делаем вывод, что все горутины отработали!
Рассмотрим пример:
```
func process(i int, wg *sync.WaitGroup) {
  fmt.Println("started Goroutine ", i)
  time.Sleep(2 * time.Second)
  fmt.Printf("Goroutine %d ended\n", i)
  wg.Done() //WaitGroup--
}

func main() {
  no := 5
  var wg sync.WaitGroup
  for i := 0; i < no; i++ {
    wg.Add(1) // WaitGroup++
    go process(i, &wg)
  }
  wg.Wait() // if WaitGroup == 0 ? До тех пор, пока это условие не выполнено - мы заблокированы в данной строке для main горутины
  fmt.Println("All go routines finished executing")
}
```