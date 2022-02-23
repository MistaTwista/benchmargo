package geneslic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	cases := []struct {
		Name   string
		List   []string
		Result []string
	}{
		{
			Name:   "simple",
			List:   []string{"one", "one", "two", "three"},
			Result: []string{"one", "two", "three"},
		},
	}

	for _, c := range cases {
		res := Uniq(c.List)
		t.Log(res)
		assert.Equal(t, c.expect, locations.HubIDAtDate(c.zmpf, c.locationID, c.date))
	}
}
