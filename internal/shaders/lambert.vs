#version 330 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTex;

out vec2 TexCoord;

uniform mat4 ProjectionMatrix;
uniform mat4 ViewMatrix;
uniform mat4 ModelMatrix;

void main(){
  gl_Position = ProjectionMatrix * ViewMatrix * ModelMatrix * vec4(aPos, 1.0);
  TexCoord = aTex;
}
