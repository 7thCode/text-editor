package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// FileData represents file information
type FileData struct {
	Content string `json:"content"`
	Path    string `json:"path"`
	Name    string `json:"name"`
}

// NewFile creates a new empty file
func (a *App) NewFile() FileData {
	return FileData{
		Content: "",
		Path:    "",
		Name:    "新規ファイル",
	}
}

// OpenFileDialog opens a file dialog and returns the selected file content
func (a *App) OpenFileDialog() (FileData, error) {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "ファイルを開く",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "テキストファイル (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Markdownファイル (*.md)",
				Pattern:     "*.md",
			},
			{
				DisplayName: "JavaScriptファイル (*.js)",
				Pattern:     "*.js",
			},
			{
				DisplayName: "Pythonファイル (*.py)",
				Pattern:     "*.py",
			},
			{
				DisplayName: "HTMLファイル (*.html)",
				Pattern:     "*.html",
			},
			{
				DisplayName: "CSSファイル (*.css)",
				Pattern:     "*.css",
			},
			{
				DisplayName: "JSONファイル (*.json)",
				Pattern:     "*.json",
			},
			{
				DisplayName: "すべてのファイル (*.*)",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		return FileData{}, err
	}

	if filePath == "" {
		return FileData{}, fmt.Errorf("ファイルが選択されませんでした")
	}

	return a.OpenFile(filePath)
}

// OpenFile opens a file and returns its content
func (a *App) OpenFile(filePath string) (FileData, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return FileData{}, err
	}

	return FileData{
		Content: string(content),
		Path:    filePath,
		Name:    filepath.Base(filePath),
	}, nil
}

// SaveFile saves content to the specified file path
func (a *App) SaveFile(filePath string, content string) error {
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}

// SaveFileDialog opens a save file dialog and saves the content
func (a *App) SaveFileDialog(content string) (FileData, error) {
	filePath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "名前を付けて保存",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "テキストファイル (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Markdownファイル (*.md)",
				Pattern:     "*.md",
			},
			{
				DisplayName: "JavaScriptファイル (*.js)",
				Pattern:     "*.js",
			},
			{
				DisplayName: "Pythonファイル (*.py)",
				Pattern:     "*.py",
			},
			{
				DisplayName: "HTMLファイル (*.html)",
				Pattern:     "*.html",
			},
			{
				DisplayName: "CSSファイル (*.css)",
				Pattern:     "*.css",
			},
			{
				DisplayName: "JSONファイル (*.json)",
				Pattern:     "*.json",
			},
			{
				DisplayName: "すべてのファイル (*.*)",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		return FileData{}, err
	}

	if filePath == "" {
		return FileData{}, fmt.Errorf("ファイル名が指定されませんでした")
	}

	return a.SaveFileAs(filePath, content)
}

// SaveFileAs saves content to a new file path
func (a *App) SaveFileAs(filePath string, content string) (FileData, error) {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return FileData{}, err
	}

	return FileData{
		Content: content,
		Path:    filePath,
		Name:    filepath.Base(filePath),
	}, nil
}

// GetHomeDir returns the user's home directory
func (a *App) GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return homeDir
}


