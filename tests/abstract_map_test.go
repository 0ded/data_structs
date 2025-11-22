package tests

import (
	"chain/abstract_map"
	"testing"
)

func TestNewAbstractMap(t *testing.T) {
	am := abstract_map.NewAbstractMap()
	if am.Size() != 0 {
		t.Errorf("expected size 0, got %d", am.Size())
	}
}

func TestSetAndGet(t *testing.T) {
	am := abstract_map.NewAbstractMap()
	am.Set("key1", "value1")
	am.Set("key2", 42)

	if value, exists := am.Get("key1"); !exists || value != "value1" {
		t.Errorf("expected 'value1', got %v", value)
	}

	if value, exists := am.Get("key2"); !exists || value != 42 {
		t.Errorf("expected 42, got %v", value)
	}
}

func TestGetOrDefault(t *testing.T) {
	am := abstract_map.NewAbstractMap()
	am.Set("key1", "value1")

	if value := am.GetOrDefault("key1", "default"); value != "value1" {
		t.Errorf("expected 'value1', got %v", value)
	}

	if value := am.GetOrDefault("non_existent_key", "default"); value != "default" {
		t.Errorf("expected 'default', got %v", value)
	}
}

func TestSize(t *testing.T) {
	am := abstract_map.NewAbstractMap()
	if am.Size() != 0 {
		t.Errorf("expected size 0, got %d", am.Size())
	}

	am.Set("key1", "value1")
	if am.Size() != 1 {
		t.Errorf("expected size 1, got %d", am.Size())
	}

	am.Set("key2", "value2")
	if am.Size() != 2 {
		t.Errorf("expected size 2, got %d", am.Size())
	}

	am.Delete("key1")
	if am.Size() != 1 {
		t.Errorf("expected size 1 after deletion, got %d", am.Size())
	}
}

func TestDelete(t *testing.T) {
	am := abstract_map.NewAbstractMap()
	am.Set("key1", "value1")
	am.Set("key2", "value2")

	am.Delete("key1")
	if _, exists := am.Get("key1"); exists {
		t.Error("key1 should not exist after deletion")
	}

	if value, exists := am.Get("key2"); !exists || value != "value2" {
		t.Errorf("expected 'value2', got %v", value)
	}
}

func TestKeys(t *testing.T) {
	am := abstract_map.NewAbstractMap()
	am.Set("key1", "value1")
	am.Set("key2", "value2")

	keys := am.Keys()
	if len(keys) != 2 {
		t.Errorf("expected 2 keys, got %d", len(keys))
	}

	// Check if specific keys are present
	foundKey := false
	for _, key := range keys {
		if key == "key1" {
			foundKey = true
			break
		}
	}
	if !foundKey {
		t.Error("key1 should be present in the keys")
	}
}
