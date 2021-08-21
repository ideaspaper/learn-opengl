# Learn OpenGL

## VAO

- VAO (Vertex Array Object): objects in which you can store data about 3D model
- VAO have slots in which you can store data
- These slots also known as attribute lists
- For example, you can store vertex positions, vertex colors, normal vectors, etc.

## VBO

- VBO (Vertex Buffer Object): is just an array of numbers
- These data can be anything
- Each VBO can be put into attribute list

## Steps

1. Create data
1. Create VAO
1. Bind the VAO
1. Create VBO
1. Bind the VBO
1. Store data into VBO
1. Store VBO into one of the attribute lists of the VAO
1. Unbind the VBO
1. Unbind the VAO
1. Get the id of the VAO
1. Enable the attribute list where the data stored, using EnableVertexAttribArray
1. Tell OpenGL to render the VAO, using DrawArrays