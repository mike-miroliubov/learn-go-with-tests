package hello

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		return helloEnglishPrefix + "World"
	}

	switch language {
	case "Spanish":
		return helloSpanishPrefix + name
	default:
		return helloEnglishPrefix + name
	}

}
