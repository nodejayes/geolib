package geometry

type ITransformable interface {
	Transform(target int) error
}

type IGeometry interface {
	ITransformable
	GetCoordinates(data interface{})
}
