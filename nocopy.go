package nocopy

type S struct{}

func (*S) Lock() {}
