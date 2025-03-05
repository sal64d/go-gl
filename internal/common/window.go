package common

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func NewWindow(width int, height int) *glfw.Window {
	err := glfw.Init()

	if err != nil {
		log.Fatalln("newWindow: failed to init glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "glfw", nil, nil)
	if err != nil {
		log.Panicln("newWindow: create window failed", err)
	}
	window.MakeContextCurrent()

	return window

}

func InitGlow() {
	// Init Glow
	if err := gl.Init(); err != nil {
		log.Panicln("main: gl init failed", err)
	}

	// print version
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("main: opengl version", version)

}

func SetupScene(program uint32, width float32, height float32) {
	gl.UseProgram(program)

	// Global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0, 0, 0, 0)

	// Projection matrix
	projection := mgl32.Perspective(
		mgl32.DegToRad(45.0),
		width/height,
		0.1,
		10.0,
	)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	log.Println("setup: projectionUniform\n", projection)

	// Camera
	camera := mgl32.LookAtV(
		mgl32.Vec3{3, 3, 3},
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{0, 1, 0},
	)
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	log.Println("setup: cameraUniform\n", camera)

}
