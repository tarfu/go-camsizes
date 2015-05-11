package main
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
	"html/template"
	"net/http"
)


type Brand struct {
	Name string
	Products []Product
}

type Product struct {
	Name string
	Sizes []Size
}

type Size struct {
	Name, Color string
	Start, End float32
}

func (s Size) GetWidth() int {
	return int(s.End) - int(s.Start)
}

type Brands struct {
	Brands []Brand
}

func (b Brands) StepSize() float32 {
	return 10
}

func (b Brands) MaxSize() float32 {
	var max float32
	for _, bb := range b.Brands {
		for _, p := range bb.Products {
			for _, s := range p.Sizes {
				if s.End > max {max = s.End}
			}
		}
	}
	return max
}

var VerticalPos float32


func main() {

	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request){
	file, e := ioutil.ReadFile("./static/cams.json")
	if e!= nil {
		fmt.Printf("Error reading file: %v\n", e)
		return
	}

	var jsontype Brands

	e = json.Unmarshal(file, &jsontype)
	if e != nil {
		fmt.Printf("Error unmarschal file: %v\n", e)
		return
	}

	funcMap := template.FuncMap{
		"add": func(a, b float32) float32 { return a+b },
		"verticalAdd": func(a float32) float32 {
			VerticalPos = VerticalPos + a
			return VerticalPos
		},
		"verticalReset": func() float32 {
			VerticalPos = float32(0)
			return VerticalPos
		},
		"verticalResetAdd": func(a float32) float32 {
			VerticalPos = float32(0)+a
			return VerticalPos
		},
		"sizeSteps": sizeSteps,
		"pasteJS": func() template.JS {
			js, e := ioutil.ReadFile("./static/cams.js")
			if e!= nil {
				fmt.Printf("Error reading JS file: %v\n", e)
				os.Exit(1)
			}
			return template.JS(js)
		},
	}

	t, e := template.New("cams.tmpl").Funcs(funcMap).ParseFiles("./template/cams.tmpl")
	if e != nil {
		fmt.Printf("Error parsing template file: %v\n", e)
		return
	}
	e =  t.ExecuteTemplate(w, "root", jsontype)
	if e != nil {
		fmt.Printf("Error executing template file: %v\n", e)
		return
	}
}

func sizeSteps(limit float32, stepsSize float32) []float32 {
	var steps []float32
	for i := float32(0.0); i<=limit+stepsSize; i = i+stepsSize {
		steps = append(steps, i)
	}
	return steps
}