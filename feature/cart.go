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

func CartProcess(menu menu.User, qty int) []Cart{
	if qty == 0 {
		fmt.Println("Input Jumlah Minimal 1")
		return nil
	}

	carts := []Cart{}

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