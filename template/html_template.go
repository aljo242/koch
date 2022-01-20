package template

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rs/zerolog/log"
)

type InfoHTML struct {
	Host string
	// TODO add more
}

// ErrFileNotHTML indicates that the file is not html
var ErrFileNotHTML = errors.New("filetype not html")

// ExecuteTemplateHTML is a util func for executing an html template
// at path and saving the new file to newPath
func ExecuteTemplateHTML(secure bool, host, path, newPath string) error {
	filePath := filepath.Clean(newPath)
	// check if file is html
	if filepath.Ext(filePath) != ".html" {
		return fmt.Errorf("%w : %s", ErrFileNotHTML, filePath)
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %v : %w", newPath, err)
	}
	defer func() {
		err := newFile.Close()
		if err != nil {
			log.Error().Err(err).Str("filename", filePath).Msg("error closing the file")
		}
	}()

	pathClean := filepath.Clean(path)
	if filepath.Ext(pathClean) != ".html" {
		return fmt.Errorf("%w : %s", ErrFileNotHTML, pathClean)
	}
	tpl, err := template.ParseFiles(pathClean)
	if err != nil {
		return fmt.Errorf("error creating template : %w", err)
	}

	var httpPrefix string
	if secure {
		httpPrefix = "https://"
	} else {
		httpPrefix = "http://"
	}

	p := InfoHTML{httpPrefix + host}

	err = tpl.Execute(newFile, p)
	if err != nil {
		return fmt.Errorf("error executing template : %w", err)
	}

	return nil
}
