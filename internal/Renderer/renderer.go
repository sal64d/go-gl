package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func Render(scene Scene, window *glfw.Window) {
	proj := mgl32.Perspective(
		mgl32.DegToRad(scene.Camera.FOV),
		scene.Camera.AspectRatio,
		0.1,
		100.0,
	)

	view := mgl32.Translate3D(0, -0.3, -3).Mul4(
		mgl32.HomogRotate3DX(mgl32.DegToRad(30)),
	)

	shaderMap := CompileShaders(proj)

	sceneGL := GenerateGLData(scene)

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

		turnTableView := view.Mul4(mgl32.HomogRotate3DY(float32(angle)))

		for i, modelGL := range sceneGL.ModelsGL {
			for j, meshGL := range modelGL.MeshesGL {
				Draw(elapsedTime, meshGL, scene.Models[i].Meshes[j].Material, shaderMap, turnTableView)
			}
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func Draw(elapsedTime float64, meshGL MeshGL, material Material, shaderMap ShaderMap, view mgl32.Mat4) {
	shader := shaderMap[material.ShaderType]

	gl.UseProgram(shader.Program)

	// apply the material color to shader
	gl.Uniform4fv(shader.uniforms[MatColor], 1, &material.Color[0])
	gl.UniformMatrix4fv(shader.uniforms[ViewMatrix], 1, false, &view[0])

	gl.BindVertexArray(meshGL.VAO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, meshGL.EBO)

	gl.DrawElements(
		gl.TRIANGLES,
		meshGL.Size,
		gl.UNSIGNED_INT,
		nil,
	)
}

func GenerateGLData(scene Scene) (sceneGL SceneGL) {
	ModelsGL := []ModelGL{}

	for _, model := range scene.Models {

		MeshesGL := []MeshGL{}

		for _, mesh := range model.Meshes {
			var VAO uint32

			gl.GenVertexArrays(1, &VAO)
			gl.BindVertexArray(VAO)

			vertexData := []float32{}
			for _, v := range mesh.Vertices {
				vertexData = append(vertexData, v.X())
				vertexData = append(vertexData, v.Y())
				vertexData = append(vertexData, v.Z())
			}

			var VBO uint32
			gl.GenBuffers(1, &VBO)
			gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
			gl.BufferData(
				gl.ARRAY_BUFFER,
				len(vertexData)*4,
				gl.Ptr(vertexData),
				gl.STATIC_DRAW,
			)

			var EBO uint32
			gl.GenBuffers(1, &EBO)
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
			gl.BufferData(
				gl.ELEMENT_ARRAY_BUFFER,
				len(mesh.Indices)*4,
				gl.Ptr(mesh.Indices),
				gl.STATIC_DRAW,
			)

			gl.VertexAttribPointerWithOffset(
				0,
				3,
				gl.FLOAT,
				false,
				3*4,
				0,
			)
			gl.EnableVertexAttribArray(0)

			Size := int32(len(mesh.Indices))
			MeshesGL = append(MeshesGL, MeshGL{VAO, VBO, EBO, Size})
		}

		ModelsGL = append(ModelsGL, ModelGL{MeshesGL})
	}

	sceneGL.ModelsGL = ModelsGL
	return
}
