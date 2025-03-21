package triangle

import (
	"go-rpg/internal/common"
	"image"
	"image/draw"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const width, height = 800, 600

func getUniformLocationToProgram(program uint32, varName string) int32 {

	gl.UseProgram(program)

	modelUniform := gl.GetUniformLocation(
		program,
		gl.Str(varName+"\x00"),
	)

	return (modelUniform)
}

func Main() {
	window := common.NewWindow(width, height)
	common.InitGlow()
	defer glfw.Terminate()

	// 1. Bind the vao
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Push all the vertex info inside vao:
	// 1.a. Gen and load vertex data
	vertices := []float32{
		0.5, 0.5, 0, 0, 0, // top right
		-0.5, 0.5, 0, 1, 0, // top left
		0.5, -0.5, 0, 0, 1, // bottom right
		-0.5, -0.5, 0, 0, 0, // bottom left
	}
	indices := []int32{
		0, 1, 2,
		1, 2, 3,
	}

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(
		gl.ELEMENT_ARRAY_BUFFER,
		len(indices)*4,
		gl.Ptr(indices),
		gl.STATIC_DRAW,
	)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(vertices)*4,
		gl.Ptr(vertices),
		gl.STATIC_DRAW,
	)

	var stride int32 = 3*4 + 2*4

	// 1.b. Specify the vertex attribute data
	gl.VertexAttribPointerWithOffset(
		0,
		3,
		gl.FLOAT,
		false,
		stride,
		0,
	)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointerWithOffset(
		1,
		2,
		gl.FLOAT,
		false,
		stride,
		3*4,
	)
	gl.EnableVertexAttribArray(1)

	// 2. Load, compile and link program
	vsSource := common.LoadShaderSource("./internal/triangle/vertexShader.vs")
	fsSource1 := common.LoadShaderSource("./internal/triangle/fragmentShader1.fs")
	program1 := common.NewProgram(vsSource, fsSource1)

	// 3. model
	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(
		program1,
		gl.Str("model\x00"),
	)

	gl.UseProgram(program1)
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	// 4. color
	color := mgl32.Vec4([]float32{0, 0, 0, 0})
	colorUniform := gl.GetUniformLocation(program1, gl.Str("color\x00"))
	gl.Uniform4fv(colorUniform, 1, &color[0])

	angle := 0.0
	prevTime := glfw.GetTime()

	// 5. Load textute
	imgFile, err := os.Open("./internal/textures/red_brick_diff_1k.jpg")
	if err != nil {
		log.Panicln("Loading image failed", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Panicln("Decoding image failed", err)
	}
	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		log.Panicln("Stride check failed")
	}

	iwidth := int32(rgba.Rect.Size().X)
	iheight := int32(rgba.Rect.Size().Y)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR) // minification filter
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		iwidth,
		iheight,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix),
	)
	gl.GenerateMipmap(texture)

	texUniform := gl.GetUniformLocation(program1, gl.Str("tex\x00"))
	gl.Uniform1i(texUniform, int32(texture-gl.TEXTURE0))

	for !window.ShouldClose() {

		gl.ClearColor(0, 0, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		currTime := glfw.GetTime()
		elapsed := currTime - prevTime
		prevTime = currTime

		angle += elapsed / 2

		model := mgl32.HomogRotate3D(
			float32(angle),
			mgl32.Vec3{0, 1, 0},
		)

		gl.UseProgram(program1)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture)

		gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
		gl.BindVertexArray(vao)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

		gl.DrawElements(
			gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil,
		)

		window.SwapBuffers()
		glfw.PollEvents()
	}

}
