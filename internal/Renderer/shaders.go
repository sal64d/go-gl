package renderer

import (
	"go-rpg/internal/common"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func CompileShaders() ShaderMap {
	lambertVs := common.LoadShaderSource("./internal/shaders/lambert.vs")
	lambertFs := common.LoadShaderSource("./internal/shaders/lambert.fs")

	// create the material
	shader := Shader{
		Program: common.NewProgram(lambertVs, lambertFs),
	}

	shader.UseProgram()

	defaultColor := mgl32.Vec4{.2, .2, .2, 1}
	gl.Uniform4fv(
		shader.GetUniformLocation(MatColor),
		1,
		&defaultColor[0],
	)

	return ShaderMap{
		Lambert: shader,
	}
}

func (self *Shader) UseProgram() {
	gl.UseProgram(self.Program)
}

func (self *Shader) GetUniformLocation(name string) int32 {
	return gl.GetUniformLocation(self.Program, gl.Str(name+"\x00"))
}
