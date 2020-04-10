//
// ternary search tree
//

package tst

type node2_t struct {
	hi_kid int
	eq_kid int
	lo_kid int
	key    rune
	value  interface{} // prefix terminator
}

type Tree2_t struct {
	root []node2_t
}

type Cursor2_t struct {
	root []node2_t
	cur  int
}

const INTMAX = 1<<32 - 1

func (self *Tree2_t) Add(str string, value interface{}) {
	cur := 0
	last := INTMAX
	for _, key := range str {
		for cur < len(self.root) && key != self.root[cur].key {
			if key < self.root[cur].key {
				if self.root[cur].lo_kid == INTMAX {
					self.root[cur].lo_kid = len(self.root)
				}
				cur = self.root[cur].lo_kid
			} else {
				if self.root[cur].hi_kid == INTMAX {
					self.root[cur].hi_kid = len(self.root)
				}
				cur = self.root[cur].hi_kid
			}
		}
		if cur >= len(self.root) {
			cur = len(self.root)
			if last != INTMAX && self.root[last].eq_kid == INTMAX {
				self.root[last].eq_kid = cur
			}
			self.root = append(self.root, node2_t{key: key, eq_kid: INTMAX, hi_kid: INTMAX, lo_kid: INTMAX})
		}
		last = cur
		cur = self.root[cur].eq_kid
	}
	if last != INTMAX {
		self.root[last].value = value
	}
}

func (self *Tree2_t) Cursor() (c *Cursor2_t) {
	return &Cursor2_t{root: self.root}
}

func (self *Cursor2_t) Fetch(key rune) (value interface{}, next bool) {
	if len(self.root) == 0 {
		return nil, false
	}
	for self.cur < len(self.root) && key != self.root[self.cur].key {
		if key < self.root[self.cur].key {
			self.cur = self.root[self.cur].lo_kid
		} else {
			self.cur = self.root[self.cur].hi_kid
		}
	}
	if self.cur == INTMAX {
		return
	}
	value = self.root[self.cur].value
	self.cur = self.root[self.cur].eq_kid
	return value, true
}

func (self *Tree2_t) Search(str string) (value interface{}) {
	c := self.Cursor()
	for _, symbol := range str {
		temp, ok := c.Fetch(symbol)
		if temp != nil {
			value = temp
		}
		if ok == false {
			return
		}
	}
	return
}
