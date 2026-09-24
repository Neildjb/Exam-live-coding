package main

import "fmt"

type Basket struct{
	name	string
	color	string
	content	[]Clothing
}

type Clothing struct{
	name	string
	color	string
	state	string
}

func main(){
	PanierBlanc:= Basket{
		name: "panier blanc",
		color: "blanc",
		content: []Clothing{
			content(VetementsBlanc)
		},
	}
	PanierNoir:= Basket{
		name: "panier noir",
		color: "noir",
		content: []Clothing{
			content(VetementsNoir)
		},
	}
	PanierCouleur:= Basket{
		name: "panier couleur",
		color: "couleur",
		content: []Clothing{
			content(VetementsCouleur)
		},
	}
	VetementsBlanc:= Clothing{
		name: "pantalon blanc",
		color: "blanc",
		state:	"propre",
	}
	VetementsNoir:= Clothing{
		name: "pantalon noir",
		color: "noir",
		state:	"propre",
	}
	VetementsCouleur:= Clothing{
		name: "pantalon couleur",
		color: "couleur",
		state:	"propre",
	}
}

func (d Basket) displayBasket() {
	fmt.println("=========== Panier blanc ===========")
	fmt.println(d.panierBlanc)
}



