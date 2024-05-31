package main

import (
	"fmt"
	"math"
)

/*
Hash Table:

there is difference between the hash function used in "hash table" with hash function used in "cryptography"
the hash fn related to "hash table" must be:

-fast
-distribute data uniformly
-consistent for the same input return the same output

Big O:
insertion: O(1)
deletion: O(1)
access: O(1)

*/

type HashTable struct {
	keyMap [][]any
}

// complete dummy hash fn
func (h *HashTable) hash(key string) int {
	var total int
	primeNumber := 31
	min := math.Min(100, float64(len(key)))
	for i := 0; i < int(min); i++ {
		r := key[i] - 96
		total = (total*primeNumber + int(r)) % len(h.keyMap)
	}
	return total
}

func (h *HashTable) Set(key string, val any) {
	idx := h.hash(key)
	bucket := h.keyMap[idx]
	//nothing in there, we create the array
	if len(bucket) == 0 {
		h.keyMap[idx] = []any{[]any{key, val}}
		return
	}
	//there is something inside of bucket
	bucket = append(bucket, []any{key, val})
	h.keyMap[idx] = bucket
}

func (h *HashTable) Get(key string) any {
	idx := h.hash(key)
	bucket := h.keyMap[idx]
	if len(bucket) == 0 {
		return nil
	}
	//we iterate over it
	for _, items := range bucket {
		item := items.([]any)
		keyFromBucket, val := item[0].(string), item[1]
		if keyFromBucket == key {
			return val
		}
	}
	return nil
}

func main() {
	h := HashTable{
		keyMap: make([][]any, 53),
	}
	h.Set("john", 12)
	h.Set("jane", "doe")
	h.Set("wiki", true)
	fmt.Println(h.Get("wiki"))
	fmt.Println(h.Get("alice"))
}
