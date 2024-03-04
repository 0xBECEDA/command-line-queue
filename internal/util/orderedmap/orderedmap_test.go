package orderedmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSet(t *testing.T) {
	om := NewOrderedMap()

	type entry struct {
		key string
		val any
	}

	entries := make([]entry, 5)
	for i := range entries {
		entries[i] = entry{
			key: fmt.Sprintf("%v", rand.Int63()), // generate random data
			val: rand.Int63(),
		}

		om.Set(entries[i].key, entries[i].val)
	}

	keys, vals := om.GetAll()
	assert.Equal(t, len(entries), len(keys))
	assert.Equal(t, len(entries), len(vals))

	for i := range entries {
		assert.Equal(t, entries[i].key, keys[i])
		assert.Equal(t, entries[i].val, vals[i])
	}

	// test updating existing value
	om.Set(entries[0].key, entries[1].val)
	value, ok := om.Get(entries[0].key)
	assert.True(t, ok)
	assert.Equal(t, entries[1].val, value)
}

func TestDelete(t *testing.T) {
	om := NewOrderedMap()

	type entry struct {
		key string
		val any
	}

	entries := make([]entry, 5)
	for i := range entries {
		entries[i] = entry{
			key: fmt.Sprintf("%v", rand.Int63()), // generate random data
			val: rand.Int63(),
		}

		om.Set(entries[i].key, entries[i].val)
	}

	// test removal non-existing value -> nothing changed
	om.DeleteItem("abrakadabra")
	keys, vals := om.GetAll()
	assert.Equal(t, len(entries), len(keys))
	assert.Equal(t, len(entries), len(vals))

	for i := range entries {
		assert.Equal(t, entries[i].key, keys[i])
		assert.Equal(t, entries[i].val, vals[i])
	}

	// test removal existing value
	itemToDeleteKey := entries[3].key
	om.DeleteItem(itemToDeleteKey)

	_, ok := om.Get(itemToDeleteKey)
	assert.False(t, ok)

	keys, vals = om.GetAll()
	assert.Equal(t, len(entries)-1, len(keys))
	assert.Equal(t, len(entries)-1, len(vals))

	expectedKeysAfterDelete := []string{entries[0].key, entries[1].key, entries[2].key, entries[4].key}
	for i := range expectedKeysAfterDelete {
		assert.Equal(t, expectedKeysAfterDelete[i], keys[i])
	}
}
