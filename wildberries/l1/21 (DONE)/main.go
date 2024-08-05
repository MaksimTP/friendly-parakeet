package main

import "fmt"

type JSON map[string]interface{}
type XML map[string]string

type NormalizationParameters struct {
	x_min, y_min, z_min float64
	x_max, y_max, z_max float64
}

func (l *NormalizationParameters) NormalizeData(data JSON) JSON {
	fmt.Println("Normalizing data...")
	return data
}

type Module struct {
	Vertices             []float64
	Polygons             []int
	CountVertices        int
	CountPolygons        int
	CountPolygonsIndices int
}

func (m *Module) GetData() JSON {
	fmt.Println("Getting necessary data...")
	return JSON{"Vertices": m.Vertices, "Polygons": m.Polygons}
}

func (m *Module) LoadData(path string) JSON {
	fmt.Println("Loading data...", path)
	return JSON{}
}

type Drawer interface {
	Draw(data JSON)
}

type QTDrawer struct {
	color   [3]int
	opacity int8
}

func (d *QTDrawer) SetColor() {
	fmt.Println("QT setting color")
}

func (d *QTDrawer) SetOpacity() {
	fmt.Println("QT setting opacity")
}

func (d *QTDrawer) Draw(data JSON) {
	d.SetColor()
	d.SetOpacity()
	fmt.Println("drawing")
}

type NewModule struct {
	Data XML
}

func (m *NewModule) GetData() XML {
	fmt.Println("Getting necessary data...")
	return m.Data
}

func (m *NewModule) LoadData(path string) XML {
	fmt.Println("Loading data...", path)
	return XML{}
}

type Adapter struct {
	Module NewModule
}

func (a *Adapter) GetData() JSON {
	a.Module.GetData()
	// Converting algorithm
	return JSON{}
}

func main() {

	// #1. OK

	// model := Module{}
	// data := model.GetData()
	// normalizer := NormalizationParameters{}
	// data = normalizer.NormalizeData(data)
	// var drawerModule Drawer = &QTDrawer{}
	// drawerModule.Draw(data)

	// #2. OK

	// modelAdap := Adapter{NewModule{}}
	// data := modelAdap.GetData()
	// normalizer := NormalizationParameters{}
	// data = normalizer.NormalizeData(data)
	// var drawerModule Drawer = &QTDrawer{}
	// drawerModule.Draw(data)

}
