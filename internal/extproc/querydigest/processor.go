package querydigest

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/kaz/pprotein/internal/collect"
)

type (
	processor struct{}
)

func (p *processor) Cacheable() bool {
	return true
}

func (p *processor) Process(snapshot *collect.Snapshot) (io.ReadCloser, error) {
	cmd := exec.Command("pt-query-digest", "--limit", "100%", "--output", "json", snapshot.Body)

	res, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("external process aborted: %w", err)
	}

	return io.NopCloser(bytes.NewBuffer(res)), nil
}