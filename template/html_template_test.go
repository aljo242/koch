package template

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO make test

func TestExecuteTemplateHTML(t *testing.T) {
	path := "./test_file/test.html"
	goodPath := "./newpath.html"
	badPath := "./newpath"
	host := "localhost"

	err := ExecuteTemplateHTML(false, host, path, goodPath)
	require.NoError(t, err)

	err = ExecuteTemplateHTML(true, host, path, goodPath)
	require.NoError(t, err)

	path = "./test_file/invalid.html"
	err = ExecuteTemplateHTML(true, host, path, goodPath)
	require.Contains(t, err.Error(), "error executing template")

	path = "./test_file/html.json"
	err = ExecuteTemplateHTML(true, host, path, goodPath)
	require.ErrorIs(t, err, ErrFileNotHTML)
	err = ExecuteTemplateHTML(true, host, path, badPath)
	require.ErrorIs(t, err, ErrFileNotHTML)
}
