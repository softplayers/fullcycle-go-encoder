package utils_test

import (
	"encoder/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
				"id": "123",
				"file_path": "c.mp4",
				"status": "hugo_brincalhao"
			}`

	err := utils.IsJson(json)

	require.Nil(t, err)

	json = "fake"

	err = utils.IsJson(json)

	require.Error(t, err)
}
