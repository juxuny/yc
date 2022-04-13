package dt

func NewID(id int64) ID {
	return ID{
		Int64: id,
		Valid: true,
	}
}
