package lsm

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestYAMAIsActive(t *testing.T) {
	var yamaFileTest = []struct {
		name        string
		fileContent string
		out         bool
	}{
		{"Not enabled, with default value", "0\n", false},
		{"Active, restricted", "1\n", true},
	}
	for _, tt := range yamaFileTest {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.CreateTemp(os.TempDir(), "yamatest")
			if err != nil {
				t.Error(err)
			}
			defer os.Remove(f.Name()) // clean up
			err = ioutil.WriteFile(f.Name(), []byte(tt.fileContent), 0644)
			if err != nil {
				t.Error(err)
			}
			lsm, err := NewLSMConfig(ConfigPath{YamaScope: f.Name()})
			if err != nil {
				t.Error(err)
			}
			enabled, err := lsm.IsYamaActive()
			if err != nil {
				t.Error(err)
			}
			if enabled != tt.out {
				t.Errorf("got %t, want %t", enabled, tt.out)
			}
		})
	}
}
