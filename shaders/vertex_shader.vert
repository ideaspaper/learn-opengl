# version 410

in vec3 position;

out vec3 color;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;

void main() {
    gl_Position = projectionMatrix * transformationMatrix * vec4(position, 1.0);
    color = vec3(position.x+0.5, 1.0, position.y+0.5);
}
