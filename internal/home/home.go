package home

import (
	"go-rpg/internal/common"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func createCube(width float32, height float32, depth float32) ([]float32, []int32) {
	bottomLeft := mgl32.Vec3([]float32{0, 0, 0})      // 0 4
	bottomRight := mgl32.Vec3([]float32{width, 0, 0}) // 1 5
	topLeft := mgl32.Vec3([]float32{0, height, 0})    // 2 6
	topRight := bottomRight.Add(topLeft)              // 3 7

	back := mgl32.Vec3([]float32{0, 0, depth})

	frontFace := []mgl32.Vec3{bottomLeft, bottomRight, topLeft, topRight}

	var combinedVecs [8]mgl32.Vec3
	for i := range frontFace {
		combinedVecs[i] = frontFace[i]
		combinedVecs[i+4] = frontFace[i].Add(back)
	}

	var vertexData []float32
	for _, v := range combinedVecs {
    // Pos
		vertexData = append(vertexData, v.X()-width/2)
		vertexData = append(vertexData, v.Y()-height/2)
		vertexData = append(vertexData, v.Z()-depth/2)

    // Color
    vertexData = append(vertexData, v.X()/2 + 0.2)
    vertexData = append(vertexData, v.Y()/2 + 0.2)
    vertexData = append(vertexData, v.Z()/2 + 0.2)
	}

  log.Println("vertex data", vertexData[12:18])

	indices := []int32{
		0, 1, 2, // front face
		2, 1, 3,

		4, 5, 6, // back face
		6, 5, 7,

		0, 4, 2, // left face
		2, 4, 6,

		1, 5, 3, // right face
		3, 5, 7,

		2, 3, 6, // top face
		6, 3, 7,

		0, 1, 4, // bottom face
		4, 1, 5,
	}

	return vertexData, indices
}

const width = 800
const height = 600

func Main() {
	window := common.NewWindow(width, height)
	common.InitGlow()
	defer glfw.Terminate()

	vertexData, indices := createCube(0.5, 1, 0.8)

	vertexShader := common.LoadShaderSource("./internal/home/vertexShader.vs")
	fragmentShader := common.LoadShaderSource("./internal/home/fragmentShader.fs")
	program := common.NewProgram(vertexShader, fragmentShader)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(vertexData)*4,
		gl.Ptr(vertexData),
		gl.STATIC_DRAW,
	)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(
		gl.ELEMENT_ARRAY_BUFFER,
		len(indices)*4,
		gl.Ptr(indices),
		gl.STATIC_DRAW,
	)

	gl.VertexAttribPointerWithOffset(
		0,
		3,
		gl.FLOAT,
		false,
		6*4,
		0,
	)
	gl.EnableVertexAttribArray(0)
  
  gl.VertexAttribPointerWithOffset(
		1,
		3,
		gl.FLOAT,
		false,
		6*4,
		3*4,
	)
	gl.EnableVertexAttribArray(1)

	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(45), width/height, 0.1, 100)
	model := mgl32.Ident4()
	view := mgl32.Translate3D(0, -0.0, -3.0).Mul4(mgl32.HomogRotate3DX(mgl32.DegToRad(30)))

	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	viewUniform := gl.GetUniformLocation(program, gl.Str("view\x00"))
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))

	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	prevTime := glfw.GetTime()
	angle := 0.0

  gl.Enable(gl.DEPTH_TEST)

	for !window.ShouldClose() {
		gl.ClearColor(0, 0, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		currTime := glfw.GetTime()
		elapsedTime := currTime - prevTime
		prevTime = currTime

		angle += elapsedTime

		gl.UseProgram(program)

		model := mgl32.HomogRotate3D(
			float32(angle),
			mgl32.Vec3{0, 1, 0},
		)
		gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

		gl.BindVertexArray(vao)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

		gl.DrawElements(
			gl.TRIANGLES,
			int32(len(indices)),
			gl.UNSIGNED_INT,
			nil,
		)

		window.SwapBuffers()
		glfw.PollEvents()
	}

}
