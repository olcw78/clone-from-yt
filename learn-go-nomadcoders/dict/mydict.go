package mydict

import "errors"

// alias of type
type Dictionary map[string]string

var errNotFound = errors.New("could not find the word")
var errKeyExists = errors.New("key already exists")
var errCantUpdate = errors.New("Cant update non-existing word")
var errCantDelete = errors.New("Cant delete non-existing word")

// Search for a word
func (d *Dictionary) Search(key string) (string, error) {
	val, exists := (*d)[key]
	if exists {
		return val, nil
	}

	return "", errNotFound
}

func (d *Dictionary) Add(key, val string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		(*d)[key] = val
	case nil:
		return errKeyExists
	}

	return nil
}

func (d *Dictionary) Update(key, val string) error {
	_, err := d.Search(key)
	if err == errNotFound {
		return errCantUpdate
	}

	(*d)[key] = val
	return nil
}

func (d *Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err == errNotFound {
		return errCantDelete
	}

	delete(*d, key)
	return nil
}
