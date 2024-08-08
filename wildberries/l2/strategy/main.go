package main

type Strategy interface {
	SearchProduct(filters []string)
}

type FirstAlgorithm struct{}

func (a FirstAlgorithm) SearchProduct(filters []string) {}

type SecondAlgorithm struct{}

func (a SecondAlgorithm) SearchProduct(filters []string) {}

type DataFinder struct {
	strategy Strategy
}

func (df DataFinder) GetData(filters []string) {
	df.strategy.SearchProduct(filters)
}

func main() {
	df := DataFinder{FirstAlgorithm{}}
	df.GetData([]string{"32", "Jeans", "L"})
}
