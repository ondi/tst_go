//
//
//

package tst

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Tst2_01(t *testing.T) {
	temp := &Tree1_t{}
	temp.Add("/", "/")
	temp.Add("/debug", "/debug")
	temp.Add("/debug/size", "/debug/size")
	temp.Add("/debug/often", "/debug/often")
	temp.Add("/debug/metrics", "/debug/metrics")

	assert.Assert(t, temp.Search("") == nil)
	assert.Assert(t, temp.Search("v1/test") == nil)
	assert.Assert(t, temp.Search("/debu") == "/")
	assert.Assert(t, temp.Search("/v1/test") == "/")
	assert.Assert(t, temp.Search("/debug/test") == "/debug")
	assert.Assert(t, temp.Search("/debug/size") == "/debug/size")
	assert.Assert(t, temp.Search("/debug/often") == "/debug/often")
	assert.Assert(t, temp.Search("/debug/often/very") == "/debug/often")
	assert.Assert(t, temp.Search("/debug/metrics") == "/debug/metrics")
	assert.Assert(t, temp.Search("/debug/metrics2") == "/debug/metrics")
}
