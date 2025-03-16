#version 330 core

in vec2 TexCoord;
out vec4 outputColor;

uniform vec4 MatColor;
uniform sampler2D MatTex;

void main(){
  outputColor = texture(MatTex, TexCoord);
}
