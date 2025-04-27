package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound      = DictionaryErr("could not find the word in the Dictionary")
	ErrAlreadyExists = DictionaryErr("definition already exists")
	ErrDoesntExist   = DictionaryErr("cannot operate a non existing value")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrDoesntExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrDoesntExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
