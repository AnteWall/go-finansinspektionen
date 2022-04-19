package insider

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDatetime_MarshalJSON(t *testing.T) {
	type test struct {
		Date DateTime `json:"date"`
	}
	ti := DateTime(time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC))
	s := test{
		Date: ti,
	}
	marshal, err := json.Marshal(&s)
	assert.Nil(t, err)
	assert.Equal(t, "{\"date\":\"2020-01-01T01:01:01Z\"}", string(marshal))

	s = test{}

	err = json.Unmarshal(marshal, &s)
	assert.Nil(t, err)
	assert.Equal(t, ti, s.Date)
}
