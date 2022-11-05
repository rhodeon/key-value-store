package main

func putValue(key string, value string) error {
	// lock store to prevent concurrent updates
	store.Lock()
	store.data[key] = value
	store.Unlock()

	return nil
}

func getValue(key string) (string, error) {
	store.RLock()
	value, exists := store.data[key]
	store.RUnlock()

	if !exists {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func deleteValue(key string) error {
	store.Lock()
	delete(store.data, key)
	store.Unlock()

	return nil
}
