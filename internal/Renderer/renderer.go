package renderer

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func CreateRenderer(scene Scene, Window *glfw.Window) Renderer {
	ShaderMap := CompileShaders()

	SceneGL := SceneGL{
		ModelsGL: generateModelGLData(scene.Models, ShaderMap),
		CameraGL: generateCameraGL(scene.Camera),
	}

	State := RendererState{
		TimeDelta: 0,
		Time:      glfw.GetTime(),
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1)

	return Renderer{
		SceneGL,
		State,
		Window,
	}
}

func (self *Renderer) Step() bool {
	if self.Window.ShouldClose() {
		return false
	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	time := glfw.GetTime()
	self.State.TimeDelta = time - self.State.Time
	self.State.Time = time

	return true
}

func (self *Renderer) Draw() {
	for _, modelGL := range self.SceneGL.ModelsGL {
		for _, meshGL := range modelGL.MeshesGL {
			self.DrawMesh(
				meshGL,
				modelGL.ModelMatrix,
			)
		}
	}

	self.Window.SwapBuffers()
	glfw.PollEvents()

}

func (self *Renderer) DrawMesh(
	meshGL MeshGL,
	modelMatrix mgl32.Mat4,
) {
	material := meshGL.Material
	shader := material.ShaderProg

	shader.UseProgram()

	// Color
	gl.Uniform4fv(
		shader.GetUniformLocation(MatColor),
		1,
		&material.Color[0],
	)

	// Projection
	gl.UniformMatrix4fv(
		shader.GetUniformLocation(ProjectionMatrix),
		1,
		false,
		&self.SceneGL.CameraGL.ProjectionMatrix[0],
	)

	// View
	gl.UniformMatrix4fv(
		shader.GetUniformLocation(ViewMatrix),
		1,
		false,
		&self.SceneGL.CameraGL.ViewMatrix[0],
	)

	// Model
	gl.UniformMatrix4fv(
		shader.GetUniformLocation(ModelMatrix),
		1,
		false,
		&modelMatrix[0],
	)

	// Diffuse texture
	diff, ok := material.TextureGLMap[Diffuse]
	if ok {
		diff.Bind(gl.TEXTURE0)
		gl.Uniform1i(
			shader.GetUniformLocation(MatDiffTex),
			int32(diff.Handle-gl.TEXTURE0),
		)

		gl.Uniform1f(
			shader.GetUniformLocation(MatDiffOpacity),
			diff.Opacity,
		)
	}

	gl.BindVertexArray(meshGL.VAO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, meshGL.EBO)

	gl.DrawElements(
		gl.TRIANGLES,
		meshGL.Size,
		gl.UNSIGNED_INT,
		nil,
	)
}

func generateModelGLData(models []Model, shaderMap ShaderMap) []ModelGL {
	ModelsGL := []ModelGL{}

	for _, model := range models {

		MeshesGL := []MeshGL{}

		for _, mesh := range model.Meshes {
			var VAO uint32

			gl.GenVertexArrays(1, &VAO)
			gl.BindVertexArray(VAO)

			vertexData := []float32{}
			for i, v := range mesh.Vertices {
				// Loc
				vertexData = append(vertexData, v.X())
				vertexData = append(vertexData, v.Y())
				vertexData = append(vertexData, v.Z())

				// UV
				vertexData = append(vertexData, mesh.UV[i].X())
				vertexData = append(vertexData, mesh.UV[i].Y())
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

			stride := int32(3*4 + 2*4)

			// Loc
			gl.VertexAttribPointerWithOffset(
				0,
				3,
				gl.FLOAT,
				false,
				stride,
				0,
			)
			gl.EnableVertexAttribArray(0)

			// Tex
			gl.VertexAttribPointerWithOffset(
				1,
				2,
				gl.FLOAT,
				false,
				stride,
				3*4,
			)
			gl.EnableVertexAttribArray(1)

			size := int32(len(mesh.Indices))

			// Load the textures
			var textureGLMap TextureGLMap = make(TextureGLMap)
			for key, texture := range mesh.Material.Texture {
				textureGLMap[key] = texture.Load()
			}

			// Load shader
			shaderProg, ok := shaderMap[mesh.Material.ShaderType]
			if !ok {
				log.Panicln("Shader undefined", mesh.Material.ShaderType)
			}

			// Build material
			materialGL := MaterialGL{
				shaderProg,
				mesh.Material.Color,
				textureGLMap,
			}

			MeshesGL = append(MeshesGL, MeshGL{VAO, VBO, EBO, size, materialGL})
		}

		ModelMatrix := mgl32.Ident4()
		ModelsGL = append(ModelsGL, ModelGL{MeshesGL, ModelMatrix})
	}

	return ModelsGL
}

func generateCameraGL(camera Camera) (cameraGL CameraGL) {
	cameraGL.ProjectionMatrix = mgl32.Perspective(
		mgl32.DegToRad(camera.FOV),
		camera.AspectRatio,
		0.1,
		100.0,
	)

	cameraGL.ViewMatrix = mgl32.Translate3D(0, -0.3, -3).Mul4(
		mgl32.HomogRotate3DX(mgl32.DegToRad(30)),
	)

	return

}
