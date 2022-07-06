//
// ternary search tree
//

package tst

type node2_t[Value_t any] struct {
	hi_kid    int
	eq_kid    int
	lo_kid    int
	key       rune
	value     Value_t
	has_value bool
}

type Tree2_t[Value_t any] struct {
	root []node2_t[Value_t]
}

type Cursor2_t[Value_t any] struct {
	root []node2_t[Value_t]
	cur  int
}

const INTMAX = 1<<32 - 1

func (self *Tree2_t[Value_t]) Add(in string, value Value_t) {
	cur := 0
	last := INTMAX
	for _, key := range in {
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
			self.root = append(self.root, node2_t[Value_t]{key: key, eq_kid: INTMAX, hi_kid: INTMAX, lo_kid: INTMAX})
		}
		last = cur
		cur = self.root[cur].eq_kid
	}
	if last != INTMAX {
		self.root[last].value = value
		self.root[last].has_value = true
	}
}

func (self *Tree2_t[Value_t]) Cursor() (c *Cursor2_t[Value_t]) {
	return &Cursor2_t[Value_t]{root: self.root}
}

func (self *Cursor2_t[Value_t]) Fetch(key rune) (value Value_t, ok bool, next bool) {
	if len(self.root) == 0 {
		return
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
	ok = self.root[self.cur].has_value
	next = true
	self.cur = self.root[self.cur].eq_kid
	return
}

func (self *Tree2_t[Value_t]) Search(in string) (value Value_t, ok bool) {
	c := self.Cursor()
	for _, symbol := range in {
		temp, okvalue, next := c.Fetch(symbol)
		if okvalue {
			value = temp
			ok = true
		}
		if next == false {
			return
		}
	}
	return
}
