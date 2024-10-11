package training

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type ShoppingBasket struct {
	products []Product
}

// Ajouter un produit au panier
func (b *ShoppingBasket) AddProductIntoMyShoppingBasket(p Product) {
	b.products = append(b.products, p)
}

// Calculer le prix total des produits dans le panier
func (b *ShoppingBasket) TotalProductsPrice() float64 {
	var totalPrice float64
	for _, product := range b.products {
		totalPrice += product.Price * float64(product.Quantity)
	}
	return totalPrice
}

// Afficher les produits du panier
func (b *ShoppingBasket) DisplayBasket() {
	if len(b.products) == 0 {
		fmt.Println("Your basket is empty!")
		return
	}
	fmt.Println("Your Shopping Basket contains:")
	for _, product := range b.products {
		fmt.Printf("%d x %s - Unit Price: %.2f\n", product.Quantity, product.Name, product.Price)
	}
}

func main() {
	// Créer un panier
	basket := ShoppingBasket{}

	// Ajouter des produits au panier
	basket.AddProductIntoMyShoppingBasket(Product{Name: "Apple", Price: 0.99, Quantity: 5})
	basket.AddProductIntoMyShoppingBasket(Product{Name: "Banana", Price: 1.20, Quantity: 3})

	// Afficher le contenu du panier
	basket.DisplayBasket()

	// Calculer et afficher le prix total des produits dans le panier
	totalPrice := basket.TotalProductsPrice()
	fmt.Printf("Total Price: %.2f\n", totalPrice)
}
