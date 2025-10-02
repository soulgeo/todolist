package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/soulgeo/todolist/internal/app"
	"github.com/soulgeo/todolist/internal/store"
	"github.com/soulgeo/todolist/internal/todo"
)

// returns a sane default path for the lists.json file
func defaultDataPath() (string, error) {
	// 1) env override
	if p := os.Getenv("TODO_DATA"); p != "" {
		return p, nil
	}

	// 2) platform-specific fallbacks
	if runtime.GOOS == "windows" {
		// UserConfigDir is usually something like: C:\Users\<user>\AppData\Local
		if cfg, err := os.UserConfigDir(); err == nil && cfg != "" {
			return filepath.Join(cfg, "todolist", "lists.json"), nil
		}
	} else {
		// default to ~/.local/share/todolist/lists.json
		home, err := os.UserHomeDir()
		if err == nil && home != "" {
			return filepath.Join(home, ".local", "share", "todolist", "lists.json"), nil
		}
	}

	// last resort: current directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, "lists.json"), nil
}

func main() {
	// compute path
	dataPath, err := defaultDataPath()
	if err != nil {
		log.Fatalf("couldn't determine data path: %v", err)
	}

	// ensure parent dir exists (use 0700/0750 depending on needs)
	if err := os.MkdirAll(filepath.Dir(dataPath), 0o700); err != nil {
		log.Fatalf("failed to create data dir: %v", err)
	}

	// create store (NewJSONStore should accept a path string and return todo.Store)
	jsonstore := store.NewJSONStore(dataPath)

	// wire service and start app
	app.SetService(todo.NewService(jsonstore))
	app.Execute()
}
