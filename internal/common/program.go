package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func NewProgram(vertexShaderSource string, fragmentShaderSource string) uint32 {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		log.Panicln("program: failed to compile vertex shader", err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		log.Panicln("program: failed to compile fragment shader", err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		msg := strings.Repeat("\x00", int(logLength+1))

		log.Panicln("program: failed to link program", msg)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)

	free()

	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))

		return 0, fmt.Errorf("shader: failed to compile shader: %v: %v", source, log)
	}

	return shader, nil

}

func LoadShaderSource(filepath string) string {
	buf, err := os.ReadFile(filepath)
	if err != nil {
		log.Panicf("load: unable to load the shader file %v: %v", filepath, err)
	}

	return string(buf) + "\x00"
}
