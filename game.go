package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
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
	prog, err := NewShaderProgram("./shaders/vertex_basic.vert", "./shaders/frag_basic.frag")
	if err != nil {
		panic(err)
	}
	game.Program = prog

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
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*floatSize, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Coords Attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(3*floatSize), gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

}

func (game *Game) Render() {
	gl.UseProgram(game.Program)
	gl.BindVertexArray(game.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
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
