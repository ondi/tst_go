//
// ternary search tree
//

package tst

const (
	FnvOffset64 = 14695981039346656037
	FnvPrime64  = 1099511628211
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
	var ok bool
	key := key3_t{hash: FnvOffset64}
	for i, v := range in {
		key.i = i
		key.hash ^= uint64(v)
		key.hash *= FnvPrime64
		if _, ok = self.root[key]; !ok {
			self.root[key] = nil
		}
	}
	temp := new(Value_t)
	*temp = value
	self.root[key] = temp
}

func (self *Tree3_t[Value_t]) Search(in string) (value Value_t, ok bool) {
	var count int
	var temp *Value_t
	key := key3_t{hash: FnvOffset64}
	for i, v := range in {
		key.i = i
		key.hash ^= uint64(v)
		key.hash *= FnvPrime64
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
