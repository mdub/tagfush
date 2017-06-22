package update

import (
	"bytes"
	"strings"
	"testing"

	"github.com/docker/distribution/reference"
	"github.com/opencontainers/go-digest"
)

type stubClient struct {
	digest digest.Digest
}

func (c *stubClient) Init(debug bool) error {
	return nil
}

func (c *stubClient) GetTags(ref reference.Named) ([]string, error) {
	return []string{}, nil
}

func (c *stubClient) Resolve(name reference.Named) (reference.Canonical, error) {
	return reference.WithDigest(name, c.digest)
}

func TestUpdateRefsInStream(t *testing.T) {

	oldDigest := "sha256:a5ebd3bc0bf3881258975f8afa1c6d24429dfd4d7dd53a299559a3e927b77fd7"
	newDigest := "sha256:08868d719684cf9cafacbaa1786ad01111332b4c1e65abd67833db603d8dab7f"
	input := "line1\nruby:2.3@" + oldDigest + "\nline 3\n"
	expectedOutput := strings.Replace(input, oldDigest, newDigest, 1)
	expectedReport := "-:2: ruby:2.3\n  was " + oldDigest + "\n  now " + newDigest + "\n"

	client := &stubClient{digest: digest.Digest(newDigest)}
	inputReader := strings.NewReader(input)
	outputWriter := new(bytes.Buffer)
	reportWriter := new(bytes.Buffer)
	u := NewUpdater(client, reportWriter)

	err := u.UpdateRefsInStream("-", inputReader, outputWriter)
	if err != nil {
		t.Error("Did not expect error, ", err)
	}

	output := outputWriter.String()
	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}

	report := reportWriter.String()
	if report != expectedReport {
		t.Errorf("expected report %q, got %q", expectedReport, report)
	}

}
