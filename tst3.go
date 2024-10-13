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

func (self *Tree3_t[Value_t]) Add(prefix string, value Value_t) (ok bool) {
	var i int
	key := key3_t{hash: FnvOffset64}
	var temp *mapped3_t[Value_t]
	for i, key.code = range prefix {
		key.pos = int32(i)
		key.hash ^= uint64(key.code)
		key.hash *= FnvPrime64
		if temp, ok = self.root[key]; !ok {
			self.root[key] = nil
		}
	}
	if temp == nil {
		self.root[key] = &mapped3_t[Value_t]{value: value}
		return true
	}
	return false
}

func (self *Tree3_t[Value_t]) Search(in string) (value Value_t, length int, found int) {
	var ok bool
	var temp *mapped3_t[Value_t]
	key := key3_t{hash: FnvOffset64}
	for length, key.code = range in {
		key.pos = int32(length)
		key.hash ^= uint64(key.code)
		key.hash *= FnvPrime64
		if temp, ok = self.root[key]; !ok {
			return
		}
		if temp != nil {
			found++
			value = temp.value
		}
	}
	return
}
