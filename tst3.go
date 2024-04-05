//
// ternary search tree
//

package tst

const (
	// key.hash = FnvOffset64
	// ...
	// key.hash ^= uint64(codepoint)
	// key.hash *= FnvPrime64
	FnvOffset64 = 14695981039346656037
	FnvPrime64  = 1099511628211
)

type key3_t struct {
	pos  int
	prev rune
	next rune
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
	var ok bool
	var prev rune
	key := key3_t{}
	for key.pos, key.next = range prefix {
		key.prev = prev
		prev = key.next
		if _, ok = self.root[key]; !ok {
			self.root[key] = nil
		}
	}
	self.root[key] = &mapped3_t[Value_t]{value: value}
}

func (self *Tree3_t[Value_t]) Search(in string) (value Value_t, ok bool) {
	var count int
	var prev rune
	var temp *mapped3_t[Value_t]
	key := key3_t{}
	for key.pos, key.next = range in {
		key.prev = prev
		prev = key.next
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
