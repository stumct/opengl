package main

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Game struct {
	Width    int
	Height   int
	VAO      uint32
	Program  uint32
	Texture1 uint32
	Texture2 uint32
	MixValue float32
	Cubes    []mgl32.Vec3
}

func NewGame(width, height int) *Game {
	return &Game{Width: width, Height: height}
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
	//fmt.Println(texture1)
	texture2, err := NewTexture("assets/awesomeface.png")
	if err != nil {
		log.Fatalln(err)
	}
	game.Texture2 = texture2
	//fmt.Println(texture2)

	// Configure the vertex data
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	/*var EBO uint32
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicesRect)*4, gl.Ptr(indicesRect), gl.STATIC_DRAW)
	*/
	/*	var VBO uint32
		gl.GenBuffers(1, &VBO)
		gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
		gl.BufferData(gl.ARRAY_BUFFER, len(verticesRectTex)*4, gl.Ptr(verticesRectTex), gl.STATIC_DRAW)
	*/
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticesCube)*4, gl.Ptr(verticesCube), gl.STATIC_DRAW)
	// Coords Attributes
	//gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	//gl.EnableVertexAttribArray(0)

	// RGB Attributes
	//gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	//gl.EnableVertexAttribArray(1)

	// Texture Attributes
	//gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	//gl.EnableVertexAttribArray(2)

	// Coords Attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Texture Attributes
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(2)

	// unbind the VAO
	gl.BindVertexArray(0)

	game.VAO = VAO
	game.Program = program

	game.Cubes = []mgl32.Vec3{
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
	//gl.Uniform1f(gl.GetUniformLocation(game.Program, gl.Str("mixValue\x00")), game.MixValue)
	//ident := mgl32.Ident4()
	//model := mgl32.HomogRotate3D(mgl32.DegToRad(float32(glfw.GetTime()*50.0)), mgl32.Vec3{0, 1, 1})
	//model := mgl32.HomogRotate3D(mgl32.DegToRad(float32(glfw.GetTime()*50.0)), mgl32.Vec3{0, 1, 1})
	model1 := mgl32.HomogRotate3DX(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
	model2 := mgl32.HomogRotate3DY(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
	model := model1.Mul4(model2)
	view := mgl32.Translate3D(0, 0, -4.0)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), 800/600, 0.1, 100.0)

	//gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("model\x00")), 1, false, &model[0])
	gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("view\x00")), 1, false, &view[0])
	gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("projection\x00")), 1, false, &projection[0])

	gl.BindVertexArray(game.VAO)
	for _, val := range game.Cubes {

		position := mgl32.Translate3D(val.X(), val.Y(), val.Z())
		model1 := mgl32.HomogRotate3DX(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
		model2 := mgl32.HomogRotate3DY(mgl32.DegToRad(float32(glfw.GetTime() * 50.0)))
		model = model1.Mul4(model2)
		model = model.Mul4(position)
		gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("model\x00")), 1, false, &model[0])
		//gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("view\x00")), 1, false, &view[0])
		//gl.UniformMatrix4fv(gl.GetUniformLocation(game.Program, gl.Str("projection\x00")), 1, false, &projection[0])
		gl.DrawArrays(gl.TRIANGLES, 0, 36)
	}
	gl.BindVertexArray(0)
	/*
		var transRotate mgl32.Mat4

		transRotate = mgl32.HomogRotate3D(1.5708, mgl32.Vec3{0.0, 0.0, 1.0})

		var transScale mgl32.Mat4
		transScale = mgl32.Scale3D(0.5, 0.5, 0.5)

		var trans mgl32.Mat4

		trans = transRotate.Mul4(transScale)
	*/
	//ident := mgl32.Ident4()
	//scale := mgl32.Scale3D(0.5, 0.5, 0.5)
	//transformLoc := gl.GetUniformLocation(game.Program, gl.Str("transform\x00"))
	//gl.UniformMatrix4fv(transformLoc, 1, false, &scale[0])
	/*
		time := glfw.GetTime()
		fmt.Println(time)

		var translate mgl32.Mat4
		translate = mgl32.Translate3D(0.5, -0.5, 0.0)
		var rotate mgl32.Mat4
		rotate = mgl32.HomogRotate3D(mgl32.DegToRad(float32(glfw.GetTime()*50.0)), mgl32.Vec3{0.0, 0.0, 1.0})
		var trans mgl32.Mat4
		trans = translate.Mul4(rotate)

		transformLoc := gl.GetUniformLocation(game.Program, gl.Str("transform\x00"))
		gl.UniformMatrix4fv(transformLoc, 1, false, &trans[0])
	*/

	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	//gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(nil))

}

var vertexShaderSource = `
#version 330 core
layout (location = 0) in vec3 position;
layout (location = 1) in vec3 color;
layout (location = 2) in vec2 texCoord;

out vec3 ourColor;
out vec2 TexCoord;

//uniform mat4 transform;
uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main()
{
    //gl_Position = projection * view * model * vec4(position, 1.0f);
    //gl_Position = projection * model * vec4(position, 1.0f);
    //gl_Position = model * vec4(position, 1.0f);
    //gl_Position = vec4(position, 1.0f);

    gl_Position = projection * view * model * vec4(position, 1.0f);
	
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
