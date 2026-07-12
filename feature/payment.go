package feature

import (
	"fmt"
	"project-golang/utils"
	"time"
)

func Payment(payment int, totalPrice int, orders []Cart) []Cart {
	utils.CallClear()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowWIB := time.Now().In(loc)

	charge := payment - totalPrice
	fmt.Println(`-----------------------STRUK PEMBAYARAN----------------------`)
	fmt.Println(`                       Popeye Chicken`)
	fmt.Println(`-------------------------------------------------------------`)
	//inisiasi looping

	for _, order := range orders {
		fmt.Printf("# %s\n", order.Menu.Name)
		fmt.Printf("%dx @ %s = %s\n", order.Qty, FormatRupiah(order.Menu.Price), FormatRupiah(order.Subtotal()))
	}
	fmt.Println(`-------------------------------------------------------------`)
	fmt.Printf("Total    : %s\n", FormatRupiah(totalPrice))
	fmt.Printf("Bayar    : %s\n", FormatRupiah(payment))
	fmt.Printf("Kembali  : %s\n", FormatRupiah(charge))
	fmt.Println("Waktu Transaksi: ", nowWIB.Format("Monday, 02 January 2006 15:04:05"))
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("Terima kasih atas pesanan Anda!")
	fmt.Println("-------------------------------------------------------------")
	orders = []Cart{}

	return orders
}
