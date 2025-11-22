package abstract_map

import (
	"data_structs/result"
	"encoding/json"
	"errors"
	"fmt"
)

// AbstractMap struct represents a map that can hold different value types.
type AbstractMap struct {
	data map[string]interface{}
	size int
}

// NewAbstractMap initializes a new AbstractMap.
func NewAbstractMap() *AbstractMap {
	return &AbstractMap{
		data: make(map[string]interface{}),
		size: 0,
	}
}

// Set adds or updates a value in the map.
func (m *AbstractMap) Set(key string, value interface{}) {
	if m.data[key] == nil {
		m.size++
	}
	m.data[key] = value
}

// Get retrieves a value by key. Returns `nil` and a boolean indicating if the key exists.
func (m *AbstractMap) Get(key string) (interface{}, bool) {
	value, exists := m.data[key]
	return value, exists
}

func (m *AbstractMap) Contains(key string) bool {
	_, exists := m.data[key]
	return exists
}

func (m *AbstractMap) GetOrDefault(key string, defaultValue interface{}) interface{} {
	value, exists := m.data[key]
	if exists {
		return value
	}
	return defaultValue
}

func (m *AbstractMap) Size() int {
	return m.size
}

// Delete removes a key-value pair from the map.
func (m *AbstractMap) Delete(key string) {
	m.size--
	delete(m.data, key)
}

// Keys returns all keys in the map.
func (m *AbstractMap) Keys() []string {
	keys := make([]string, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// String returns a string representation of the map.
func (m *AbstractMap) String() string {
	str := "{ "
	for k, v := range m.data {
		str += fmt.Sprintf("%s: %v, ", k, v)
	}
	str += "}"
	return str
}

func From_json(json_str string) result.Result[AbstractMap] {
	tempMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(json_str), &tempMap)
	dictionary := AbstractMap{data: tempMap, size: 0}

	if err != nil {
		return result.Err[AbstractMap](errors.New("Error unmarshalling json"))
	}

	for key, value := range tempMap {
		dictionary.Set(key, value)
	}

	return result.Ok(AbstractMap{data: tempMap, size: dictionary.size})
}
