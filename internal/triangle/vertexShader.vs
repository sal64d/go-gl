#version 330 core

uniform mat4 model;

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 texCoord;

out vec2 TexCoord;

void main()
{
  gl_Position = model * vec4(aPos, 1.0);
  TexCoord = texCoord;
}
