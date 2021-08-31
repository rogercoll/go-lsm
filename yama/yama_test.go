package yama

import (
	"io/ioutil"
	"testing"
)

func TestIsEnabled(t *testing.T) {
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
			f, err := ioutil.TempFile("", "testyamafile")
			if err != nil {
				t.Error(err)
			}
			ioutil.WriteFile(f.Name(), []byte(tt.fileContent), 0644)
			ptrace_scope_file = f.Name()
			enabled, err := IsEnabled()
			if err != nil {
				t.Error(err)
			}
			if enabled != tt.out {
				t.Errorf("got %t, want %t", enabled, tt.out)
			}
		})
	}
}
