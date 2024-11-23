package hello

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return helloPrefix + "World"
	}
	return helloPrefix + name
}
