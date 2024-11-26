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
	Día                int // Día de la fecha
	Mes                time.Month // Mes de la fecha en enum time.Month
	Año                int // Año de la fecha
	FechaEnFormatoTime time.Time // Fecha en formato time.Time
}

type Mes struct {
	Nombre       string // Nombre del mes
	Año          int // Año al que pertenecerá el mes
	DíaSeñalado  int // Día elegido por el usuario
	SemanaInicio int // Semana en la que empieza el mes
	SemanaFin    int // Semana en la que termina el mes
	Semana       map[int][7]int // Semanas del mes organizadas en un mapa de 7 días, indexado por el número de la semana
	TotalSemanas int // Total de semanas que tiene el mes
}

/** Compilamos la expresión regular para tenerla disponible en cualquier punto
del código sin tener que recompilar cada vez
**/
var expresiónRegular *regexp.Regexp = regexp.MustCompile(`^\s*(?:(\d{1,2})\s+)?(?:(\d{1,2})\s+)?(\d{4})$`)

// Constantes para el formato del texto con colores
const reset = "\033[0m"
const textoCian = "\033[36m"
const textoRojo = "\033[31m"
const textoBlancoSobreFondoNegro = "\033[37;40m"

/** Función para extraer los argumentos de la línea de comandos que van
precedidos de '-'
**/
func procesarFlags() (bool, bool, int, error) {

	númeroDeMeses := 0

	soloUno := flag.Bool("1", false, "Muestra solo un mes")
	triada := flag.Bool("3", false, "Muestra un mes junto con el anterior y el posterior")
	mostrarNúmerosDeSemana := flag.Bool("week-numbering", false, "Muestra los números de la semana")
	totalMeses := flag.Int("months", 0, "Muestra el número de meses indicado. Máximo 12")

	flag.Parse()

	if *soloUno {
		númeroDeMeses = 1
	} else {
		if *triada {
			númeroDeMeses = 3
			*triada = true
		} else {
			if *totalMeses > 0 {
				if *totalMeses > 12 {
					*totalMeses = 12
				} else {
					númeroDeMeses = *totalMeses
				}
			}
		}
	}

	return *triada, *mostrarNúmerosDeSemana, númeroDeMeses, nil
}

/** Función para extraer los argumentos de la línea de comandos que no van
precedidos de '-'. Se corresponden con la fecha introducida por el usuario en
alguno de estos formatos: día mes año mes año, año
**/
func procesarFecha(númeroDeMeses int) (fechaSolicitada Fecha, numMeses int, err error) {

	numMeses = númeroDeMeses
	err = nil
	fechaDeHoy := time.Now()
	fechaSolicitada = Fecha{
		Día: fechaDeHoy.Day(),
		Mes: fechaDeHoy.Month(),
		Año: fechaDeHoy.Year(),
	}
	fechaSolicitada.FechaEnFormatoTime = time.Date(fechaSolicitada.Año, fechaSolicitada.Mes, fechaSolicitada.Día, 0, 0, 0, 0, time.UTC)

	args := flag.Args()
	numArgs := len(args)

	if numArgs > 0 {
		fecha := strings.Join(args, " ")
		fragmentos := expresiónRegular.FindStringSubmatch(fecha)

		if fragmentos != nil {
			switch numArgs {

			case 1:
				if fragmentos[3] != "" {
					fechaSolicitada.Año, _ = strconv.Atoi(fragmentos[3])
					if fechaSolicitada.Año < 0 {
						err = fmt.Errorf("el año debe ser mayor que 0")
					}
				}
				if numMeses == 0 {
					numMeses = 12
				}
			case 2:
				if fragmentos[1] != "" {
					númeroMes, _ := strconv.Atoi(fragmentos[1])
					fechaSolicitada.Mes = time.Month(númeroMes)
					if fechaSolicitada.Mes <= 0 || fechaSolicitada.Mes > 12 {
						err = fmt.Errorf("el mes debe estar entre 1 y 12")
					}
				}
				if fragmentos[3] != "" {
					fechaSolicitada.Año, _ = strconv.Atoi(fragmentos[3])
					if fechaSolicitada.Año < 0 {
						err = fmt.Errorf("el año debe ser mayor que 0")
					}
				}
			case 3:
				if fragmentos[1] != "" {
					fechaSolicitada.Día, _ = strconv.Atoi(fragmentos[1])
					if fechaSolicitada.Día <= 0 || fechaSolicitada.Día > 31 {
						err = fmt.Errorf("el día debe estar entre 1 y 31")
					}
				}
				if fragmentos[2] != "" {
					númeroMes, _ := strconv.Atoi(fragmentos[2])
					fechaSolicitada.Mes = time.Month(númeroMes)
					if fechaSolicitada.Mes <= 0 || fechaSolicitada.Mes > 12 {
						err = fmt.Errorf("el mes debe estar entre 1 y 12")
					}
				}
				if fragmentos[3] != "" {
					fechaSolicitada.Año, _ = strconv.Atoi(fragmentos[3])
					if fechaSolicitada.Año < 0 {
						err = fmt.Errorf("el año debe ser mayor que 0")
					}
				}
			}
		}

		fechaSolicitada.FechaEnFormatoTime = time.Date(fechaSolicitada.Año, fechaSolicitada.Mes, fechaSolicitada.Día, 0, 0, 0, 0, time.UTC)
	} else {
		fmt.Println("No has introducido ninguna fecha. Por defecto cogemos la fecha de hoy: ", fechaSolicitada.Día, fechaSolicitada.Mes, fechaSolicitada.Año)
	}
	return
}


