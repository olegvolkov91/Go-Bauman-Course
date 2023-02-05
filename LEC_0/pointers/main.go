package main

import "fmt"

type User struct {
	Name string
}

// для того, чтобы метод мог изменять внутренние значения структуры, надо ссылаться на эту структуру с помощью *
func (u *User) changeName(name string) {
	u.Name = name
}

func main() {
	var senseOfLife int = 42
	var p *int = &senseOfLife // &... - операция взятия адреса в памяти

	fmt.Printf("Type of pointer %T; Value of pointer %v\n", p, p)

	var zeroPointer *int // создание пустого указателя
	fmt.Printf("Type of pointer %T; Value of pointer %v\n", zeroPointer, zeroPointer)

	u := User{"Nikita"}
	fmt.Printf("User before changing name -> %v\n", u)

	u.changeName("Viktor")
	fmt.Printf("User after changing name -> %v\n", u)

	var age uint8 = 30
	changeAge(&age) // <- можно не создавать отдельную переменную под поинтер, а сразу передавать указатель

	// Указатели на массивы и почему так не стоит делать
	fmt.Printf("\n")
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation", arr)

	fmt.Printf("\n")

	nums := [3]int{1, 2, 3}
	fmt.Println("Slice before mutation", nums)
	mutationSlc(nums[:])
	fmt.Println("Slice after mutation", nums)

}

func mutationSlc(slc []int) {
	slc[0] = 909
}

func changeAge(age *uint8) {
	*age = 31
}

func mutation(arr *[3]int) {
	//(*arr)[0] = 909
	//(*arr)[2] = 1000
	arr[0] = 909
	arr[2] = 1000
}
