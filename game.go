package main

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Game struct {
	VAO      uint32
	Program  uint32
	Texture1 uint32
	Texture2 uint32
	MixValue float32
}

func NewGame() *Game {
	return &Game{}
}

func (game *Game) Setup() {

	// Configure the vertex and fragment shaders
	program, err := NewShaderProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		panic(err)
	}

	// Load Textures
	texture1, err := NewTexture("assets/container.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	game.Texture1 = texture1
	fmt.Println(texture1)
	texture2, err := NewTexture("assets/awesomeface.png")
	if err != nil {
		log.Fatalln(err)
	}
	game.Texture2 = texture2
	fmt.Println(texture2)

	// Configure the vertex data
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	var EBO uint32
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicesRect)*4, gl.Ptr(indicesRect), gl.STATIC_DRAW)

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticesRectTex)*4, gl.Ptr(verticesRectTex), gl.STATIC_DRAW)

	// Coords Attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// RGB Attributes
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Texture Attributes
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

	// unbind the VAO
	gl.BindVertexArray(0)

	game.VAO = VAO
	game.Program = program

	/*
		vec := glm.Vec4{1.0, 0.0, 0.0, 1.0}

			var trans glm.Mat4

			trans = glm.Translate3D(1.0, 1.0, 0.0)

			newvec := trans.Mul4x1(vec)
			fmt.Println(newvec)
	*/
}

func (game *Game) Render() {
	gl.UseProgram(game.Program)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, game.Texture1)
	gl.Uniform1i(gl.GetUniformLocation(game.Program, gl.Str("ourTexture1\x00")), 0)

	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, game.Texture2)
	gl.Uniform1i(gl.GetUniformLocation(game.Program, gl.Str("ourTexture2\x00")), 1)

	// Set current value of uniform mix
	gl.Uniform1f(gl.GetUniformLocation(game.Program, gl.Str("mixValue\x00")), game.MixValue)

	gl.BindVertexArray(game.VAO)
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	//gl.DrawArrays(gl.TRIANGLES, 0, 4)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(nil))
	gl.BindVertexArray(0)
}

var vertexShaderSource = `
#version 330 core
layout (location = 0) in vec3 position;
layout (location = 1) in vec3 color;
layout (location = 2) in vec2 texCoord;

out vec3 ourColor;
out vec2 TexCoord;

void main()
{
    gl_Position = vec4(position, 1.0f);
    ourColor = color;
    // We swap the y-axis by substracing our coordinates from 1. This is done because most images have the top y-axis inversed with OpenGL's top y-axis.
	// TexCoord = texCoord;
	TexCoord = vec2(texCoord.x, 1.0 - texCoord.y);
}` + "\x00"

var fragmentShaderSource = `
#version 330 core
in vec3 ourColor;
in vec2 TexCoord;

out vec4 color;

uniform float mixValue;

uniform sampler2D ourTexture1;
uniform sampler2D ourTexture2;

void main()
{
    color = mix(texture(ourTexture1, TexCoord), texture(ourTexture2, TexCoord), mixValue);
}` + "\x00"

var vertices = []float32{
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