// Función para extraer todos los argumentos de la línea de comandos
func ExtraerArgumentos() (bool, bool, int, Fecha, error) {

	triada, mostrarNúmerosDeSemana, númeroDeMeses, _ := procesarFlags()
	fechaSolicitada, númeroDeMeses, err := procesarFecha(númeroDeMeses)

	fmt.Printf("Quieres que te muestre el calendario con %d meses", númeroDeMeses)
	if triada {
		fmt.Printf(" junto con el anterior y el posterior")
	}
	if mostrarNúmerosDeSemana {
		fmt.Printf(" y con los números de la semana")
	}
	fmt.Printf(" para la fecha %d/%s/%d\n\n", fechaSolicitada.Día, fechaSolicitada.Mes, fechaSolicitada.Año)

	return triada, mostrarNúmerosDeSemana, númeroDeMeses, fechaSolicitada, err
}

// Función para calcular el número de días que tiene ese mes
func CalcularDíasDelMes(año int, mes time.Month) int {
	//Calcular primer día del mes siguiente
	primerDíaMesSiguiente := time.Date(año, mes+1, 1, 0, 0, 0, 0, time.UTC)

	//Calcular el día inmediatamente anterior a primerDíaMesSiguiente
	últimoDíaMes := primerDíaMesSiguiente.AddDate(0, 0, -1)

	return últimoDíaMes.Day()
}

/** Función para construir e inicializar una estructura de tipo Mes en base a una fecha
**/
func InicializarCalendarioDelMes(fechaSolicitada Fecha) (mes Mes) {

	numDías := CalcularDíasDelMes(fechaSolicitada.Año, fechaSolicitada.Mes)

	fechaInicial := time.Date(fechaSolicitada.Año, fechaSolicitada.Mes, 1, 0, 0, 0, 0, time.UTC)
	_, numSemanaInicial := fechaInicial.ISOWeek()

	fechaFinal := fechaInicial.AddDate(0, 0, numDías-1)
	_, numSemanaFinal := fechaFinal.ISOWeek()

	totalSemanas := numSemanaFinal - numSemanaInicial + 1
	if (totalSemanas) < 0 {
		totalSemanas = totalSemanas + 52
	}

	mes = Mes{Nombre: fechaSolicitada.Mes.String(), Año: fechaSolicitada.Año, DíaSeñalado: fechaSolicitada.Día, SemanaInicio: numSemanaInicial, SemanaFin: numSemanaFinal, Semana: make(map[int][7]int), TotalSemanas: totalSemanas}

	diaDeLaSemana := int(fechaInicial.Weekday())
	if diaDeLaSemana == 0 {
		diaDeLaSemana = 6
	} else {
		diaDeLaSemana = diaDeLaSemana - 1
	}

	for inc, d, p := 0, 1, diaDeLaSemana; inc < mes.TotalSemanas; inc++ {
		s := numSemanaInicial + inc
		if s >= 53 {
			s = s - 52
		}
		semana := [7]int{0, 0, 0, 0, 0, 0, 0}

		for pos := p; pos <= 6 && d <= numDías; pos, d = pos+1, d+1 {
			semana[pos] = d
		}

		mes.Semana[s] = semana
		p = 0
	}
	return mes
}

// Función para sacar por pantalla el nombre del mes y el año
func (mes *Mes) PintarNombreMes() (err error) {

	totalEspacios := 24 - len(mes.Nombre)
	EspaciosIzquierda := totalEspacios / 2
	EspaciosDerecha := totalEspacios/2 + totalEspacios%2

	cabecera := strings.Repeat(" ", EspaciosIzquierda) + mes.Nombre + " " + strconv.Itoa(mes.Año) + strings.Repeat(" ", EspaciosDerecha)
	fmt.Print(textoCian + cabecera + reset)

	return nil
}

