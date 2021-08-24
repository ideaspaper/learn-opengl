package helpers

import "github.com/go-gl/mathgl/mgl32"

func CreateTransformationMatrix(position, angles, scales []float32) mgl32.Mat4 {
	result := mgl32.Ident4()
	result = result.Mul4(mgl32.Translate3D(position[0], position[1], position[2]))
	result = result.Mul4(mgl32.HomogRotate3DX(angles[0]))
	result = result.Mul4(mgl32.HomogRotate3DX(angles[1]))
	result = result.Mul4(mgl32.HomogRotate3DX(angles[2]))
	result = result.Mul4(mgl32.Scale3D(scales[0], scales[1], scales[2]))
	return result
}
