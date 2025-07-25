package tempfile_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/buchgr/bazel-remote/v2/utils/tempfile"
)

func TestTempfileCreator(t *testing.T) {
	tfc := tempfile.NewCreator()

	dir, err := os.MkdirTemp("", "foo")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.RemoveAll(dir) }()

	targetFileBase := path.Join(dir, "foo")
	tf, random, err := tfc.Create(targetFileBase, false)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Remove(tf.Name()) }()

	if random == "" {
		t.Fatalf("Expected non-empty random string in the filename: %q",
			tf.Name())
	}

	if !strings.Contains(tf.Name(), random) {
		t.Fatalf("Expected filename %q to contain random string %q",
			tf.Name(), random)
	}

	expectedPrefix := targetFileBase
	if !strings.HasPrefix(tf.Name(), expectedPrefix) {
		t.Fatalf("Expected tempfile \"%s\" to have prefix \"%s\"",
			tf.Name(), expectedPrefix)
	}
}
