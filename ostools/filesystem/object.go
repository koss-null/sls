package filesystem

type FSObject interface {
	Path() string
	Name() string

	IsFolder() bool
	IsFile() bool

	WeightBit() int
	CountWeightBit() chan int
}
