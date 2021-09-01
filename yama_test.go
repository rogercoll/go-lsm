package lsm

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestYAMAIsEnabled(t *testing.T) {
	var yamaFileTest = []struct {
		name        string
		fileContent string
		out         bool
	}{
		{"Not enabled, with default value", "0\n", false},
		{"Enabled, restricted", "1\n", true},
	}
	for _, tt := range yamaFileTest {
		t.Run(tt.name, func(t *testing.T) {
			totalFile := strings.Split(ptrace_scope_file, "/")
			tmpDir := os.TempDir() + strings.Join(totalFile[:len(totalFile)-1], "/")
			err := os.MkdirAll(tmpDir, 0777)
			if err != nil {
				t.Error(err)
			}
			f, err := os.Create(os.TempDir() + ptrace_scope_file)
			if err != nil {
				t.Error(err)
			}
			defer os.Remove(f.Name()) // clean up
			ioutil.WriteFile(f.Name(), []byte(tt.fileContent), 0644)
			lsm, err := NewLSMConfig("./", os.TempDir())
			if err != nil {
				t.Error(err)
			}
			enabled, err := lsm.IsYamaEnabled()
			if err != nil {
				t.Error(err)
			}
			if enabled != tt.out {
				t.Errorf("got %t, want %t", enabled, tt.out)
			}
		})
	}
}
