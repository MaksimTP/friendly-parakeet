package main

type Vertex struct {
	x, y, z float32
}

func (v *Vertex) Transform(params interface{}) {

}

type BaseFileReader interface {
	LoadFigure(path string) Figure
}

type ObjFileReader struct {
	Figure_ Figure
}

func (fr ObjFileReader) LoadFigure(path string) Figure {
	return Figure{}
}

type Figure struct {
	polygons []int
	vertices []Vertex
}

func (f *Figure) TransformFigure(params interface{}) {
	for _, v := range f.vertices {
		v.Transform([]int{1, 2, 3})
	}
}

type BaseDrawer interface {
	Draw(figure Figure)
}

type QTDrawer struct{}

func (qtd QTDrawer) Draw(figure Figure) {

}

type Facade struct {
	reader_ BaseFileReader
	drawer_ BaseDrawer
	Figure_ Figure
}

func (f *Facade) LoadFigure(path string) Figure {
	return f.reader_.LoadFigure(path)
}

func (f *Facade) Draw() {
	f.drawer_.Draw(f.Figure_)
}

func main() {
	facade := &Facade{reader_: ObjFileReader{}, drawer_: QTDrawer{}}
	facade.LoadFigure("file.obj")
	facade.Draw()
}
