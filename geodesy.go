package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("temp/main.html", "temp/footer.html", "temp/header.html")
	t.ExecuteTemplate(w, "main", nil)
}

func result(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("temp/result.html", "temp/footer.html", "temp/header.html")
	X1 := r.FormValue("X1")
	Y1 := r.FormValue("Y1")
	X2 := r.FormValue("X2")
	Y2 := r.FormValue("Y2")
	x1, _ := strconv.ParseFloat(X1, 64)
	x2, _ := strconv.ParseFloat(X2, 64)
	y1, _ := strconv.ParseFloat(Y1, 64)
	y2, _ := strconv.ParseFloat(Y2, 64)
	n, u, m:= atanNumber(x1, y1, x2, y2)
	fmt.Println(n, u, m)
	t.ExecuteTemplate(w, "result", nil)
	fmt.Fprint(w, "Результат вычесления: ", n,"° ",u,"′ ", m,"″ ")
}

func atanNumber(x1, y1, x2, y2 float64) (int, int, int) {
	x := x2-x1
	y := y2-y1
	num := y/x
	res := math.Atan(num)
	res *=  (180/math.Pi)
	deg := int(res)
	min1:= (res-float64(deg))*60
	min := int(min1)
	sec1:= int((min1 - float64(min))*60)
	sec := int(sec1)
	if x<0 && y>0 {
		deg = 179 + deg
		min = 59 + min
		sec = 60 + sec
	}else if x<0 && y<0{
		deg = 180 + deg
	}else if x > 0 && y < 0 {
		deg = 359 + deg
		min = 59 + min
		sec = 60 + sec
	}
	return deg, min, sec
}

func handlefunc()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/result", result)
	http.ListenAndServe(":8080", nil)
}
func main() {
	handlefunc()
}
