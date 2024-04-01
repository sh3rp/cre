package model

import (
	"fmt"
	"github.com/sh3rp/cre/internal/model"
)

var DefaultDownloadURL = "https://go.dev/dl"

func (gp GoPlatform) SrcURL() string {
	return fmt.Sprintf("%s/go%d.%d.%d.src.%s",
		DefaultDownloadURL,
		gp.Version.Major,
		gp.Version.Minor,
		gp.Version.Patch,
		gp.Extension)
}

func (gp GoPlatform) BinaryURL() string {
	return fmt.Sprintf("%s/go%d.%d.%d.%s-%s.%s",
		DefaultDownloadURL,
		gp.Version.Major,
		gp.Version.Minor,
		gp.Version.Patch,
		gp.Goos,
		gp.Goarch,
		gp.Extension)
}

func (gp GoPlatform) ToConfig() *model.ConfigObject {
	return nil
}

type GoPlatform struct {
	Version   GoVersion `json:"version"`
	Goos      string    `json:"os"`
	Goarch    string    `json:"arch"`
	Extension string    `json:"extension"`
}

type GoVersion struct {
	Major int    `json:"major"`
	Minor int    `json:"minor"`
	Patch int    `json:"patch"`
	Label string `json:"label"`
}
