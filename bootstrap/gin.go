package bootstrap

import (
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
)

func newGinEntry(boot []byte) *rkgin.GinEntry {
	rkentry.BootstrapBuiltInEntryFromYAML(boot)
	rkentry.GlobalAppCtx.GetAppInfoEntry().AppName = "pass-through"

	res := rkgin.RegisterGinEntryYAML(boot)

	return res["pass-through"].(*rkgin.GinEntry)
}
