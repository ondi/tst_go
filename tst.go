//
// ternary search tree
//

package tst_go

import "os"
import "fmt"

type TernaryNode_t struct {
	eq_kid * TernaryNode_t
	hi_kid * TernaryNode_t
	lo_kid * TernaryNode_t
	key rune
	value string	// prefix termination
}

type TernaryTree_t struct {
	root * TernaryNode_t
}

func (self * TernaryTree_t) Add(str string, value string) {
	var key rune
	cur := &self.root
	var last ** TernaryNode_t
	for _, key = range str {
		for *cur != nil && key != (*cur).key {
			if key < (*cur).key {
				cur = &(*cur).lo_kid
			} else {
				cur = &(*cur).hi_kid
			}
		}
		if *cur == nil {
			*cur = &TernaryNode_t{}
			(*cur).key = key
		}
		last = cur
		cur = &(*cur).eq_kid
	}
	if last != nil {
		(*last).value = value
	}
}

func (self * TernaryTree_t) Search(str string) (int, int, string, bool) {
	var n int
	var prev int
	var last_eq int
	var key rune
	var value string
	cur := self.root
	for n, key = range str {
		for cur != nil && key != cur.key {
			if key < cur.key {
				cur = cur.lo_kid
			} else {
				cur = cur.hi_kid
			}
		}
		if cur == nil {
			return n, prev, value, false
		}
		last_eq++
		if len(cur.value) > 0 {
			value = cur.value
			last_eq = 0
		} else if last_eq > 1 {
			value = ""
		}
		fmt.Fprintf(os.Stderr, "TST EQ '%v' %v %v %v '%v'\n", str[n:], n, prev, last_eq, value)
		cur = cur.eq_kid
		prev = n
	}
	return len(str), prev, value, cur != nil
}
