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

type Cursor1_t struct {
	cur **node1_t
}

func (self *Tree1_t) Add(str string, value string) {
	cur := &self.root
	var last **node1_t
	for _, key := range str {
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

func (self *Tree1_t) Cursor() *Cursor1_t {
	return &Cursor1_t{cur: &self.root}
}

func (self *Cursor1_t) Fetch(key rune) (value interface{}, next bool) {
	for *self.cur != nil && key != (*self.cur).key {
		if key < (*self.cur).key {
			self.cur = &(*self.cur).lo_kid
		} else {
			self.cur = &(*self.cur).hi_kid
		}
	}
	if *self.cur == nil {
		return
	}
	value = (*self.cur).value
	self.cur = &(*self.cur).eq_kid
	return value, true
}

func (self *Tree1_t) Search(str string) (found interface{}) {
	c := self.Cursor()
	for _, symbol := range str {
		value, ok := c.Fetch(symbol)
		if value != nil {
			found = value
		}
		if ok == false {
			return
		}
	}
	return
}
