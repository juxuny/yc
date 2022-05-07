package dt

func NewID(id uint64) ID {
	return ID{
		Uint64: id,
		Valid:  true,
	}
}

func NewIDPointer(id uint64) *ID {
	return &ID{
		Uint64: id,
		Valid:  true,
	}
}

func InvalidID() ID {
	return ID{}
}

func InvalidIDPointer() *ID {
	return &ID{}
}
