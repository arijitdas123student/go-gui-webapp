package main

import (
	"github.com/Equanox/gotron"
)

func main() {
	// Create a new browser window instance
	window, err := gotron.New("webapp")
	if err != nil {
		panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 640
	window.WindowOptions.Height = 480
	window.WindowOptions.Title = "Golang GUI App"

	// Start the browser window.
	// This will establish a golang <=> nodejs bridge using websockets,
	// to control ElectronBrowserWindow with our window object.
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// Open dev tools must be used after window.Start
	// window.OpenDevTools()

	// Wait for the application to close
	<-done
}
