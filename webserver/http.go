package webserver

import (
	"fmt"
	"github.com/codeedu/golang-cobra-example/app"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	Port string
}
//Quando o Webservice for executado será diponibilisado
//ResponseWrite - Respnde para o usário
//Request - Requece a requisição
func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {
	//Realiza a conversão de Str para float
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
	if err != nil {
		//Se der erro o valor será zero
		a = 0
	}
	//Realiza a convesão de Str para Floar
	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	if err != nil {
		//Se der erro o valor será zero
		b = 0
	}
	//Executa o aplicativo implementado
	calc := app.NewCalc()
	calc.A = a
	calc.B = b
	w.Write([]byte(fmt.Sprintf("Resultado é %f", calc.Sum())))
}
//Registra os EndPoint
func (s Server) Serve() {
	http.HandleFunc("/", s.SumHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
