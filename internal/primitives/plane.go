package primitives

import (
	renderer "go-rpg/internal/Renderer"

	"github.com/go-gl/mathgl/mgl32"
)

func CreatePlane(
	width float32,
	height float32,
	scale_u float32,
	scale_v float32,
	material renderer.Material,
) renderer.Mesh {
	repos := mgl32.Vec3{-0.5, -0.5, 0}

	base := []mgl32.Vec3{
		mgl32.Vec3{0, 0, 0}.Add(repos),
		mgl32.Vec3{1, 0, 0}.Add(repos),
		mgl32.Vec3{0, 1, 0}.Add(repos),
		mgl32.Vec3{1, 1, 0}.Add(repos),
	}

	scaler := mgl32.Scale2D(width, height)
	uvScaler := mgl32.Scale2D(scale_u, scale_v)

	var scaled [4]mgl32.Vec3
	var uv [4]mgl32.Vec2
	for i, v := range base {
		scaled[i] = scaler.Mul3x1(v)
		uv[i] = uvScaler.Mul3x1(v).Vec2()
	}

	return renderer.Mesh{
		Vertices: scaled[:],
		Indices: []int32{
			0, 1, 2,
			2, 1, 3,
		},
		UV:           uv[:],
		Material:     material,
		FeatureFlags: renderer.EnableUV | renderer.EnableEBO,
	}
}
