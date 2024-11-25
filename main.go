package main

import (
	"fmt"
	"learn-go-with-tests/dependency_injection"
	"learn-go-with-tests/hello"
	"log"
	"net/http"
	"os"
)

type ArrayList[T any] struct {
	elements []T
}

func (l *ArrayList[T]) Add(el T) {
	l.elements = append(l.elements, el)
}

func (list *ArrayList[T]) Get(index int) (T, bool) {
	return list.elements[0], true
}

type ArrayList2 struct {
	elements []interface{}
}

func (l *ArrayList2) Add(el interface{}) {
	l.elements = append(l.elements, el)
}

func (list *ArrayList2) Get(index int) (interface{}, bool) {
	return list.elements[0], true
}

func main() {
	fmt.Println(hello.Hello("User", ""))
	dependency_injection.Greet(os.Stdout, "User")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		dependency_injection.Greet(writer, "world")
	})))

	//list := ArrayList[int]{}
	//list.Add(1)
	//list.Add("asdasda") // wont compile
	//el := list.Get(0)   // type is int
	//
	//list2 := ArrayList2{}
	//list2.Add(1)
	//list2.Add("asdasda") // compiles ok!
	//el2 := list2.Get(0)  // type is interfac{}
}
