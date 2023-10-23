package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestAnimal(t *testing.T) {
	lancheTests := []struct {
		name string
		got  string
		want string
	}{
		{name: "Teste Homem", got: problemas.Animal("vertebrado", "mamifero", "onivoro"), want: "homem"},
		{name: "Teste Aguia", got: problemas.Animal("vertebrado", "ave", "carnivoro"), want: "aguia"},
		{name: "Teste Pulga", got: problemas.Animal("invertebrado", "inseto", "hematofago"), want: "pulga"},
		{name: "Teste Sanguessuga", got: problemas.Animal("invertebrado", "anelideo", "hematofago"), want: "sanguessuga"},
	}

	for _, test := range lancheTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}
}
