package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDisplayCartServiceResponse_ToString(t *testing.T) {
	dcsr := &DisplayCartServiceResponse{
		Items:              nil,
		TotalPrice:         0,
		AppliedPromotionID: 0,
		TotalDiscount:      0,
	}

	res, err := json.Marshal(dcsr)
	require.Nil(t, err)

	expected := string(res)
	actual := dcsr.ToString()

	assert.Equal(t, expected, actual)
}
