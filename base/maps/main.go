package main

import (
	"errors"
)

func main() {

}

type Dictionary map[string]string

const (
	UnknownWordError      = DictionaryErr("UnknownWord")
	AlreadyHaveKeyValue   = DictionaryErr("AlreadyHaveKeyValue")
	DontHaveKey           = DictionaryErr("DontHaveKey")
	CantDeleteNonExistKey = DictionaryErr("CantDeleteNonExistKey")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value := d[key]

	if value == "" {
		return "", UnknownWordError
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch {

	case errors.Is(err, UnknownWordError):
		d[key] = value
	case err == nil:
		return AlreadyHaveKeyValue
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case UnknownWordError:
		return DontHaveKey
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	if err != nil {
		return CantDeleteNonExistKey
	}

	delete(d, key)

	return nil
}
