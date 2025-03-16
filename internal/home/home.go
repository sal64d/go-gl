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
		Color:      mgl32.Vec4{0.2, 0.2, 0.2, 1},
	}

	wallMaterial := renderer.Material{
		ShaderType: renderer.Lambert,
		Color:      mgl32.Vec4{1.0, 0.2, 0.2, 1},
	}

	ground := primitives.CreatePlane(2, 2, groundMaterial)
	ground.TransformMesh(mgl32.HomogRotate3DX(mgl32.DegToRad(-90)))

	cube := primitives.CreateCube(.5, .5, .5, wallMaterial)
	cube.TransformMesh(mgl32.Translate3D(0, .25, 0))

	scene := renderer.Scene{
		Models: []renderer.Model{
			{
				Meshes: []renderer.Mesh{ground},
			},
			{
				Meshes: []renderer.Mesh{cube},
			},
		},
		Camera: renderer.Camera{
			FOV:         45,
			AspectRatio: width / height,
		},
	}

	r := renderer.CreateRenderer(scene, window)

	angle := 0.0
	for r.Step() {
		angle += r.State.TimeDelta
		r.SceneGL.CameraGL.ViewMatrix = mgl32.Translate3D(0, -0.5, -3).Mul4(
			mgl32.HomogRotate3DY(float32(angle)),
		)
		r.Draw()
	}
}
