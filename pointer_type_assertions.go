package main

import "fmt"

type A struct {
	Zero int
}

func main() {
	var a *A
	var b interface{}

	b = a

	//Ошибка, т.к. a == nil
	/*
		if t, ok := b.(*A); ok {
			fmt.Println(t.Zero)
		}
	*/

	//Более длинная конструкция
	if t, ok := b.(*A); ok && t != nil {
		fmt.Println(t.Zero)
	}

	//Т.к. краткая форма type assertion возвращает дефолтный тип, если конвертация
	//не удалась, то дефолтным типом для *A будет нулевой указатель, т.е. nil
	//поэтому применительно к указателям конструкцию выше можно упростить до:
	if t := b.(*A); t != nil {
		fmt.Println(t.Zero)
	} //поскольку и в сучае когда b == nil и в случае когда конвертация не удалась t будет nil
}
