var entrada = require('fs').readFileSync('stdin', 'utf8')
var lines = entrada.split('\n')

var raio = parseFloat(lines.shift())
var diametro = 3.14159*(raio**2)
console.log(`A=${diametro.toFixed(4)}`)