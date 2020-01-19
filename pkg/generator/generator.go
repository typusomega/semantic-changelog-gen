package generator

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Generator interface {
	Generate(dir string) error
}

func New() Generator {
	return &generator{}
}

type generator struct {
}

func (it *generator) Generate(dir string) error {
	dirPath, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	logrus.Infof("Generating changelog for git repository %v", dirPath)

	return nil
}
