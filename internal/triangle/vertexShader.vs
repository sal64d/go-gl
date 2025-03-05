#version 330 core

uniform mat4 model;

layout (location = 0) in vec3 aPos;


void main()
{
  gl_Position = model * vec4(aPos.x, aPos.y, aPos.z, 1.0);
}
