package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// calculate the memory size of floats used to calculate total memory size of float arrays
const floatSize = 4

type Game struct {
	Width    int
	Height   int
	VAO      uint32
	Program  uint32
	Texture1 uint32
	Texture2 uint32
	Textures map[string]uint32
	MixValue float32
	Cubes    []mgl32.Vec3
}

func NewGame(width, height int) *Game {
	return &Game{
		Width:    width,
		Height:   height,
		Textures: map[string]uint32{},
	}
}

func (game *Game) Setup() {

	// Configure the vertex and fragment shaders
	prog, err := NewShaderProgram("./shaders/basic_tex.vert", "./shaders/basic_tex.frag")
	if err != nil {
		panic(err)
	}
	game.Program = prog

	// Load the textures
	tex, err := LoadTextures("./texture")
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
	gl.UseProgram(game.Program)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, game.Textures["container.jpg"])
	gl.Uniform1i(gl.GetUniformLocation(game.Program, gl.Str("texture1\x00")), 0)

	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, game.Textures["awesomeface.png"])
	gl.Uniform1i(gl.GetUniformLocation(game.Program, gl.Str("texture2\x00")), 1)

	gl.BindVertexArray(game.VAO)

	for _, pos := range positions {
		model0 := mgl32.Translate3D(pos.X(), pos.Y(), pos.Z())
		model1 := mgl32.HomogRotate3DX(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
		model2 := mgl32.HomogRotate3DY(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
		model := model0.Mul4(model1).Mul4(model2)
		view := mgl32.Translate3D(0, 0, -4.0)

		projection := mgl32.Perspective(mgl32.DegToRad(45.0), 800/600, 0.1, 100.0)

		gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("model\x00")), 1, false, &model[0])
		gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("view\x00")), 1, false, &view[0])
		gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("projection\x00")), 1, false, &projection[0])

		gl.DrawArrays(gl.TRIANGLES, 0, 36)
	}

	//gl.BindTexture(gl.TEXTURE_2D, game.Textures["container.jpg"])
	//gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(nil))
	gl.BindVertexArray(0)
}

var vertices = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	0.0, 0.5, 0.0,
}

var vertices2 = []float32{
	-0.5, -0.5, 0.0, 1.0, 0.0, 0.0,
	0.5, -0.5, 0.0, 0.0, 1.0, 0.0,
	0.0, 0.5, 0.0, 0.0, 0.0, 1.0,
}

var verticesRect = []float32{
	// Positions          // Colors
	0.5, 0.5, 0.0, 1.0, 0.0, 0.0, // Top Right
	0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // Bottom Right
	-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, // Bottom Left
	-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, // Top Left
}

var indicesRect = []int32{
	0, 1, 3, // First Triangle
	1, 2, 3, // Second Triangle
}
var verticesRectTex = []float32{
	// Positions          // Colors           // Texture Coords
	0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, // Top Right
	0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, // Bottom Right
	-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, // Bottom Left
	-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, // Top Left
}

var verticesCube = []float32{
	-0.5, -0.5, -0.5, 0.0, 0.0,
	0.5, -0.5, -0.5, 1.0, 0.0,
	0.5, 0.5, -0.5, 1.0, 1.0,
	0.5, 0.5, -0.5, 1.0, 1.0,
	-0.5, 0.5, -0.5, 0.0, 1.0,
	-0.5, -0.5, -0.5, 0.0, 0.0,

	-0.5, -0.5, 0.5, 0.0, 0.0,
	0.5, -0.5, 0.5, 1.0, 0.0,
	0.5, 0.5, 0.5, 1.0, 1.0,
	0.5, 0.5, 0.5, 1.0, 1.0,
	-0.5, 0.5, 0.5, 0.0, 1.0,
	-0.5, -0.5, 0.5, 0.0, 0.0,

	-0.5, 0.5, 0.5, 1.0, 0.0,
	-0.5, 0.5, -0.5, 1.0, 1.0,
	-0.5, -0.5, -0.5, 0.0, 1.0,
	-0.5, -0.5, -0.5, 0.0, 1.0,
	-0.5, -0.5, 0.5, 0.0, 0.0,
	-0.5, 0.5, 0.5, 1.0, 0.0,

	0.5, 0.5, 0.5, 1.0, 0.0,
	0.5, 0.5, -0.5, 1.0, 1.0,
	0.5, -0.5, -0.5, 0.0, 1.0,
	0.5, -0.5, -0.5, 0.0, 1.0,
	0.5, -0.5, 0.5, 0.0, 0.0,
	0.5, 0.5, 0.5, 1.0, 0.0,

	-0.5, -0.5, -0.5, 0.0, 1.0,
	0.5, -0.5, -0.5, 1.0, 1.0,
	0.5, -0.5, 0.5, 1.0, 0.0,
	0.5, -0.5, 0.5, 1.0, 0.0,
	-0.5, -0.5, 0.5, 0.0, 0.0,
	-0.5, -0.5, -0.5, 0.0, 1.0,

	-0.5, 0.5, -0.5, 0.0, 1.0,
	0.5, 0.5, -0.5, 1.0, 1.0,
	0.5, 0.5, 0.5, 1.0, 0.0,
	0.5, 0.5, 0.5, 1.0, 0.0,
	-0.5, 0.5, 0.5, 0.0, 0.0,
	-0.5, 0.5, -0.5, 0.0, 1.0,
}

var positions = []mgl32.Vec3{
	mgl32.Vec3{0.0, 0.0, 0.0},
	mgl32.Vec3{2.0, 5.0, -15.0},
	mgl32.Vec3{-1.5, -2.2, -2.5},
	mgl32.Vec3{-3.8, -2.0, -12.3},
	mgl32.Vec3{2.4, -0.4, -3.5},
	mgl32.Vec3{-1.7, 3.0, -7.5},
	mgl32.Vec3{1.3, -2.0, -2.5},
	mgl32.Vec3{1.5, 2.0, -2.5},
	mgl32.Vec3{1.5, 0.2, -1.5},
	mgl32.Vec3{-1.3, 1.0, -1.5},
}
