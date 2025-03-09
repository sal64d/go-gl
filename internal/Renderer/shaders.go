package renderer

import (
	"go-rpg/internal/common"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func CompileShaders(defaultProjection mgl32.Mat4) ShaderMap {
	lambertVs := common.LoadShaderSource("./internal/shaders/lambert.vs")
	lambertFs := common.LoadShaderSource("./internal/shaders/lambert.fs")

	// create the material
	mat := common.NewProgram(lambertVs, lambertFs)

	// register uniform
	colorUniform := gl.GetUniformLocation(mat, gl.Str(MatColor))
	defaultColor := mgl32.Vec4{.2, .2, .2, 1}

	projUniform := gl.GetUniformLocation(mat, gl.Str(ProjectionMatrix))

	viewUniform := gl.GetUniformLocation(mat, gl.Str(ViewMatrix))
	defaultView := mgl32.Ident4()

	modelUniform := gl.GetUniformLocation(mat, gl.Str(ModelMatrix))
	defaultModel := mgl32.Ident4()

	gl.UseProgram(mat)
	gl.Uniform4fv(colorUniform, 1, &defaultColor[0])
	gl.UniformMatrix4fv(projUniform, 1, false, &defaultProjection[0])
	gl.UniformMatrix4fv(viewUniform, 1, false, &defaultView[0])
	gl.UniformMatrix4fv(modelUniform, 1, false, &defaultModel[0])

	return ShaderMap{
		Lambert: Shader{
			Program: mat,
			uniforms: UniformMap{
				ProjectionMatrix: projUniform,
				ViewMatrix:       viewUniform,
				ModelMatrix:      modelUniform,
				MatColor:         colorUniform,
			},
		},
	}
}
