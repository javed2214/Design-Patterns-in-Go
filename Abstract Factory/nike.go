package main

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "Nike Shoe",
			Size: 10,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			Logo: "Nike Shirt",
			Size: 30,
		},
	}
}
