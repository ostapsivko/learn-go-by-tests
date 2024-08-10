package dictionary

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("this word already has a definition")
	ErrWordDoesNotExist = DictionaryErr("this word is not in the dictionary")
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = meaning
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return err
}

func (d Dictionary) Update(word, newMeaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newMeaning
	default:
		return err
	}

	return err
}
