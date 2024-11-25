package maps

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const ErrorNotFound = DictionaryError("could not find the word you were looking for")
const ErrorWordExists = DictionaryError("cannot add word because it already exists")

func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]

	if !ok {
		return "", ErrorNotFound
	}

	return result, nil
}

func (d Dictionary) Add(term string, definition string) error {
	_, exists := d[term]
	if exists {
		return ErrorWordExists
	}

	d[term] = definition
	return nil
}

func (d Dictionary) Update(term string, definition string) error {
	_, err := d.Search(term)
	if err != nil {
		return err
	}

	d[term] = definition
	return nil
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
