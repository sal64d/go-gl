package primitives

import (
	renderer "go-rpg/internal/Renderer"

	"github.com/go-gl/mathgl/mgl32"
)

func CreateCube2(
	width float32,
	height float32,
	depth float32,
	scale_u float32,
	scale_v float32,
	material renderer.Material,
) renderer.Mesh {

	frontBottomLeft := mgl32.Vec3{0, 0, 0}  // 0 4
	frontBottomRight := mgl32.Vec3{1, 0, 0} // 1 5
	frontTopLeft := mgl32.Vec3{0, 1, 0}     // 2 6
	frontTopRight := mgl32.Vec3{1, 1, 0}    // 3 7

	backBottomLeft := mgl32.Vec3{0, 0, 1}  // 0 4
	backBottomRight := mgl32.Vec3{1, 0, 1} // 1 5
	backTopLeft := mgl32.Vec3{0, 1, 1}     // 2 6
	backTopRight := mgl32.Vec3{1, 1, 1}    // 3 7

	vertices := []mgl32.Vec3{
		frontBottomLeft, frontTopLeft, frontTopRight,
		frontTopRight, frontBottomRight, frontBottomLeft,

		frontBottomRight, frontTopRight, backTopRight,
		backTopRight, backBottomRight, frontBottomRight,

		backBottomRight, backTopRight, backTopLeft,
		backTopLeft, backBottomLeft, backBottomRight,

		backBottomLeft, backTopLeft, frontTopLeft,
		frontTopLeft, frontBottomLeft, backBottomLeft,

		frontTopLeft, backTopLeft, backTopRight,
		backTopRight, frontTopRight, frontTopLeft,

		backBottomLeft, frontBottomLeft, frontBottomRight,
		frontBottomRight, backBottomRight, backBottomLeft,
	}

	topLeft := mgl32.Vec2{0, 1}
	topRight := mgl32.Vec2{1, 1}
	bottomLeft := mgl32.Vec2{0, 0}
	bottomRight := mgl32.Vec2{1, 0}

	uv := []mgl32.Vec2{
		bottomLeft, topLeft, topRight,
		topRight, bottomRight, bottomLeft,

		bottomLeft, topLeft, topRight,
		topRight, bottomRight, bottomLeft,

		bottomLeft, topLeft, topRight,
		topRight, bottomRight, bottomLeft,

		bottomLeft, topLeft, topRight,
		topRight, bottomRight, bottomLeft,

		bottomLeft, topLeft, topRight,
		topRight, bottomRight, bottomLeft,

		bottomLeft, topLeft, topRight,
		topRight, bottomRight, bottomLeft,
	}

	vertScaler := mgl32.Scale3D(width, height, depth)
	uvScaler := mgl32.Scale2D(scale_u, scale_v)

	for i := range vertices {
		// center and scale vert
		centerd := vertices[i].Add(mgl32.Vec3{-0.5, -0.5, -0.5})
		scaled := vertScaler.Mul4x1(centerd.Vec4(0))
		vertices[i] = scaled.Vec3()

		// scale uv
		uv[i] = uvScaler.Mul3x1(uv[i].Vec3(0)).Vec2()
	}

	return renderer.Mesh{
		Vertices:     vertices,
		UV:           uv,
		Material:     material,
		FeatureFlags: renderer.EnableUV,
	}

}

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
		Indices:      indices,
		Vertices:     combinedVecs[:],
		UV:           uv[:],
		Material:     material,
		FeatureFlags: renderer.EnableEBO | renderer.EnableUV,
	}
}
