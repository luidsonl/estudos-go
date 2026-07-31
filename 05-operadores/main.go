package main

import "fmt"

func main() {
	fmt.Println("=== OPERADORES ARITMÉTICOS ===")
	a, b := 15, 4
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d (divisão inteira)\n", a/b)
	fmt.Printf("a %% b = %d (módulo/resto)\n", a%b)

	fmt.Println("\n=== CONCATENAÇÃO DE STRINGS ===")
	fmt.Println("Go" + "lang")

	// a++ e a-- são statements, não expressions
	fmt.Printf("\na (antes do ++) = %d\n", a)
	a++
	fmt.Printf("a (depois de ++) = %d\n", a)

	fmt.Println("\n=== OPERADORES DE COMPARAÇÃO ===")
	fmt.Printf("a (atual) = %d, b = %d\n", a, b)
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a >= b:", a >= b)

	fmt.Println("\n=== OPERADORES LÓGICOS (short-circuit) ===")
	x, y := true, false
	fmt.Printf("x = %t, y = %t\n", x, y)
	fmt.Println("x && y:", x && y)
	fmt.Println("x || y:", x || y)
	fmt.Println("!x:", !x)

	fmt.Println("\n=== OPERADORES BITWISE ===")
	m, n := 0b1100, 0b1010 // 12 e 10
	fmt.Printf("m = %04b (%d)\n", m, m)
	fmt.Printf("n = %04b (%d)\n", n, n)
	fmt.Printf("m & n  (AND)  = %04b (%d)\n", m&n, m&n)
	fmt.Printf("m | n  (OR)   = %04b (%d)\n", m|n, m|n)
	fmt.Printf("m ^ n  (XOR)  = %04b (%d)\n", m^n, m^n)
	fmt.Printf("m &^ n (AND NOT) = %04b (%d)\n", m&^n, m&^n)
	fmt.Printf("m << 2         = %04b (%d)\n", m<<2, m<<2)
	fmt.Printf("m >> 2         = %04b (%d)\n", m>>2, m>>2)
	fmt.Printf("^m     (NOT)   = %08b (%d)\n", ^m, ^m)

	fmt.Println("\n=== XOR: PROPRIEDADES ÚTEIS ===")
	// Toggle com XOR
	toggle := 0
	for i := 0; i < 5; i++ {
		fmt.Printf("  antes: toggle = %d", toggle)
		toggle ^= 1
		fmt.Printf("  -> depois: %d\n", toggle)
	}

	// Troca sem variável temporária (apenas demonstração, em Go use a, b = b, a)
	p, q := 7, 3
	fmt.Printf("\nAntes: p = %d, q = %d\n", p, q)
	p ^= q
	q ^= p
	p ^= q
	fmt.Printf("Depois (XOR swap): p = %d, q = %d\n", p, q)

	fmt.Println("\n=== ATRIBUIÇÃO COMPOSTA ===")
	v := 10
	fmt.Printf("v = %d\n", v)
	fmt.Printf("v += 5  (v = %d + 5) => ", v)
	v += 5
	fmt.Printf("%d\n", v)
	fmt.Printf("v -= 3  (v = %d - 3) => ", v)
	v -= 3
	fmt.Printf("%d\n", v)
	fmt.Printf("v *= 2  (v = %d * 2) => ", v)
	v *= 2
	fmt.Printf("%d\n", v)
	fmt.Printf("v /= 4  (v = %d / 4) => ", v)
	v /= 4
	fmt.Printf("%d\n", v)
	fmt.Printf("v %%= 3  (v = %d %% 3) => ", v)
	v %= 3
	fmt.Printf("%d\n", v)
}
