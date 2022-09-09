package main

import (
	"embed"
	_ "embed"

	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"

	"pass-trougth/bootstrap"
)

//go:embed cia-config/docs
var docsFS embed.FS

func init() {
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.SWEntryType, "pass-through", &docsFS)
}

//go:embed boot.yaml
var boot []byte

func main() {
	bootstrap.Run(boot)
}
