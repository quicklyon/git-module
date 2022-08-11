package git

import (
	"bytes"
	"io"
)

// GetDiffOrPatch generates either diff or formatted patch data between given revisions
func (repo *Repository) GetDiffOrPatch(base, head string, w io.Writer, formatted bool) error {
	if formatted {
		return repo.GetPatch(base, head, w)
	}
	return repo.GetDiff(base, head, w)
}

// GetDiff generates and returns patch data between given revisions.
func (repo *Repository) GetDiff(base, head string, w io.Writer) error {
	return NewCommand("diff", "-p", "--binary", base, head).
		RunInDirPipeline(w, nil, repo.path)
}

// GetPatch generates and returns format-patch data between given revisions.
func (repo *Repository) GetPatch(base, head string, w io.Writer) error {
	stderr := new(bytes.Buffer)
	err := NewCommand("format-patch", "--binary", "--stdout", base+"..."+head).
		RunInDirPipeline(w, stderr, repo.path)
	if err != nil && bytes.Contains(stderr.Bytes(), []byte("no merge base")) {
		return NewCommand("format-patch", "--binary", "--stdout", base, head).
			RunInDirPipeline(w, nil, repo.path)
	}
	return err
}
