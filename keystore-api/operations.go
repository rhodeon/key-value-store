package main

func putValue(key string, value string) error {
	store[key] = value
	return nil
}

func getValue(key string) (string, error) {
	value, exists := store[key]
	if !exists {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func deleteValue(key string) error {
	delete(store, key)
	return nil
}
