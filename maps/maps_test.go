package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Find a known word", func(t *testing.T) {
		// given
		dict := Dictionary{"test": "this is just a test"}

		// when
		result, _ := dict.Search("test")

		// then
		assert.Equal(t, "this is just a test", result)
	})

	t.Run("Error if an unknown word", func(t *testing.T) {
		// given
		dict := Dictionary{"test": "this is just a test"}

		// when
		_, err := dict.Search("foo")

		// then
		assert.EqualError(t, err, "could not find the word you were looking for")
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add a new word", func(t *testing.T) {
		// given
		dict := Dictionary{}

		// when
		_ = dict.Add("test", "this is just a test")

		// then
		result, err := dict.Search("test")
		assert.NoError(t, err)
		assert.Equal(t, "this is just a test", result)
	})

	t.Run("Fail to replace an existing word", func(t *testing.T) {
		// given
		dict := Dictionary{"test": "this is just a test"}

		// when
		err := dict.Add("test", "replacing original definition")

		// then
		assert.EqualError(t, err, "cannot add word because it already exists")
		// original is not replaced
		result, _ := dict.Search("test")
		assert.Equal(t, "this is just a test", result)
	})
}

func TestUpdate(t *testing.T) {
	// given
	dict := Dictionary{"test": "this is just a test"}
	newDefinition := "replacing original definition"

	// when
	err := dict.Update("test", newDefinition)

	// then
	assert.NoError(t, err)
	result, _ := dict.Search("test")
	assert.Equal(t, newDefinition, result)
}

func TestDelete(t *testing.T) {
	// given
	dict := Dictionary{
		"test":  "this is just a test",
		"test2": "this is another test",
	}

	// when
	dict.Delete("test")

	// then
	_, err := dict.Search("test")
	assert.EqualError(t, err, string(ErrorNotFound))
}
