package renderer

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

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

type Renderer struct {
  SceneGL SceneGL
  State RendererState
  Window *glfw.Window
  ShaderMap ShaderMap
}

type RendererState struct {
  TimeDelta float64
  Time float64
}

type SceneGL struct {
	ModelsGL []ModelGL
  CameraGL CameraGL
}

type CameraGL struct {
  ProjectionMatrix mgl32.Mat4
  ViewMatrix mgl32.Mat4
}

type ModelGL struct {
	MeshesGL []MeshGL
  ModelMatrix mgl32.Mat4
}

type MeshGL struct {
	VAO  uint32
	VBO  uint32
	EBO  uint32
	Size int32
  Material Material
}

type ShaderType int

const (
	Lambert = iota
)

const (
  ProjectionMatrix = "ProjectionMatrix"
  ViewMatrix = "ViewMatrix"
  ModelMatrix = "ModelMatrix"

  MatColor = "MatColor"
)

type Material struct {
	ShaderType ShaderType
  Color mgl32.Vec4
}

type Shader struct {
  Program uint32
}

type ShaderMap map[ShaderType]Shader

