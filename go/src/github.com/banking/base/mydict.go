package banking

import "errors"

var (
	errNotFound   = errors.New("Not found")
	errCantUpdate = errors.New("Fail to update.")
	errCantDelete = errors.New("Fail to delete.")
	errWordExists = errors.New("The word is already exist.")
)

// Dictionary
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]
	if ok {
		return result, nil
	} else {
		return "", errNotFound
	}
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else {
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantUpdate
	} else {
		d[word] = def
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantDelete
	} else {
		delete(d, word)
	}

	return nil
}
