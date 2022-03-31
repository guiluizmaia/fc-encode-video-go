package utils_test

import (
	"encoder/framework/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsJson(t *testing.T) {
	json := `{
				"id": "ab7b7ec9-c12f-4223-853a-5ee35c8595ad",
				"file_path": "convite.mp4",
 				"status": "pending"
			 }`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `wes`
	err = utils.IsJson(json)
	require.Error(t, err)
}
