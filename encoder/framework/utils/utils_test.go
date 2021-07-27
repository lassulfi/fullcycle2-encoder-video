package utils_test

import (
	"encoder/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
		"id": "c1c0241f-5cf2-45ae-a4e1-8ec3c9894b75",
		"file_path": "convite.mp4",
		"status": "pending"
	}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = "string"
	err = utils.IsJson(json)
	require.Error(t, err)
}
