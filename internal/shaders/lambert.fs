#version 330 core

in vec2 TexCoord;
out vec4 outputColor;

uniform vec4 MatColor;
uniform sampler2D MatDiffTex;
uniform float MatDiffOpacity;

void main(){
  outputColor = mix(MatColor, texture(MatDiffTex, TexCoord), MatDiffOpacity);
}
