var input = require('fs').readFileSync('./stdin', 'utf8')
var lines = input.split('\n')

var valorA = lines.shift()
var valorB = lines.shift()
var valorC = lines.shift()

console.log(valorA)
console.log(valorB)
console.log(valorC)

if (valorA == "vertebrado") {
  if (valorB == "ave") {
    if (valorC == "carnivoro") {
      console.log("aguia")
    } else if (valorC == "onivoro") {
      console.log("pomba")
    }
  } else if (valorB == "mamifero") {
    if (valorC == "onivoro") {
      console.log("homem")
    } else if (valorC == "herbivoro") {
      console.log("vaca")
    }
  }
} else if (valorA == "invertebrado") {
  if (valorB == "inseto") {
    if (valorC == "hematofago") {
      console.log("pulga")
    } else if (valorC == "herbivoro") {
      console.log("lagarta")
    }
  } else if (valorB == "anelideo") {
    if (valorC == "hematofago") {
      console.log("sanguessuga")
    } else if (valorC == "onivoro") {
      console.log("minhoca")
    }
  }
}