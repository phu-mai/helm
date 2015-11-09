package action

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPluginName(t *testing.T) {
	if PluginName("foo") != "helm-foo" {
		t.Errorf("Expected helm-foo, got %s", PluginName("foo"))
	}
}

func TestPlugin(t *testing.T) {
	f := "../testdata"
	p := "plugin"
	a := []string{"-a", "-b", "-c"}

	os.Setenv("PATH", os.ExpandEnv("$PATH:"+helmRoot+"/testdata"))

	buf, err := ioutil.TempFile("", "helm-plugin-test")
	if err != nil {
		t.Fatal(err)
	}

	oldout := os.Stdout
	os.Stdout = buf
	defer func() { os.Stdout = oldout; buf.Close(); os.Remove(buf.Name()) }()

	fakeUpdate(f)
	Plugin(f, p, a)

	buf.Seek(0, 0)
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Errorf("Failed to read tmp file: %s", err)
	}

	if strings.TrimSpace(string(b)) != "HELLO -a -b -c" {
		t.Errorf("Expected 'HELLO -a -b -c', got %v", string(b))
	}
}
