package triangle

import (
	"go-rpg/internal/common"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const width, height = 800, 600

func Main() {
	window := common.NewWindow(width, height)
	common.InitGlow()
	defer glfw.Terminate()

	// 1. Bind the vao
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Push all the vertex info inside vao:
	// 1.a. Gen and load vertex data
	data := []float32{
		-0.5, -0.5, 0,
		-0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(data)*4,
		gl.Ptr(data),
		gl.STATIC_DRAW,
	)

	// 1.b. Specify the vertex attribute data
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)

	// 2. Load, compile and link program
	vsSource := common.LoadShaderSource("./internal/triangle/vertexShader.vs")
	fsSource := common.LoadShaderSource("./internal/triangle/fragmentShader.fs")
	program := common.NewProgram(vsSource, fsSource)

	for !window.ShouldClose() {

		gl.UseProgram(program)
		gl.BindVertexArray(vao)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}

}
