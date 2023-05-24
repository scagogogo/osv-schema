package model

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUnmarshalFromJson(t *testing.T) {
	bytes, err := os.ReadFile("test_data/GHSA-vxv8-r8q2-63xw.json")
	assert.Nil(t, err)

	json, err := UnmarshalFromJson[any, any](bytes)
	assert.Nil(t, err)
	if err != nil {
		t.Log(err.Error())
	}
	assert.NotNil(t, json)

}
