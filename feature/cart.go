package feature

import (
	"fmt"
	"project-golang/menu"
)

type Cart struct {
	IdCart int
	Menu menu.User
	Qty int
}

func (c Cart) Subtotal() int{
	return c.Qty * c.Menu.Price

} 

func CartProcess(menu menu.User, qty int, carts []Cart) []Cart{
	if qty == 0 {
		fmt.Println("Input Jumlah Minimal 1")
		return nil
	}

	for _, i := range carts{
		if i.Menu.IdMenu == menu.IdMenu{
			i.Qty += qty
			return carts
		}
	}

	idCart := 1
	if len(carts) > 0 {
		idCart = carts[len(carts)-1].IdCart + 1
	}

	carts = append(carts, Cart{
		IdCart: idCart,
		Menu: menu,
		Qty: qty,
	})

	return carts
}

func DisplayCart(Carts []Cart) int{
	totalHarga := 0

	fmt.Println(`--------------------------------------------------------------`)
	fmt.Println("---------------------Keranjang--------------------------------")
	
	for i, dataCart := range Carts{
		fmt.Printf("%d. %s\n", i+1, dataCart.Menu.Name)
		fmt.Printf("%dx @ %s\n", dataCart.Qty, FormatRupiah(dataCart.Menu.Price))
		fmt.Printf("Subtotal: %s\n", FormatRupiah(dataCart.Subtotal()))
		totalHarga += dataCart.Subtotal()
	}

	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("TOTAL HARGA: %s\n", FormatRupiah(totalHarga))
	fmt.Println("-------------------------------------------------------------")
	return totalHarga
}