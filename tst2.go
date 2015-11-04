//
// ternary search tree
//

package tst

import "unicode/utf8"

type TernaryNode2_t struct {
	hi_kid int
	eq_kid int
	lo_kid int
	key rune
	value string	// prefix terminator
}

type TernaryTree2_t struct {
	root []TernaryNode2_t
}

type Cursor2_t struct {
	root []TernaryNode2_t
	cur int
}

const INTMAX = 1 << 32 - 1

func (self * TernaryTree2_t) Add(str string, value string) {
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
			if last < len(self.root) && self.root[last].eq_kid == INTMAX {
				self.root[last].eq_kid = cur
			}
			self.root = append(self.root, TernaryNode2_t{key: key, eq_kid: INTMAX, hi_kid: INTMAX, lo_kid: INTMAX})
		}
		last = cur
		cur = self.root[cur].eq_kid
	}
	if last != INTMAX {
		self.root[last].value = value
	}
}

func (self * TernaryTree2_t) Cursor() (c * Cursor2_t) {
	c = &Cursor2_t{}
	c.root = self.root
	return
}

func (self * Cursor2_t) Fetch(key rune) (value string, next bool) {
	for self.cur < len(self.root) && key != self.root[self.cur].key {
		if key < self.root[self.cur].key {
			self.cur = self.root[self.cur].lo_kid
		} else {
			self.cur = self.root[self.cur].hi_kid
		}
	}
	if self.cur == INTMAX {
		return value, false
	}
	if len(self.root[self.cur].value) > 0 {
		value = self.root[self.cur].value
	}
	self.cur = self.root[self.cur].eq_kid
	return value, self.cur != INTMAX
}

func (self * TernaryTree2_t) Search(str string) (int, int, string, bool) {
	var found string
	var value string
	var next bool
	last := 0
	c := self.Cursor()
	for n, key := range str {
		value, next = c.Fetch(key)
		if len(value) > 0 {
			found = value
			last = n + utf8.RuneLen(key)
		}
		if next == false {
			return last, n, found, false
		}
	}
	return last, len(str), found, next
}
