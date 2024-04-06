//
// ternary search tree
//

package tst

const (
	FnvOffset64 = 14695981039346656037
	FnvPrime64  = 1099511628211
)

type key3_t struct {
	hash uint64
	pos  int32
	code int32
}

type mapped3_t[Value_t any] struct {
	value Value_t
}

type Tree3_t[Value_t any] struct {
	root map[key3_t]*mapped3_t[Value_t]
}

func NewTree3[Value_t any]() *Tree3_t[Value_t] {
	return &Tree3_t[Value_t]{
		root: map[key3_t]*mapped3_t[Value_t]{},
	}
}

func (self *Tree3_t[Value_t]) Add(prefix string, value Value_t) {
	var i int
	var ok bool
	key := key3_t{hash: FnvOffset64}
	for i, key.code = range prefix {
		key.pos = int32(i)
		key.hash ^= uint64(key.code)
		key.hash *= FnvPrime64
		if _, ok = self.root[key]; !ok {
			self.root[key] = nil
		}
	}
	self.root[key] = &mapped3_t[Value_t]{value: value}
}

func (self *Tree3_t[Value_t]) Search(in string) (value Value_t, ok bool) {
	var i, count int
	var temp *mapped3_t[Value_t]
	key := key3_t{hash: FnvOffset64}
	for i, key.code = range in {
		key.pos = int32(i)
		key.hash ^= uint64(key.code)
		key.hash *= FnvPrime64
		if temp, ok = self.root[key]; !ok {
			return value, count > 0
		}
		if temp != nil {
			value = temp.value
			count++
		}
	}
	return value, count > 0
}
