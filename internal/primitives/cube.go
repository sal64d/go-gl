package primitives

import (
	renderer "go-rpg/internal/Renderer"

	"github.com/go-gl/mathgl/mgl32"
)

func CreateCube(width float32, height float32, depth float32, material renderer.Material) renderer.Mesh {
	bottomLeft := mgl32.Vec3([]float32{0, 0, 0})  // 0 4
	bottomRight := mgl32.Vec3([]float32{1, 0, 0}) // 1 5
	topLeft := mgl32.Vec3([]float32{0, 1, 0})     // 2 6
	topRight := bottomRight.Add(topLeft)          // 3 7

	back := mgl32.Vec3([]float32{0, 0, 1})

	frontFace := []mgl32.Vec3{
		bottomLeft,
		bottomRight,
		topLeft,
		topRight,
	}

	faceUV := []mgl32.Vec2{
		bottomLeft.Vec2(),
		bottomRight.Vec2(),
		topLeft.Vec2(),
		topRight.Vec2(),
	}

	var combinedVecs [8]mgl32.Vec3
	var uv [8]mgl32.Vec2

	for i := range frontFace {
		combinedVecs[i] = frontFace[i]
		combinedVecs[i+4] = frontFace[i].Add(back)
		uv[i] = faceUV[i]
	}
	uv[4] = faceUV[1]
	uv[5] = faceUV[0]
	uv[6] = faceUV[3]
	uv[7] = faceUV[2]

	scaler := mgl32.Scale3D(width, height, depth)

	for i, v := range combinedVecs {
		// center it
		centerd := v.Add(mgl32.Vec3{-0.5, -0.5, -0.5})
		scaled := scaler.Mul4x1(centerd.Vec4(0))

		combinedVecs[i] = scaled.Vec3()
	}

	indices := []int32{
		0, 1, 2, // front face
		2, 1, 3,

		4, 5, 6, // back face
		6, 5, 7,

		0, 4, 2, // left face
		2, 4, 6,

		1, 5, 3, // right face
		3, 5, 7,

		2, 3, 6, // top face
		6, 3, 7,

		0, 1, 4, // bottom face
		4, 1, 5,
	}

	return renderer.Mesh{
		Indices:  indices,
		Vertices: combinedVecs[:],
		UV:       uv[:],
		Material: material,
	}
}
