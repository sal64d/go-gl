package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func Render(scene Scene, window *glfw.Window) {
	shaderMap := CompileShaders()

	sceneGL := GenerateGLData(scene)

	prevTime := glfw.GetTime()

	gl.Enable(gl.DEPTH_TEST)

	for !window.ShouldClose() {
		gl.ClearColor(0, 0, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		currTime := glfw.GetTime()
		elapsedTime := currTime - prevTime
		prevTime = currTime

		for i, modelGL := range sceneGL.ModelsGL {
			for j, meshGL := range modelGL.MeshesGL {
				Draw(elapsedTime, meshGL, scene.Models[i].Meshes[j].Material, shaderMap)
			}
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func Draw(elapsedTime float64, meshGL MeshGL, material Material, shaderMap ShaderMap) {
	shader := shaderMap[material.ShaderType]

	gl.UseProgram(shader.Program)

  // apply the material color to shader
  gl.Uniform4fv(shader.uniforms[MatColor], 1, &material.Color[0])

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
      for _,v := range mesh.Vertices {
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
