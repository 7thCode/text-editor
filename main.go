package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	// File メニュー
	fileMenu := menu.NewMenu()
	fileMenu.AddText("新規作成", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "new-file")
	})
	fileMenu.AddText("開く...", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "open-file")
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("保存", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "save-file")
	})
	fileMenu.AddText("名前を付けて保存...", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "save-file-as")
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("終了", nil, func(cd *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Edit メニュー
	editMenu := menu.NewMenu()
	editMenu.AddText("元に戻す", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "undo")
	})
	editMenu.AddSeparator()
	editMenu.AddText("切り取り", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "cut")
	})
	editMenu.AddText("コピー", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "copy")
	})
	editMenu.AddText("貼り付け", nil, func(cd *menu.CallbackData) {
		runtime.EventsEmit(app.ctx, "paste")
	})

	// メインメニュー
	mainMenu := menu.NewMenu()
	mainMenu.Append(menu.SubMenu("ファイル", fileMenu))
	mainMenu.Append(menu.SubMenu("編集", editMenu))

	err := wails.Run(&options.App{
		Title:  "テキストエディター",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Menu:             mainMenu,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
