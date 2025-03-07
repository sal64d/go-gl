package triangle

import (
	"go-rpg/internal/common"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const width, height = 800, 600

func attachModelToProgram(program uint32) int32 {

	gl.UseProgram(program)

	modelUniform := gl.GetUniformLocation(
		program,
		gl.Str("model\x00"),
	)

	return (modelUniform)
}

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
	vertices := []float32{
		0.5, 0.5, 0, // top right
		-0.5, 0.5, 0, // top left
		0.5, -0.5, 0, // bottom right
		-0.5, -0.5, 0, // bottom left
	}
	indices := []int32{
		0, 1, 2,
		1, 2, 3,
	}

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(
		gl.ELEMENT_ARRAY_BUFFER,
		len(indices)*4,
		gl.Ptr(indices),
		gl.STATIC_DRAW,
	)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(vertices)*4,
		gl.Ptr(vertices),
		gl.STATIC_DRAW,
	)

	// 1.b. Specify the vertex attribute data
	gl.VertexAttribPointerWithOffset(
		0,
		3,
		gl.FLOAT,
		false,
		3*4,
		0,
	)
	gl.EnableVertexAttribArray(0)

	// 2. Load, compile and link program
	vsSource := common.LoadShaderSource("./internal/triangle/vertexShader.vs")
	fsSource1 := common.LoadShaderSource("./internal/triangle/fragmentShader1.fs")
	fsSource2 := common.LoadShaderSource("./internal/triangle/fragmentShader2.fs")
	program1 := common.NewProgram(vsSource, fsSource1)
	program2 := common.NewProgram(vsSource, fsSource2)

	// 3. model
	model := mgl32.Ident4()

	model1 := attachModelToProgram(program1)
	model2 := attachModelToProgram(program2)

	gl.UniformMatrix4fv(model1, 1, false, &model[0])
	gl.UniformMatrix4fv(model2, 1, false, &model[0])

	angle := 0.0
	prevTime := glfw.GetTime()

	for !window.ShouldClose() {

		gl.ClearColor(0, 0, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		currTime := glfw.GetTime()
		elapsed := currTime - prevTime
		prevTime = currTime

		angle += elapsed / 2

		model := mgl32.HomogRotate3D(
			float32(angle),
			mgl32.Vec3{0, 1, 0},
		)

		if int(angle) % 10  > 5 {
			gl.UseProgram(program2)
		} else {
			gl.UseProgram(program1)
		}
		gl.UniformMatrix4fv(model1, 1, false, &model[0])
		gl.BindVertexArray(vao)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

		gl.DrawElements(
			gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil,
		)

		window.SwapBuffers()
		glfw.PollEvents()
	}

}
