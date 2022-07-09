var input = require('fs').readFileSync('stdin', 'utf8')
var lines = input.split('\n')

var distancia = parseInt(lines.shift())
var combustivel = parseFloat(lines.shift())

var consumomedio = distancia/combustivel

console.log(`${consumomedio.toFixed(3)} km/l`)