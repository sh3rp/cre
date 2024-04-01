package model

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGoPlatform(t *testing.T) {
	p := GoPlatform{
		Goarch:    "arm64",
		Goos:      "darwin",
		Extension: "tar.gz",
		Version: GoVersion{
			Major: 1,
			Minor: 22,
			Patch: 1,
			Label: "1.22",
		},
	}
	bytes, _ := json.MarshalIndent(p, "", "  ")
	os.WriteFile("go_version.json", bytes, 0700)

}
