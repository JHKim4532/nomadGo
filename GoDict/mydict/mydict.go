package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

// key Not Found error
var (
	errNotFound   = errors.New("not found")
	errCantUpdate = errors.New("Cant update non-existion word")
	errWordExists = errors.New("That word is already exist")
	errCantDelete = errors.New("Can't Delete a word not existed")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add words to dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	// if err == errNotFound {
	// d[word] = def
	// } else if err == nil {
	// return errWordExists
	// }
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
