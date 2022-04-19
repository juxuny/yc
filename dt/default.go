package dt

func NewID(id uint64) ID {
	return ID{
		Uint64: id,
		Valid:  true,
	}
}
