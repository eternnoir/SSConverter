package ssconverter

import (
	"bytes"
)

type Converter interface {
	CheckCommand() bool
	BuildSite() (sitePath string, err error)
	GetSiteBytes(archiveType string) (buffer *bytes.Buffer, err error)
}