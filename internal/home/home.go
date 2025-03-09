package home

import (
	renderer "go-rpg/internal/Renderer"
	"go-rpg/internal/common"
	"go-rpg/internal/primitives"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const width = 800
const height = 600

// ground
// + grass texture

// wasd movement
// mouse movement

// simple flat house with
// - 4 walls with intersection checks
// - 3 windows on the wall
// - a door that opens / closes on interaction

// Scene
// models
// meshes
// texturs

type GameState struct {
	playerPosition mgl32.Vec3
	isDoorOpen     bool
}

func Main() {
	window := common.NewWindow(width, height)
	common.InitGlow()
	defer glfw.Terminate()

	groundMaterial := renderer.Material{
		ShaderType: renderer.Lambert,
		Color:      mgl32.Vec4{1, 0, 0, 1},
	}
	mesh := primitives.CreatePlane(.5, .5, groundMaterial)
	mesh.TransformMesh(mgl32.HomogRotate3DX(mgl32.DegToRad(45)))

	scene := renderer.Scene{Models: []renderer.Model{
		{
			Meshes: []renderer.Mesh{mesh},
		},
	}}

	renderer.Render(scene, window)
}

//gl.UseProgram(program)

//projection := mgl32.Perspective(mgl32.DegToRad(45), width/height, 0.1, 100)
//model := mgl32.Ident4()
//view := mgl32.Translate3D(0, -0.0, -3.0).Mul4(mgl32.HomogRotate3DX(mgl32.DegToRad(30)))

//modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
//viewUniform := gl.GetUniformLocation(program, gl.Str("view\x00"))
//projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))

//gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
//gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])
//gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

//prevTime := glfw.GetTime()
//angle := 0.0

//for !window.ShouldClose() {
//	gl.ClearColor(0, 0, 0, 1)
//	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

//	currTime := glfw.GetTime()
//	elapsedTime := currTime - prevTime
//	prevTime = currTime

//	angle += elapsedTime

//	gl.UseProgram(program)

//	model := mgl32.HomogRotate3D(
//		float32(angle),
//		mgl32.Vec3{0, 1, 0},
//	)
//	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

//	gl.BindVertexArray(vao)
//	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

//	gl.DrawElements(
//		gl.TRIANGLES,
//		int32(len(indices)),
//		gl.UNSIGNED_INT,
//		nil,
//	)

//	window.SwapBuffers()
//	glfw.PollEvents()
//}
