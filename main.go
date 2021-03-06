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
	// make sure that we display any errors that are encountered
	//glfw.SetErrorCallback(errorCallback)

	// Initialise GLFW
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Provide window hints for GLFW
	glfw.WindowHint(glfw.Samples, 4)                            // desired number of samples to use for mulitsampling
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

	// disable v-sync for max FPS if the driver allows it
	glfw.SwapInterval(0)

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
	camera := NewDefaultCamera()
	camera.SetSpeed(5.00)
	game := NewGame(width, height, camera)
	game.Setup()
	/////////////////////////////////////////////

	// Key callback function to handle key press
	// We register the callback functions after we've created the window and before the game loop is initiated.
	window.SetKeyCallback(game.KeyEventHandler())
	window.SetCursorPosCallback(game.CursorEventHandler())
	window.SetScrollCallback(game.ScrollEventHandler())
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

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

// handle GLFW errors by printing them out
func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}
