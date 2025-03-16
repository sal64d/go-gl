package renderer

import "github.com/go-gl/mathgl/mgl32"

func (a *Mesh) TransformMesh(transformMatrix mgl32.Mat4) {
	for i, v := range a.Vertices {
		a.Vertices[i] = transformMatrix.Mul4x1(mgl32.Vec4{v.X(), v.Y(), v.Z(), 1.0}).Vec3()
	}
}
