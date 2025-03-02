package main

import (
	"go-rpg/internal/common"
	"go-rpg/internal/cube"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const width, height = 640, 480

func init() {
	runtime.LockOSThread()
}

func main() {
	log.Printf("main: start")

	// Init glfw and glow
	window := common.NewWindow(width, height)
	common.InitGlow()
	defer glfw.Terminate()

	// Init program
	vertexShader, fragmentShader := common.LoadShaderSource("./internal/cube/vertexShader.vs"), common.LoadShaderSource("./internal/cube/fragmentShader.fs")
	program := common.NewProgram(vertexShader, fragmentShader)

	modelUniform, vao := setupScene(program)

	angle := 0.0
	previousTime := glfw.GetTime()

	for !window.ShouldClose() {

		currentTime := glfw.GetTime()
		elapsed := currentTime - previousTime
		previousTime = currentTime

		angle += elapsed

		drawScene(angle, program, modelUniform, vao)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func setupScene(program uint32) (modelUniform int32, vao uint32) {
	log.Println("setup: started")

	common.SetupScene(program, width, height)

	// model
	model := mgl32.Ident4()
	modelUniform = gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	// Load texture
	//texture, err := newTexture("square.png")
	//if err != nil {
	//	log.Fatalln("setup: Error while loading texture")
	//}

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(cube.CubeVertices)*4,
		gl.Ptr(cube.CubeVertices),
		gl.STATIC_DRAW,
	)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointerWithOffset(texCoordAttrib, 2, gl.FLOAT, false, 5*4, 3*4)

	return
}

func drawScene(angle float64, program uint32, modelUniform int32, vao uint32) {
	// log.Println("draw: started")
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	model := mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

	// Render
	gl.UseProgram(program)
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	gl.BindVertexArray(vao)

	// gl.ActiveTexture(gl.TEXTURE0)
	// gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
