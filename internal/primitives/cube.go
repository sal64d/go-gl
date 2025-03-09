package primitives

import "github.com/go-gl/mathgl/mgl32"

func CreateCube(width float32, height float32, depth float32) ([]float32, []int32) {
	bottomLeft := mgl32.Vec3([]float32{0, 0, 0})      // 0 4
	bottomRight := mgl32.Vec3([]float32{width, 0, 0}) // 1 5
	topLeft := mgl32.Vec3([]float32{0, height, 0})    // 2 6
	topRight := bottomRight.Add(topLeft)              // 3 7

	back := mgl32.Vec3([]float32{0, 0, depth})

	frontFace := []mgl32.Vec3{bottomLeft, bottomRight, topLeft, topRight}

	var combinedVecs [8]mgl32.Vec3
	for i := range frontFace {
		combinedVecs[i] = frontFace[i]
		combinedVecs[i+4] = frontFace[i].Add(back)
	}

	var vertexData []float32
	for _, v := range combinedVecs {
    // Pos
		vertexData = append(vertexData, v.X()-width/2)
		vertexData = append(vertexData, v.Y()-height/2)
		vertexData = append(vertexData, v.Z()-depth/2)

    // Color
    vertexData = append(vertexData, v.X()/2 + 0.2)
    vertexData = append(vertexData, v.Y()/2 + 0.2)
    vertexData = append(vertexData, v.Z()/2 + 0.2)
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

	return vertexData, indices
}
