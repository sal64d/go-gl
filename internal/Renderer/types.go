package renderer

import "github.com/go-gl/mathgl/mgl32"

type Scene struct {
	Models []Model
	Camera Camera
}

type Camera struct {
	FOV         float32
	AspectRatio float32
}

type Model struct {
	Meshes []Mesh
}

type Mesh struct {
	Vertices []mgl32.Vec3
	Indices  []int32
	Material Material
}

type SceneGL struct {
	ModelsGL []ModelGL
}

type ModelGL struct {
	MeshesGL []MeshGL
}

type MeshGL struct {
	VAO  uint32
	VBO  uint32
	EBO  uint32
	Size int32
}

type ShaderType int

const (
	Lambert = iota
)

type StdUniform string

const (
  ProjectionMatrix = "proj\x00"
  ViewMatrix = "view\x00"
  ModelMatrix = "model\x00"

  MatColor = "color\x00"
)

type Material struct {
	ShaderType ShaderType
  Color mgl32.Vec4
}

type Shader struct {
  Program uint32
  uniforms UniformMap
}

type ShaderMap map[ShaderType]Shader
type UniformMap map[StdUniform]int32
