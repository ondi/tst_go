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
		for {
			if *cur == nil {
				*cur = &node1_t{key: key}
				last = cur
				cur = &(*cur).eq_kid
				break
			} else if key < (*cur).key {
				cur = &(*cur).lo_kid
			} else if key > (*cur).key {
				cur = &(*cur).hi_kid
			} else {
				cur = &(*cur).eq_kid
				last = cur
				break
			}
		}
	}
	(*last).value = value
}

func Fetch(root *node1_t, key rune) (next *node1_t, value interface{}) {
	next = root
	for next != nil && key != next.key {
		if key < next.key {
			next = next.lo_kid
		} else {
			next = next.hi_kid
		}
	}
	if next == nil {
		return
	}
	value = next.value
	next = next.eq_kid
	return
}

func (self *Tree1_t) Search(str string) (value interface{}) {
	next := self.Root()
	var temp interface{}
	for _, symbol := range str {
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
