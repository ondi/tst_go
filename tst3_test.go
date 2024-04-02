//
//
//

package tst

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Tst3_01(t *testing.T) {
	temp := NewTree3[string]()

	temp.Add("/", "/")
	temp.Add("/debug", "/debug")
	temp.Add("/debug/size", "/debug/size")
	temp.Add("/debug/often", "/debug/often")
	temp.Add("/debug/metrics", "/debug/metrics")

	var ok bool
	var value string

	value, ok = temp.Search("")
	assert.Assert(t, ok == false)

	value, ok = temp.Search("v1/test")
	assert.Assert(t, ok == false)

	value, ok = temp.Search("/debu")
	assert.Assert(t, ok == true)
	assert.Assert(t, value == "/", value)

	value, ok = temp.Search("/v1/test")
	assert.Assert(t, value == "/")

	value, ok = temp.Search("/debug/test")
	assert.Assert(t, value == "/debug")

	value, ok = temp.Search("/debug/size")
	assert.Assert(t, value == "/debug/size")

	value, ok = temp.Search("/debug/often")
	assert.Assert(t, value == "/debug/often")

	value, ok = temp.Search("/debug/often/very")
	assert.Assert(t, value == "/debug/often")

	value, ok = temp.Search("/debug/metrics")
	assert.Assert(t, value == "/debug/metrics")

	value, ok = temp.Search("/debug/metrics2")
	assert.Assert(t, value == "/debug/metrics")
}
