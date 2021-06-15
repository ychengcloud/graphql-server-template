package assets

import (
	"embed"
)

//go:embed assets/*
var Assets embed.FS

//go:embed index/*
var GraphiqlTmpl embed.FS
