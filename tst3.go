//
// ternary search tree
//

package tst

import (
	"hash/fnv"
)

type key3_t struct {
	i    int
	hash uint64
}

type Tree3_t[Value_t any] struct {
	root map[key3_t]*Value_t
}

func NewTree3[Value_t any]() *Tree3_t[Value_t] {
	return &Tree3_t[Value_t]{
		root: map[key3_t]*Value_t{},
	}
}

func (self *Tree3_t[Value_t]) Add(in string, value Value_t) {
	h := fnv.New64a()
	var ok bool
	var key key3_t
	for i := 0; i < len(in); i++ {
		h.Write([]byte{in[i]})
		key.i, key.hash = i, h.Sum64()
		if _, ok = self.root[key]; !ok {
			self.root[key] = nil
		}
	}
	temp := new(Value_t)
	*temp = value
	self.root[key] = temp
}

func (self *Tree3_t[Value_t]) Search(in string) (value Value_t, ok bool) {
	h := fnv.New64a()
	var count int
	var key key3_t
	var temp *Value_t
	for i := 0; i < len(in); i++ {
		h.Write([]byte{in[i]})
		key.i, key.hash = i, h.Sum64()
		if temp, ok = self.root[key]; !ok {
			return value, count > 0
		}
		if temp != nil {
			value = *temp
			count++
		}
	}
	return value, count > 0
}
