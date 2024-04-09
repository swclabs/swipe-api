// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"fmt"
	"log"

	"swclabs/swipe-api/delivery"
	"swclabs/swipe-api/delivery/adapter"
	"swclabs/swipe-api/internal/config"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := delivery.NewServer(addr)
	adapter := adapter.New(adapter.TypeBase)

	if err := server.Connect(adapter); err != nil {
		log.Fatal(err)
	}
}
