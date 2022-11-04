package main

import "errors"

var store = make(map[string]string)
var ErrNoSuchKey = errors.New("no such key")

func main() {

}

func Put(key string, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, exists := store[key]
	if !exists {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}
