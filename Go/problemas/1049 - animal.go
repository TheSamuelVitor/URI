package problemas

func Animal(a, b, c string) string {

	if a == "vertebrado" {
		if b == "ave" {
			if c == "carnivoro" {
				return "aguia"
			}
			if c == "onivoro" {
				return "pomba"
			}
		}

		if b == "mamifero" {
			if c == "onivoro" {
				return "homem"
			}
			if c == "herbivoro" {
				return "vaca"
			}
		}
	}

	if a == "invertebrado" {

		if b == "inseto" {
			if c == "hematofogo" {
				return "pulga"
			}
			if c == "herbivoro" {
				return "lagarta"
			}
		}

		if b == "anelideo" {
			if c == "hematofogo" {
				return "sanguessuga"
			}
			if c == "onivoro" {
				return "minhoca"
			}
		}
	}

	return ""

}
