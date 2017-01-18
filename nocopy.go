package nocopy

// S should be embedded into a struct which mustn't be copied.
// This doesn't prevent copying, but will cause `go vet` to
// show a warning.
// Usual implementations would be:
//
// 	type DoNotCopyThis struct {
//		noCopy nocopy.S
// 	}
type S struct{}

// Lock sets the lock method to something `go vet` will
// understand. For details, see
// https://github.com/golang/go/issues/8005#issuecomment-190753527.
func (*S) Lock() {}
