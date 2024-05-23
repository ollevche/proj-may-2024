package hashtable

import "fmt"

// HashTable implements a map with string keys and string values.
type HashTable struct {
	values []*entry
}

type entry struct {
	k, v string
	next *entry
}

func New() *HashTable {
	return &HashTable{
		values: make([]*entry, 25),
	}
}

func (t *HashTable) Set(k, v string) {
	idx := t.getIndexByKey(k)

	if t.values[idx] == nil {
		t.values[idx] = &entry{
			k: k,
			v: v,
		}
		return
	}

	e := t.values[idx]
	for {
		if e.k == k {
			e.v = v
			return
		}

		if e.next != nil {
			e = e.next
			continue
		}

		e.next = &entry{
			k: k,
			v: v,
		}
		return
	}
}

func (t *HashTable) Get(k string) (string, bool) {
	idx := t.getIndexByKey(k)

	e := t.values[idx]
	for {
		if e == nil {
			return "", false
		}

		if e.k == k {
			return e.v, true
		}

		e = e.next
	}
}

func (t *HashTable) Delete(k string) {
	idx := t.getIndexByKey(k)

	e := t.values[idx]
	if e == nil {
		return
	}

	if e.k == k {
		t.values[idx] = nil
		return
	}

	prev, current := e, e.next

	for {
		if current == nil {
			return
		}

		if current.k == k {
			prev.next = current.next
			return
		}

		prev, current = current, current.next
	}
}

func (t *HashTable) getIndexByKey(k string) int {
	hash := getHashByKey(k)

	idx := int(hash) % len(t.values)

	fmt.Printf("Key = `%s`; hash = `%v`; idx = `%v`\n", k, hash, idx)

	return idx
}

func getHashByKey(k string) uint8 {
	// a jenkins one-at-a-time-hash
	// refer https://en.wikipedia.org/wiki/Jenkins_hash_function

	var hash uint8

	for _, ch := range k {
		hash += uint8(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}
