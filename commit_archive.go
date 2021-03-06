package git

import (
	"fmt"
	"path/filepath"
	"strings"
)

type ArchiveType int

const (
	ZIP ArchiveType = iota + 1
	TARGZ
)

func (c *Commit) CreateArchive(target string, archiveType ArchiveType) error {
	var format string
	switch archiveType {
	case ZIP:
		format = "zip"
	case TARGZ:
		format = "tar.gz"
	default:
		return fmt.Errorf("unknown format: %v", archiveType)
	}

	_, err := NewCommand("archive", "--prefix="+filepath.Base(strings.TrimSuffix(c.repo.Path, ".git"))+"/", "--format="+format, "-o", target, c.ID.String()).RunInDir(c.repo.Path)
	return err
}
