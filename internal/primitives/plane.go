package primitives

import (
	renderer "go-rpg/internal/Renderer"

	"github.com/go-gl/mathgl/mgl32"
)

func CreatePlane(width float32, height float32, material renderer.Material) renderer.Mesh {
	repos := mgl32.Vec3{-width / 2, -height / 2, 0}
	return renderer.Mesh{
		Vertices: []mgl32.Vec3{
			mgl32.Vec3{0, 0, 0}.Add(repos),
			mgl32.Vec3{width, 0, 0}.Add(repos),
			mgl32.Vec3{0, height, 0}.Add(repos),
			mgl32.Vec3{width, height, 0}.Add(repos),
		},
		Indices: []int32{
			0, 1, 2,
			2, 1, 3,
		},
		Material: material,
	}
}
