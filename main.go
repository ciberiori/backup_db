package main

import (
	"dumpfx/fx"
	"fmt"
	"time"
)

func main() {

	nameFile := "backup-" + time.Now().UTC().Format("02-01-2006") + ".sql"

	fmt.Println("Comenzando backup de la db")
	fx.BackupDB(nameFile)
	fmt.Println("Subiendo archivos a Google Drive")
	fx.Pushfiles(nameFile)
	fmt.Println("Listando archivos en Google drive")
	fx.ListFiles()
}
