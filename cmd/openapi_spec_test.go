package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAPISpec(t *testing.T) {
	response := `{
		"openapi": "3.0.0",
		"info": {
			"title": "Glean API",
			"version": "1.0.0"
		}
	}`

	_, cleanup := testutils.SetupTestWithResponse(t, []byte(response))
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdOpenAPISpec()
	cmd.SetOut(b)

	err := cmd.Execute()
	require.NoError(t, err)

	assert.Contains(t, b.String(), "Glean API")
}
