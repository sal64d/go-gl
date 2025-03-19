package home

import (
	renderer "go-rpg/internal/Renderer"
	"go-rpg/internal/common"
	"go-rpg/internal/primitives"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const width = 1280
const height = 800

type GameState struct {
	playerPosition mgl32.Vec3
	isDoorOpen     bool
}

func Main() {
	window := common.NewWindow(width, height)
	common.InitGlow()
	defer glfw.Terminate()

	GroudTexture := renderer.Texture{
		Filepath: "./internal/textures/forrest_ground_03_diff_1k.jpg",
		Wrap_s:   gl.MIRRORED_REPEAT,
		Wrap_r:   gl.MIRRORED_REPEAT,
		Opacity:  1.0,
	}

	BrickTexture := renderer.Texture{
		Filepath: "./internal/textures/red_brick_diff_1k.jpg",
		Wrap_s:   gl.MIRRORED_REPEAT,
		Wrap_r:   gl.MIRRORED_REPEAT,
		Opacity:  1.0,
	}

	groundMaterial := renderer.Material{
		ShaderType: renderer.Lambert,
		Color:      mgl32.Vec4{0.2, 0.2, 0.2, 1},
		Texture:    renderer.TextureMap{renderer.Diffuse: GroudTexture},
	}

	wallMaterial := renderer.Material{
		ShaderType: renderer.Lambert,
		Color:      mgl32.Vec4{1.0, 0.2, 0.2, 1},
		Texture:    renderer.TextureMap{renderer.Diffuse: BrickTexture},
	}

  ground := primitives.CreatePlane(2, 2, 2, 2, groundMaterial)
	ground.TransformMesh(mgl32.HomogRotate3DX(mgl32.DegToRad(-90)))

	cube := primitives.CreateCube2(.5, 1, .5, 1.5, 1.5, wallMaterial)
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
		r.SceneGL.CameraGL.ViewMatrix = mgl32.Translate3D(0, -0.2, -3).Mul4(
			mgl32.HomogRotate3DX(float32(mgl32.DegToRad(30))),
		).Mul4(
			mgl32.HomogRotate3DY(float32(angle)),
		)
		r.Draw()
	}
}
