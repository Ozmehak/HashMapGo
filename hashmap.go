package main

import (
	"crypto/sha256"
	"fmt"
)

type keyValue struct {
	key   string
	value string
}

const arrayLength = 6

func main() {

	var keyValueArray [arrayLength][]keyValue

	for i := 0; i < arrayLength; i++ {
		addKeyValue(&keyValueArray, fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	value := getValueByKey(&keyValueArray, "key3")
	fmt.Println(value)
}

func addKeyValue(keyValueArray *[arrayLength][]keyValue, key string, value string) {
	hash := calculateHash(key)

	index := hash % arrayLength

	(*keyValueArray)[index] = append((*keyValueArray)[index], keyValue{Key: key, Value: value})

}

func getValueByKey(keyValueArray *[arrayLength][]keyValue, key string) string {
	hash := calculateHash(key)

	index := hash % arrayLength

	for _, kv := range (*keyValueArray)[index] {
		if kv.Key == key {
			return kv.Value
		}
	}

	return "not found"
}

func calculateHash(key string) uint32 {
	hash := sha256.New()

	hash.Write([]byte(key))

	hashValue := hash.Sum(nil)
	return uint32(hashValue[0])<<24 | uint32(hashValue[1])<<16 | uint32(hashValue[2])<<8 | uint32(hashValue[3])
}
