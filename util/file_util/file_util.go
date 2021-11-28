package file_util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const dirName = ".koch"

var KochDir string
var ConfigFile string
var BaseDir string
var OutputDir string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	KochDir = filepath.Join(home, dirName)
	err = EnsureDir(KochDir)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	ConfigFile = filepath.Join(KochDir, "config/config.json")
	BaseDir = filepath.Join(KochDir, "res")
	OutputDir = filepath.Join(KochDir, "static")
}

// Exists is a basic file util that says if a dir or file exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// EnsureDir checks if a dir exists and creates it if it does not exsits
func EnsureDir(dir string) error {
	if !Exists(dir) {
		err := os.Mkdir(dir, 0750)
		if err != nil {
			return fmt.Errorf("error creating directory %v : %w", dir, err)
		}
	}
	return nil
}

// CopyFile copies filename src to dst
func CopyFile(src, dst string) error {
	filePath := filepath.Clean(src)
	in, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v : %w", src, err)
	}
	defer func() {
		err := in.Close()
		if err != nil {
			log.Error().Err(err).Str("filename", filePath).Msg("error closing the file")
		}
	}()

	// DST FILE
	filePath = filepath.Clean(dst)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v : %w", src, err)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Error().Err(err).Str("filename", filePath).Msg("error closing the file")
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		return fmt.Errorf("error copying %v to %v : %w", src, dst, err)
	}

	return nil
}
