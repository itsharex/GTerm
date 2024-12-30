package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/OpenToolkitLab/GTerm/backend/cmd"
	"github.com/OpenToolkitLab/GTerm/backend/consts"
	"github.com/OpenToolkitLab/GTerm/backend/pkg/base"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	app := cmd.NewApp()

	if err := wails.Run(&options.App{
		Title:     consts.ProjectName,
		Width:     1266,
		Height:    768,
		Frameless: !app.PreferencesSrv.IsDarwin(),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)
		},
		Bind: app.Bind(),
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableFramelessWindowDecorations: false,
			BackdropType:                      windows.Tabbed,
		},
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   fmt.Sprintf("%s %s", consts.ProjectName, base.Version),
				Message: "Copyright © 2023 OpenToolkitLab All rights reserved",
				Icon:    icon,
			},
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
	}); err != nil {
		panic(err)
	}
}
