package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {

	// Initialise GLFW
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Provide window hints for GLFW
	glfw.WindowHint(glfw.ContextVersionMajor, 3)                // OpenGL Version 3.3
	glfw.WindowHint(glfw.ContextVersionMinor, 3)                // OpenGL Version 3.3
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // use the core version
	glfw.WindowHint(glfw.Resizable, glfw.False)                 // disable window resizing

	// Create the window object
	window, err := glfw.CreateWindow(800, 600, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	// Make the context of our window the main context on the current thread
	window.MakeContextCurrent()

	if runtime.GOOS == "windows" {
		window.SetPos(32, 64)
	}

	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Get the width and heigh of the window from GLFW
	width, height := window.GetFramebufferSize()
	// Set the location of the lower left corner of the window.
	// And the width and height of the rendering window in pixels
	gl.Viewport(0, 0, int32(width), int32(height))

	// Setup OpenGL options
	gl.Enable(gl.DEPTH_TEST)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	///////////////////////////////////////////
	game := NewGame(width, height)
	game.Setup()
	/////////////////////////////////////////////

	// Key callback function to handle key press
	// We register the callback functions after we've created the window and before the game loop is initiated.
	window.SetKeyCallback(keycallback(game))

	// Game Loop
	for !window.ShouldClose() {
		// Check and call events
		glfw.PollEvents()

		// Rendering
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.Clear(gl.DEPTH_BUFFER_BIT)

		// Run the main game render method
		game.Render()

		// Swap the buffers
		window.SwapBuffers()

	}
}
func keycallback(game *Game) func(*glfw.Window, glfw.Key, int, glfw.Action, glfw.ModifierKey) {
	return func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		// When a user presses the escape key, we set the WindowShouldClose property to true,
		// closing the application
		if key == glfw.KeyEscape && action == glfw.Press {
			w.SetShouldClose(true)
		}

		if key == glfw.KeyUp && action == glfw.Press {
			//fmt.Println("KeyUp")
			if game.MixValue < 1.0 {
				game.MixValue = game.MixValue + 0.1
			}
		}

		if key == glfw.KeyDown && action == glfw.Press {
			//fmt.Println("KeyDown")
			if game.MixValue > 0.0 {
				game.MixValue = game.MixValue - 0.1
			}
		}
	}
}