// Función para sacar por pantalla el nombre de los días de la semana
func (mes *Mes) PintarNombreDías() (err error) {
	días := [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	fmt.Print("    ")
	for i, d := range días {
		fmt.Print(textoCian + d + reset)
		if i < 6 {
			fmt.Print(" ")
		}
	}
	return nil
}

// Función para sacar por pantalla la semana solicitada
func (mes *Mes) PintarSemana(s int) (err error) {
	for _, d := range mes.Semana[s] {
		if d == 0 {
			fmt.Print("  ")
		} else {
			espacioIzquierda := 2 - len(strconv.Itoa(d))
			cadenaDía := strings.Repeat(" ", espacioIzquierda) + strconv.Itoa(d)
			if d == mes.DíaSeñalado {
				// Imprimir el texto con el estilo deseado
				fmt.Print(textoBlancoSobreFondoNegro + cadenaDía + reset)
			} else {
				fmt.Print(cadenaDía)
			}
		}
		fmt.Print(" ")
	}
	return nil
}

/** Función para sacar por pantalla la información de una semana concreta de ese mes: número de semana si fue solicitado y los días de esa semana
**/
func (mes *Mes) PintarDías(incremento int, mostrarNúmerosDeSemana bool) (err error) {

	if incremento < (*mes).TotalSemanas {
		s := mes.SemanaInicio + incremento
		if s > 52 {
			s = s - 52
		}
		if mostrarNúmerosDeSemana {
			espacioIzquierda := 2 - len(strconv.Itoa(s))
			fmt.Printf(strings.Repeat(" ", espacioIzquierda)+textoRojo+"%d: "+reset, s)
		} else {
			fmt.Print("    ")
		}
		mes.PintarSemana(s)
	} else {
		fmt.Print("                         ")

	}

	return nil
}

/** Función para pintar el calendario dependiendo de los argumentos con los que
se llamó al programa: número de meses, mostrar números de semana, fecha
específica
**/

func PintarCalendario(triada bool, mostrarNúmerosDeSemana bool, númeroDeMeses int, fechaSolicitada Fecha) (err error) {

	/** Creamos un slice de arrays de 3 elementos para pintar como mucho 3
	meses por línea. El slice inicialmente está vacío y tendrá capacidad para 4
	arrays de 3 elementos (12 meses como máximo)
	**/
	
	mesesAPintar := make([][3]Mes, 0, 4)

	if triada {
		mesAnterior := (fechaSolicitada.FechaEnFormatoTime.AddDate(0, -1, 0))
		mesSiguiente := (fechaSolicitada.FechaEnFormatoTime.AddDate(0, 1, 0))

		mesesAPintar = append(mesesAPintar, [3]Mes{
			InicializarCalendarioDelMes(Fecha{Día: 0, Mes: mesAnterior.Month(), Año: mesAnterior.Year(), FechaEnFormatoTime: fechaSolicitada.FechaEnFormatoTime.AddDate(0, -1, 0)}),
			InicializarCalendarioDelMes(fechaSolicitada),
			InicializarCalendarioDelMes(Fecha{Día: 0, Mes: mesSiguiente.Month(), Año: mesSiguiente.Year(), FechaEnFormatoTime: fechaSolicitada.FechaEnFormatoTime.AddDate(0, 1, 0)})})
	} else {
		mesActual := fechaSolicitada.FechaEnFormatoTime
		for i := 0; i < númeroDeMeses; i += 3 {
			var meses [3]Mes

			for j := 0; j < 3 && i+j < númeroDeMeses; j++ {

				if (i + j) == 0 {
					meses[j] = InicializarCalendarioDelMes(fechaSolicitada)
				} else {
					meses[j] = InicializarCalendarioDelMes(Fecha{Día: 0, Mes: mesActual.Month(), Año: mesActual.Year(), FechaEnFormatoTime: mesActual})
				}
				mesActual = mesActual.AddDate(0, 1, 0)
			}
			mesesAPintar = append(mesesAPintar, meses)
		}
	}
	for _, array := range mesesAPintar {

		for _, m := range array {
			if m.Nombre != "" {
				m.PintarNombreMes()
			}
		}
		fmt.Println()

		for _, m := range array {
			if m.Nombre != "" {
				m.PintarNombreDías()
				fmt.Print("     ")
			}
		}
		fmt.Println()

		máximoNúmeroDeSemanas := 0
		for _, m := range array {
			if máximoNúmeroDeSemanas < m.TotalSemanas {
				máximoNúmeroDeSemanas = m.TotalSemanas
			}
		}
		for inc := 0; inc < máximoNúmeroDeSemanas; inc++ {
			for _, m := range array {
				if m.Nombre != "" {
					m.PintarDías(inc, mostrarNúmerosDeSemana)
					fmt.Print("    ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	return nil
}

func main() {

	triada, mostrarNúmerosDeSemana, númeroDeMeses, fechaSolicitada, err := ExtraerArgumentos()
	if err != nil {
		fmt.Printf("Error procesando los argumentos: %s", err)
		return
	}
	err = PintarCalendario(triada, mostrarNúmerosDeSemana, númeroDeMeses, fechaSolicitada)

	if err != nil {
		fmt.Printf("Error pintando el calendario: %s", err)
	}

}
