var input = require('fs').readFileSync('stdin', 'utf8')
var lines = input.split('\n')

var number = parseInt(lines.shift())
var qtdhoras = parseInt(lines.shift())
var horas = parseFloat(lines.shift())

var salario = qtdhoras*horas

console.log(`NUMBER = ${number}`)
console.log(`SALARY = U$ ${salario.toFixed(2)}`)