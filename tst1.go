//
// ternary search tree
//

package tst

type node1_t[Value_t any] struct {
	eq_kid    *node1_t[Value_t]
	hi_kid    *node1_t[Value_t]
	lo_kid    *node1_t[Value_t]
	key       rune
	value     Value_t
	has_value bool
}

type Tree1_t[Value_t any] struct {
	root *node1_t[Value_t]
}

func (self *Tree1_t[Value_t]) Root() *node1_t[Value_t] {
	return self.root
}

func (self *Tree1_t[Value_t]) Add(in string, value Value_t) {
	cur := &self.root
	var last **node1_t[Value_t]
	for _, key := range in {
		for *cur != nil && key != (*cur).key {
			if key < (*cur).key {
				cur = &(*cur).lo_kid
			} else {
				cur = &(*cur).hi_kid
			}
		}
		if *cur == nil {
			*cur = &node1_t[Value_t]{key: key}
		}
		last = cur
		cur = &(*cur).eq_kid
	}
	if last != nil {
		(*last).value = value
		(*last).has_value = true
	}
}

func Fetch[Value_t any](root *node1_t[Value_t], key rune) (next *node1_t[Value_t], value Value_t, ok bool) {
	for root != nil && key != root.key {
		if key < root.key {
			root = root.lo_kid
		} else {
			root = root.hi_kid
		}
	}
	if root == nil {
		return
	}
	value = root.value
	ok = root.has_value
	next = root.eq_kid
	return
}

func (self *Tree1_t[Value_t]) Search(in string) (value Value_t, ok bool) {
	next := self.Root()
	var okfetch bool
	var temp Value_t
	for _, symbol := range in {
		next, temp, okfetch = Fetch(next, symbol)
		if okfetch {
			ok = true
			value = temp
		}
		if next == nil {
			return
		}
	}
	return
}
