package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {

}

type Fecha struct {
	Día int
	Mes time.Month
	Año int
}

type Mes struct {
	Nombre       string
	Año          int
	DíaSeñalado  int
	SemanaInicio int
	SemanaFin    int
	Semana       map[int][7]int
}

// Compilamos la expresión regular para tenerla disponible en cualquier punto del código sin tener que recompilar cada vez
var expresiónRegular *regexp.Regexp = regexp.MustCompile(`^\s*(?:(?P<dia>\d{1,2})\s+(?P<mes>\d{1,2})\s+)?(?P<anyo>\d{4})\s*$`)

func procesarFlags() (bool, bool, int, error) {

	númeroDeMeses := 1

	soloUno := flag.Bool("1", false, "Muestra solo un mes")
	triada := flag.Bool("3", false, "Muestra un mes junto con el anterior y el posterior")
	mostrarNúmerosDeSemana := flag.Bool("week-numbering", false, "Muestra los números de la semana")
	totalMeses := flag.Int("months", 0, "Muestra el número de meses indicado")

	flag.Parse()

	if *soloUno {
		númeroDeMeses = 1
	} else {
		if *triada {
			númeroDeMeses = 3
			*triada = true
		} else {
			if *totalMeses > 0 {
				númeroDeMeses = *totalMeses
			}
		}
	}

	return *triada, *mostrarNúmerosDeSemana, númeroDeMeses, nil
}

func procesarFecha() (Fecha, error) {

	fechaDeHoy := time.Now()
	fechaSolicitada := Fecha{
		Día: fechaDeHoy.Day(),
		Mes: fechaDeHoy.Month(),
		Año: fechaDeHoy.Year(),
	}

	args := flag.Args()

	if len(args) > 0 {

		fmt.Println("Has introducido los argumentos: ", args)
		fecha := strings.Join(args, " ")
		fmt.Println("Has introducido la fecha: ", fecha)

		fragmento := expresiónRegular.FindStringSubmatch(fecha)

		if fragmento != nil {
			if fragmento[1] != "" {
				fechaSolicitada.Día, _ = strconv.Atoi(fragmento[1])
			}
			if fragmento[2] != "" {
				númeroMes, _ := strconv.Atoi(fragmento[2])
				fechaSolicitada.Mes = time.Month(númeroMes)
			}
			if fragmento[3] != "" {
				fechaSolicitada.Año, _ = strconv.Atoi(fragmento[3])
			}
		}
	} else {
		fmt.Println("No has introducido ninguna fecha. Por defecto cogemos la fecha de hoy: ", fechaSolicitada.Día, fechaSolicitada.Mes, fechaSolicitada.Año)
	}
	return fechaSolicitada, nil
}

func ExtraerArgumentos() (bool, bool, int, Fecha) {

	triada, mostrarNúmerosDeSemana, númeroDeMeses, _ := procesarFlags()
	fechaSolicitada, _ := procesarFecha()

	fmt.Printf("Quieres que te muestre el calendario con %d meses", númeroDeMeses)
	if triada {
		fmt.Printf(" junto con el anterior y el posterior")
	}
	if mostrarNúmerosDeSemana {
		fmt.Printf(" y con los números de la semana")
	}
	fmt.Printf(" para la fecha dia %d/%s/%d\n", fechaSolicitada.Día, fechaSolicitada.Mes, fechaSolicitada.Año)

	return triada, mostrarNúmerosDeSemana, númeroDeMeses, fechaSolicitada
}

func CalcularDíasDelMes(año int, mes time.Month) int {
	//Calcular primer día del mes siguiente
	primerDíaMesSiguiente := time.Date(año, mes+1, 1, 0, 0, 0, 0, time.UTC)

	//Calcular el día inmediatamente anterior a primerDíaMesSiguiente
	últimoDíaMes := primerDíaMesSiguiente.AddDate(0, 0, -1)

	return últimoDíaMes.Day()
}

func InicializarCalendarioDelMes(fechaSolicitada Fecha) (mes Mes) {

	numDías := CalcularDíasDelMes(fechaSolicitada.Año, fechaSolicitada.Mes)

	fmt.Println("El mes tiene", numDías, "días")

	fechaInicial := time.Date(fechaSolicitada.Año, fechaSolicitada.Mes, 1, 0, 0, 0, 0, time.UTC)
	_, numSemanaInicial := fechaInicial.ISOWeek()

	fmt.Println("El mes comienza en la semana", numSemanaInicial)

	fechaFinal := fechaInicial.AddDate(0, 0, numDías-1)
	_, numSemanaFinal := fechaFinal.ISOWeek()

	fmt.Println("El mes termina en la semana", numSemanaFinal)

	mes = Mes{Nombre: fechaSolicitada.Mes.String(), Año: fechaSolicitada.Año, DíaSeñalado: fechaSolicitada.Día, SemanaInicio: numSemanaInicial, SemanaFin: numSemanaFinal, Semana: make(map[int][7]int)}

	for s, d, p := numSemanaInicial, 1, int(fechaInicial.Weekday()); s <= numSemanaFinal; s++ {

		semana := [7]int{0, 0, 0, 0, 0, 0, 0}

		for pos := p; pos <= 6 && d <= numDías; pos, d = pos+1, d+1 {
			semana[pos] = d
		}

		mes.Semana[s] = semana
		p = 0
	}
	return mes
}

func (mes *Mes) Pintar(mostrarNúmerosDeSemana bool) (err error) {

	totalEspacios := 21 - len(mes.Nombre)
	EspaciosIzquierda := totalEspacios / 2
	EspaciosDerecha := totalEspacios/2 + totalEspacios%2

	fmt.Println("Semana de inicio: " + strconv.Itoa(mes.SemanaInicio) + " Semana de fin: " + strconv.Itoa(mes.SemanaFin))

	cabecera := strings.Repeat(" ", EspaciosIzquierda) + mes.Nombre + " " + strconv.Itoa(mes.Año) + strings.Repeat(" ", EspaciosDerecha)
	fmt.Println(cabecera)

	días := [7]string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	fmt.Printf("   ")
	for _, d := range días {
		fmt.Printf(" "+d)
	}
	fmt.Println()

	for s := mes.SemanaInicio; s <= mes.SemanaFin; s++ {
		if mostrarNúmerosDeSemana {
			espacioIzquierda := 2 - len(strconv.Itoa(s))
			fmt.Printf(strings.Repeat(" ", espacioIzquierda) + "%d: ", s)
		} else {
			fmt.Printf("    ")
		}
		for _, d := range mes.Semana[s] {
			if d == 0 {
				fmt.Printf("  ")
			}else{
				espacioIzquierda := 2 - len(strconv.Itoa(d))
				fmt.Printf(strings.Repeat(" ", espacioIzquierda) + "%d",d)
			}
			fmt.Printf(" ")
		}
		fmt.Println()
	}

	return nil
}

func PintarCalendario(triada bool, mostrarNúmerosDeSemana bool, númeroDeMeses int, fechaSolicitada Fecha) (err error) {

	mes := InicializarCalendarioDelMes(fechaSolicitada)
	mes.Pintar(mostrarNúmerosDeSemana)
	return nil
}

func main() {

	triada, mostrarNúmerosDeSemana, númeroDeMeses, fechaSolicitada := ExtraerArgumentos()
	err := PintarCalendario(triada, mostrarNúmerosDeSemana, númeroDeMeses, fechaSolicitada)

	if err != nil {
		fmt.Printf("Error pintando el calendario: %s", err)
	}

}
