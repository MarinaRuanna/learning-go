package main

import (
	"fmt"
	"runtime"
)

/*
EXERCÍCIO 1:

- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
1. 42
2. "James Bond"
3. true
- Agora demonstre os valores nestas variáveis, com:
1. Uma única declaração print
2. Múltiplas declarações print
*/

/*
EXERCÍCIO 2:

Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis.
Utilize os seguintes identificadores e tipos para estas variáveis:
    1. Identificador "x" deverá ter tipo int
    2. Identificador "y" deverá ter tipo string
    3. Identificador "z" deverá ter tipo bool
- Na função main:
    1. Demonstre os valores de cada identificador
    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
	Resposta: valor zero
*/

/*
EXERCÍCIO 3:

Utilizando a solução do exercício anterior:
    1. Em package-level scope, atribua os seguintes valores às variáveis:
        1. para "x" atribua 42
        2. para "y" atribua "James Bond"
        3. para "z" atribua true
    2. Na função main:
        1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
		Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
*/

/*
EXERCÍCIO 4:

Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
*/

/*
EXERCÍCIO 5:

Utilizando a solução do exercício anterior:
    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y".
	O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    2. Na função main:
        1. Isto já deve estar feito:
            1. Demonstre o valor da variável "x"
            2. Demonstre o tipo da variável "x"
            3. Atribua 42 à variável "x" utilizando o operador "="
            4. Demonstre o valor da variável "x"
        2. Agora faça também:
            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e,
				utilizando o operador "=", atribua o valor de "x" a "y"
            2. Demonstre o valor de "y"
            3. Demonstre o tipo de "y"
*/

func main() {
	Exemples()
	Declarations()
}

func Exemples() {
	var y int
	var z = true

	type tipo int

	var x tipo
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	s := fmt.Sprintf("%v\n%v\n", y, z)
	fmt.Println(s)
	fmt.Println(x)
	fmt.Printf("%T", x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("y: %T\n", y)
	fmt.Printf("y: %v\n", y)
	fmt.Println()
}

func Declarations() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v\n%v\n%v\n", x, y, z)
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)
}
