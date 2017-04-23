package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// calculate the memory size of floats used to calculate total memory size of float arrays
const floatSize = 4

type Game struct {
	Width  int
	Height int
	VAO    uint32
	Camera *Camera

	ShaderPrograms map[string]uint32
	Textures       map[string]uint32
	InputKeys      map[glfw.Key]bool

	MixValue float32
	Cubes    []mgl32.Vec3

	deltaTime float32
	lastFrame float32
}

func NewGame(width, height int, camera *Camera) *Game {
	return &Game{
		Width:          width,
		Height:         height,
		Camera:         camera,
		ShaderPrograms: map[string]uint32{},
		Textures:       map[string]uint32{},
		InputKeys:      map[glfw.Key]bool{},
	}
}

func (game *Game) Setup() {

	// Configure the vertex and fragment shaders
	prog, err := NewShaderProgram("./shaders/basic_tex.vert", "./shaders/basic_tex.frag")
	if err != nil {
		panic(err)
	}
	game.ShaderPrograms["BasicTextureShaders"] = prog

	// Load the textures
	tex, err := LoadTextures("./textures")
	if err != nil {
		panic(err)
	}
	game.Textures = tex

	// Configure the vertex data
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	game.VAO = VAO

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticesCube)*floatSize, gl.Ptr(verticesCube), gl.STATIC_DRAW)

	/*var EBO uint32
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicesRect)*floatSize, gl.Ptr(indicesRect), gl.STATIC_DRAW)*/

	// Coords Attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(5*floatSize), gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	// Color attribute
	/*gl.VertexAttribPointer(1, 3, gl.FLOAT, false, int32(8*floatSize), gl.PtrOffset(3*floatSize))
	gl.EnableVertexAttribArray(1)*/
	// TexCoord attribute
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, int32(5*floatSize), gl.PtrOffset(3*floatSize))
	gl.EnableVertexAttribArray(2)

	gl.BindVertexArray(0)

}

func (game *Game) Render() {

	game.UpdateTimes(float32(glfw.GetTime()))
	game.UpdateCameraPosition()

	gl.UseProgram(game.ShaderPrograms["BasicTextureShaders"])

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, game.Textures["container.jpg"])
	gl.Uniform1i(gl.GetUniformLocation(game.ShaderPrograms["BasicTextureShaders"], gl.Str("texture1\x00")), 0)

	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, game.Textures["awesomeface.png"])
	gl.Uniform1i(gl.GetUniformLocation(game.ShaderPrograms["BasicTextureShaders"], gl.Str("texture2\x00")), 1)

	gl.BindVertexArray(game.VAO)

	for _, pos := range positions {
		model0 := mgl32.Translate3D(pos.X(), pos.Y(), pos.Z())
		model1 := mgl32.HomogRotate3DX(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
		model2 := mgl32.HomogRotate3DY(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
		model := model0.Mul4(model1).Mul4(model2)
		//view := mgl32.Translate3D(0, 0, -4.0)
		//radius := 10.0
		//camX := math.Sin(glfw.GetTime()) * radius
		//camZ := math.Cos(glfw.GetTime()) * radius
		//view = glm::lookAt(glm::vec3(camX, 0.0, camZ), glm::vec3(0.0, 0.0, 0.0), glm::vec3(0.0, 1.0, 0.0));
		//view := mgl32.LookAtV(mgl32.Vec3{float32(camX), 0.0, float32(camZ)}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{0.0, 1.0, 0.0})
		view := game.Camera.CurrentView()

		projection := mgl32.Perspective(mgl32.DegToRad(float32(game.Camera.FOV)), 800/600, 0.1, 100.0)

		gl.UniformMatrix4fv(gl.GetUniformLocation(game.ShaderPrograms["BasicTextureShaders"], gl.Str("model\x00")), 1, false, &model[0])
		gl.UniformMatrix4fv(gl.GetUniformLocation(game.ShaderPrograms["BasicTextureShaders"], gl.Str("view\x00")), 1, false, &view[0])
		gl.UniformMatrix4fv(gl.GetUniformLocation(game.ShaderPrograms["BasicTextureShaders"], gl.Str("projection\x00")), 1, false, &projection[0])

		gl.DrawArrays(gl.TRIANGLES, 0, 36)
	}

	//gl.BindTexture(gl.TEXTURE_2D, game.Textures["container.jpg"])
	//gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(nil))
	gl.BindVertexArray(0)
}

func (game *Game) UpdateTimes(time float32) {
	game.deltaTime = time - game.lastFrame
	game.lastFrame = time
}
func (game *Game) UpdateCameraPosition() {
	game.Camera.SetDelta(float32(game.deltaTime))
	if game.InputKeys[glfw.KeyW] {
		game.Camera.MoveForward()
	}
	if game.InputKeys[glfw.KeyS] {
		game.Camera.MoveBackward()
	}
	if game.InputKeys[glfw.KeyA] {
		game.Camera.MoveLeft()
	}
	if game.InputKeys[glfw.KeyD] {
		game.Camera.MoveRight()
	}
}

func (game *Game) CursorEventHandler() func(w *glfw.Window, xpos float64, ypos float64) {
	return func(w *glfw.Window, xpos float64, ypos float64) {
		game.Camera.HandleCursorEvent(xpos, ypos)
	}
}

func (game *Game) ScrollEventHandler() func(w *glfw.Window, xoff float64, yoff float64) {
	return func(w *glfw.Window, xoff float64, yoff float64) {
		game.Camera.HandleScrollEvent(xoff, yoff)
	}
}

func (game *Game) KeyEventHandler() func(*glfw.Window, glfw.Key, int, glfw.Action, glfw.ModifierKey) {
	return func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		// When a user presses the escape key, we set the WindowShouldClose property to true,
		// closing the application
		if key == glfw.KeyEscape && action == glfw.Press {
			w.SetShouldClose(true)
		}

		if action == glfw.Press {
			game.InputKeys[key] = true
		} else if action == glfw.Release {
			game.InputKeys[key] = false
		}
	}
}
