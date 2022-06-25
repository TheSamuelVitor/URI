var entrada = require('fs').readFileSync('stdin', 'utf8')
var linha = entrada.split('\n')

var raio = parseFloat(linha.shift())
var diametro = (3.1415)*(raio**2)
console.log(`A = ${diametro}`)