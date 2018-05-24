package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigMustGet(t *testing.T) {
	var assert = assert.New(t)
	var cfg = MustGet()
	assert.Equal("", cfg.GithubAccessToken)

}
