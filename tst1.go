//
// ternary search tree
//

package tst

type node1_t struct {
	eq_kid *node1_t
	hi_kid *node1_t
	lo_kid *node1_t
	key    rune
	value  interface{} // prefix terminator
}

type Tree1_t struct {
	root *node1_t
}

func (self *Tree1_t) Root() *node1_t {
	return self.root
}

func (self *Tree1_t) Add(in string, value interface{}) {
	cur := &self.root
	var last **node1_t
	for _, key := range in {
		for *cur != nil && key != (*cur).key {
			if key < (*cur).key {
				cur = &(*cur).lo_kid
			} else {
				cur = &(*cur).hi_kid
			}
		}
		if *cur == nil {
			*cur = &node1_t{key: key}
		}
		last = cur
		cur = &(*cur).eq_kid
	}
	if last != nil {
		(*last).value = value
	}
}

func Fetch(root *node1_t, key rune) (next *node1_t, value interface{}) {
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
	next = root.eq_kid
	return
}

func (self *Tree1_t) Search(in string) (value interface{}) {
	next := self.Root()
	var temp interface{}
	for _, symbol := range in {
		next, temp = Fetch(next, symbol)
		if temp != nil {
			value = temp
		}
		if next == nil {
			return
		}
	}
	return
}
