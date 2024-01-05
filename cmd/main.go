package main

import (
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Generate random data for bar chart
func generatebarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(100)})
	}
	return items
}

// Pie Categories
var Categories = []string{"A", "B", "C", "D", "E", "F"}

// Generate random data for pie chart
func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.PieData{Name: Categories[i], Value: rand.Intn(100)})
	}
	return items
}

// Geo Data Provinces
var geoData = []opts.GeoData{
	{Name: "北京", Value: []float64{116.40, 39.90, float64(rand.Intn(100))}},
	{Name: "上海", Value: []float64{121.47, 31.23, float64(rand.Intn(100))}},
	{Name: "重庆", Value: []float64{106.55, 29.56, float64(rand.Intn(100))}},
	{Name: "武汉", Value: []float64{114.31, 30.52, float64(rand.Intn(100))}},
	{Name: "台湾", Value: []float64{121.30, 25.03, float64(rand.Intn(100))}},
	{Name: "香港", Value: []float64{114.17, 22.28, float64(rand.Intn(100))}},
}

var guangdongData = []opts.GeoData{
	{Name: "汕头", Value: []float64{116.69, 23.39, float64(rand.Intn(100))}},
	{Name: "深圳", Value: []float64{114.07, 22.62, float64(rand.Intn(100))}},
	{Name: "广州", Value: []float64{113.23, 23.16, float64(rand.Intn(100))}},
}

func main() {

	home := homeHandler{}
	http.HandleFunc("/", home.ServeHTTP)
	http.HandleFunc("/chartpage", ChartMainPage)
	http.HandleFunc("/geopage", GeoMainPage)
	http.ListenAndServe(":8000", nil)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home Page\n"))
	w.Write([]byte("Charts ---> /chartpage\n"))
	w.Write([]byte("Geo ---> /geopage\n"))
}

func ChartMainPage(w http.ResponseWriter, r *http.Request) {
	ChartBar(w, r)
	ChartPie(w, r)
}

func GeoMainPage(w http.ResponseWriter, r* http.Request) {
	geoBase(w, r)
	geoGuangdong(w, r)
}

func ChartBar(w http.ResponseWriter, r *http.Request)  {
	// Create a new bar instance
	bar := charts.NewBar()
	// Set some global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Consumo Semestral",
		Subtitle: "Resumen Mensual",
	}))
	// Put Data into instance
	bar.SetXAxis([]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio"}).
	AddSeries("Consumo A", generatebarItems()).AddSeries("Consumo B", generatebarItems())
	bar.Render(w)
}

func ChartPie(w http.ResponseWriter, r *http.Request) {
	// Create a new pie instance
	pie := charts.NewPie()
	// Set some global options
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Distribución"}),
	)
	// Put data into instance
	pie.AddSeries("Pie", generatePieItems()).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{b}: {c}",}))
	pie.Render(w)
}

func geoBase(w http.ResponseWriter, r *http.Request) {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "China Map"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map:       "china",
			ItemStyle: &opts.ItemStyle{Color: "#cfd1d1"},
		}),
	)

	geo.AddSeries("geo", types.ChartEffectScatter, geoData,
		charts.WithRippleEffectOpts(opts.RippleEffect{
			Period:    4,
			Scale:     6,
			BrushType: "stroke",
		}),
	)
	geo.Render(w)
}

func geoGuangdong(w http.ResponseWriter, r *http.Request) {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Guangdong Province"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map: "广东",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)

	geo.AddSeries("geo", types.ChartScatter, guangdongData)
	geo.Render(w)
}