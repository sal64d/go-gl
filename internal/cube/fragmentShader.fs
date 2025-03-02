#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;
in vec3 normal;

out vec4 outputColor;

void main() {
  outputColor = texture(tex, fragTexCoord);
  outputColor[0] = normal[0] + 0.2;
  outputColor[1] = normal[1] + 0.2;
  outputColor[2] = normal[2] + 0.2;
}

