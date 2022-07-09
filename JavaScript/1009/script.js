var input = require('fs').readFileSync('stdin', 'utf8')
var lines = input.split('\n')

var nome = lines.shift()
var salariofixo = parseFloat(lines.shift())
var vendas = parseFloat(lines.shift())

var salarioComissao = salariofixo + (vendas*15/100)

console.log(`TOTAL = R$ ${salarioComissao.toFixed(2)}`)