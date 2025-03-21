package renderer

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func LoadTextureFromFile(
	filePath string,
	wrap_s int32,
	wrap_r int32,
	opacity float32,
) TextureGL {
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Panicln("Load texture failed", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Panicln("Decode texture failed", err)
	}

	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(
		rgba,
		rgba.Bounds(),
		img,
		image.Pt(0, 0),
		draw.Src,
	)

	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	target := uint32(gl.TEXTURE_2D)

	var handle uint32
	gl.GenTextures(1, &handle)

	texture := TextureGL{opacity, target, handle}
	texture.Bind(gl.TEXTURE0)

	gl.TextureParameteri(texture.Target, gl.TEXTURE_WRAP_S, wrap_s)
	gl.TextureParameteri(texture.Target, gl.TEXTURE_WRAP_R, wrap_r)

	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		width,
		height,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix),
	)
	gl.GenerateMipmap(texture.Handle)

	return texture
}

func (self *TextureGL) Bind(bindTo uint32) {
	gl.ActiveTexture(bindTo)
	gl.BindTexture(self.Target, self.Handle)
}

func (self *Texture) Load() TextureGL {
	return LoadTextureFromFile(
		self.Filepath,
		self.Wrap_s,
		self.Wrap_r,
		self.Opacity,
	)
}
