package repository

import (
	"github.com/kcwebapply/bm/provider"
)

var (
	fileName    = ""
	maxTextSize = 60

	contentPath = ""
)

func init() {
	fileName = provider.FileName
	contentPath = provider.ContentPath
}
