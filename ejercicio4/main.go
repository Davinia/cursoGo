package main

import (
	"fmt"
)

func init() {
}

type ConjuntoDeEnteros map[int]bool

/*TODO: Utilizar cuando veamos programación genérica
type Conjunto interface{
	Crear()
	Añadir(interface{}) (error)
	Borrar(interface{}) (error)
	Existe(interface{}) (bool, error)
}*/

func (c *ConjuntoDeEnteros) Crear (){

	fmt.Println("Vamos a crear el conjunto...")
	if *c == nil{
		*c = make(map[int]bool)
		fmt.Println("Conjunto creado")
	}else{
		fmt.Println("El conjunto ya estaba creado")
	}

}

func (c *ConjuntoDeEnteros) Añadir (elemento int) (err error){

	fmt.Println("Vamos a añadir el elemento: ", elemento)
	if *c != nil{
		if !(*c)[elemento]{
			(*c)[elemento] = true
		}else{
			fmt.Println("El elemento ya existe")
			err = fmt.Errorf("%s","El elemento ya existe")
			return
		}
	}else{
			fmt.Println("El mapa no está inicializado")
			err = fmt.Errorf("%s","El mapa no está inicializado")
			return
	}
	return nil
}

func (c *ConjuntoDeEnteros) Borrar (elemento int) (err error){

	fmt.Println("Vamos a borrar el elemento: ", elemento)
	if *c != nil{
		if (*c)[elemento]{
			delete(*c,elemento)
		}else{
			fmt.Println("El elemento no existe")
			err = fmt.Errorf("%s","El elemento no existe")
			return
		}
	}else{
			fmt.Println("El mapa no está inicializado")
			err = fmt.Errorf("%s","El mapa no está inicializado")
			return
	}
	return nil
}

func (c *ConjuntoDeEnteros) Existe (elemento int) (existe bool, err error){

	fmt.Println("Vamos a comprobar si existe el elemento: ", elemento)
	if *c != nil{
		if (*c)[elemento]{
			existe = true
			fmt.Println("El elemento existe")
		}else{
			existe = false
			fmt.Println("El elemento no existe")
		}
	}else{
		existe = false
		fmt.Println("El mapa no está inicializado")
		err = fmt.Errorf("%s","El mapa no está inicializado")
	}
	return existe, err
}

func (c *ConjuntoDeEnteros) Imprimir (){

	fmt.Println("Tu conjunto contiene estos elementos: ")
	if *c != nil{
		for k,v := range (*c){
			if v {
				fmt.Printf("%v ",k)
			}
		}
		fmt.Println()
	}else{
		fmt.Println("El mapa no está inicializado")
	}
}

func main() {

	var miConjunto ConjuntoDeEnteros
	miConjunto.Crear()
	miConjunto.Añadir(1)
	miConjunto.Añadir(2)
	miConjunto.Añadir(2)
	miConjunto.Existe(2)
	miConjunto.Existe(3)
	miConjunto.Añadir(3)
	miConjunto.Borrar(2)
	miConjunto.Existe(2)
	miConjunto.Existe(3)
	miConjunto.Imprimir()

}
