package main

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "Adidas Shoe",
			Size: 12,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			Logo: "Adidas Shirt",
			Size: 32,
		},
	}
}
