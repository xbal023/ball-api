package botwa

import (
	"context"
	"fmt"
	"os"
	"sync"

	v "github.com/TwiN/go-color"
	// Library Whatsmeow
	"go.mau.fi/whatsmeow"
	db "go.mau.fi/whatsmeow/store/sqlstore"
	ngelog "go.mau.fi/whatsmeow/util/log"
	// Sqlite
	_ "github.com/mattn/go-sqlite3"
	// Print to terminal code
	qrCode "github.com/mdp/qrterminal"
	// Module local
	handler "github.com/bolaxd/dumn/bot.wa/utils/handler"
)

func RunBot(wg *sync.WaitGroup) {
	defer wg.Done()
	dbName := os.Getenv("SESSION_NAME")
	typeDb := os.Getenv("DB_TYPE")
	dbLog := ngelog.Stdout("Database", "ERROR", true)
	container, err := db.New(typeDb, "file:bot.wa/"+dbName+"?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	ds, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	logs := ngelog.Stdout("Client", "ERROR", false)
	conn := whatsmeow.NewClient(ds, logs)
	msgUp := handler.MessageUpsert(conn)
	conn.SetForceActiveDeliveryReceipts(true)
	conn.AddEventHandler(msgUp)
	
	if conn.Store.ID == nil {
		qrChan, _ := conn.GetQRChannel(context.Background())
		err = conn.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Silahkan SCAN QR Code dibawah ini:\n")
				qrCode.GenerateHalfBlock(v.InPurple(evt.Code), qrCode.M, os.Stdout);
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = conn.Connect()
		fmt.Println(v.InCyan("Koneksi telah tersambung whatsapp web..."))
		fmt.Println(v.InPurple("------------------------------------------------"))
		if err != nil {
			panic(err)
		}
	}

	conn.Disconnect()
}
