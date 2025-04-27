package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word in the Dictionary")
var ErrAlreadyExists = errors.New("definition already exists")

func (d *Dictionary) Search(word string) (string, error) {

	definition, ok := (*d)[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d *Dictionary) Add(word, definition string) error {

	_, ok := (*d)[word]

	if !ok {
		(*d)[word] = definition
		return nil
	}

	return ErrAlreadyExists
}
