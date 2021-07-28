package domain_test

import (
  "testing"
  "encoder/domain"
  "github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
  video := domain.NewVideo()
  err := video.Validate()

  require.Error(t, err)
}
