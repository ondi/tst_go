//
// ternary search tree
//

package tst

import "unicode/utf8"

// import "github.com/ondi/go-log"

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

func (self * TernaryTree2_t) Add(str string, value string) {
	var key rune
	var cur int
	var last int
	if len(self.root) == 0 {
		cur = -1
		last = -1
	}
	for _, key = range str {
		// log.Trace("enter loop key = '%c' cur = %v, last = %v, len = %v", key, cur, last, len(self.root))
		for cur != -1 && key != self.root[cur].key {
			last = cur
			if key < self.root[cur].key {
				// log.Trace("'%v' < '%v': going to [%v].lo_kid = %v", key, self.root[cur].key, cur, self.root[cur].lo_kid)
				if self.root[cur].lo_kid == -1 {
					self.root[cur].lo_kid = len(self.root)
					// log.Trace("self.root[%v].lo_kid = %v", last, self.root[cur].lo_kid)
					cur = -1
				} else {
					cur = self.root[cur].lo_kid
				}
			} else {
				// log.Trace("'%v' > '%v': going to [%v].hi_kid = %v", key, self.root[cur].key, cur, self.root[cur].hi_kid)
				if self.root[cur].hi_kid == -1 {
					self.root[cur].hi_kid = len(self.root)
					// log.Trace("self.root[%v].hi_kid = %v", last, self.root[cur].hi_kid)
					cur = -1
				} else {
					cur = self.root[cur].hi_kid
				}
			}
		}
		if cur == -1 {
			cur = len(self.root)
			// log.Trace("creating new node cur = %v, last = %v", cur, last)
			if last != -1 && self.root[last].eq_kid == -1 {
				self.root[last].eq_kid = cur
				// log.Trace("self.root[%v].eq_kid = %v", last, cur)
			}
			self.root = append(self.root, TernaryNode2_t{key: key, eq_kid: -1, hi_kid: -1, lo_kid: -1})
		}
		last = cur
		cur = self.root[cur].eq_kid
	}
	if last != -1 {
		self.root[last].value = value
	}
	// log.Trace("%v", self.root)
}

func (self * TernaryTree2_t) Search(str string) (int, int, string, bool) {
	var key rune
	var value string
	var n int
	var last int
	var cur int
	if len(self.root) == 0 {
		cur = -1
	}
	for n, key = range str {
		// log.Trace("search '%c' cur = %v", key, cur)
		for cur != -1 && key != self.root[cur].key {
			if key < self.root[cur].key {
				cur = self.root[cur].lo_kid
			} else {
				cur = self.root[cur].hi_kid
			}
		}
		if cur == -1 {
			return last, n, value, false
		}
		if len(self.root[cur].value) > 0 {
			value = self.root[cur].value
			_, size := utf8.DecodeRuneInString(str[n:])
			last = n + size
		}
		cur = self.root[cur].eq_kid
	}
	return last, len(str), value, cur != -1
}
