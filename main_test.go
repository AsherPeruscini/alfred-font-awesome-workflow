package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("FAW_ICONS_YAML_PATH", "workflow/icons.yml")
	os.Exit(m.Run())
}

func setTestEnvHelper(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}

func TestMain_newApp(t *testing.T) {
	actual := newApp().Version

	expected := "0.1.0"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
