// Copyright 2023 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateServerInfo(t *testing.T) {
	err := os.RemoveAll(filepath.Join(repoPath, "info"))
	require.NoError(t, err)
	err = UpdateServerInfo(repoPath, UpdateServerInfoOptions{Force: true})
	require.NoError(t, err)
	assert.True(t, isFile(filepath.Join(repoPath, "info", "refs")))
}

func TestReceivePack(t *testing.T) {
	got, err := ReceivePack(repoPath, ReceivePackOptions{HTTPBackendInfoRefs: true})
	require.NoError(t, err)
	const contains = "report-status report-status-v2 delete-refs side-band-64k quiet atomic ofs-delta object-format=sha1 agent=git/"
	assert.Contains(t, string(got), contains)
}

func TestUploadPack(t *testing.T) {
	got, err := UploadPack(repoPath,
		UploadPackOptions{
			StatelessRPC:        true,
			Strict:              true,
			HTTPBackendInfoRefs: true,
		},
	)
	require.NoError(t, err)
	const contains = "multi_ack thin-pack side-band side-band-64k ofs-delta shallow deepen-since deepen-not deepen-relative no-progress include-tag multi_ack_detailed no-done symref=HEAD:refs/heads/master object-format=sha1 agent=git/"
	assert.Contains(t, string(got), contains)
}
