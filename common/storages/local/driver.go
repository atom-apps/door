package local

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/atom-apps/door/common/storages"
	"github.com/atom-apps/door/database/models"
)

var _ (storages.Storage) = (*LocalDriver)(nil)

type LocalDriver struct {
	model *models.Driver
}

func New(params *models.Driver) storages.Storage {
	return &LocalDriver{model: params}
}

func (driver *LocalDriver) ID() uint64 {
	return driver.model.ID
}

// Append implements storages.Storage.
func (*LocalDriver) Append(string, io.Reader, int64) (string, error) {
	panic("unimplemented")
}

// Copy implements storages.Storage.
func (*LocalDriver) Copy(string, string) error {
	panic("unimplemented")
}

// Delete implements storages.Storage.
func (*LocalDriver) Delete([]string) error {
	panic("unimplemented")
}

// Exists implements storages.Storage.
func (*LocalDriver) Exists(path string) (bool, error) {
	panic("unimplemented")
}

// Get implements storages.Storage.
func (*LocalDriver) Get(string) (io.ReadCloser, error) {
	panic("unimplemented")
}

// GetURL implements storages.Storage.
func (*LocalDriver) GetURL(string) (string, error) {
	panic("unimplemented")
}

// IsDir implements storages.Storage.
func (*LocalDriver) IsDir(path string) (bool, error) {
	panic("unimplemented")
}

// IsFile implements storages.Storage.
func (*LocalDriver) IsFile(path string) (bool, error) {
	panic("unimplemented")
}

// List implements storages.Storage.
func (*LocalDriver) List(dir, marker string, limit uint) ([]string, string, error) {
	panic("unimplemented")
}

// Move implements storages.Storage.
func (*LocalDriver) Move(string, string) error {
	panic("unimplemented")
}

// Put implements storages.Storage.
func (driver *LocalDriver) Put(dst string, in io.Reader) (string, error) {
	dst = filepath.Join(driver.model.Bucket, dst)
	if !strings.HasPrefix(dst, driver.model.Bucket) {
		return "", errors.New("invalid save path")
	}

	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return "", errors.New("create dir failed, err:" + err.Error())
	}

	fd, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	if _, err := io.Copy(fd, in); err != nil {
		return "", err
	}

	return dst[len(driver.model.Bucket):], nil
}

// PutFile implements storages.Storage.
func (*LocalDriver) PutFile(string, string) (string, error) {
	panic("unimplemented")
}
