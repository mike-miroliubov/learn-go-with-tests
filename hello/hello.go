package hello

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func Hello(name string) string {
	return "Hello World, " + name
}

func main() {
	fmt.Println(Hello("User"))
}
