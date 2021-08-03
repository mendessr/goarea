package goarea

import "math"

// Pi é uma proporção número definida pela relação entre
// o perímetro de uma cincunferência e seu diâmetro
const Pi = 3.1416

//Circ é responsavel por calcular a area da cincunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi

}

// Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visivel!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
dasdas
