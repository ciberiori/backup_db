package fx

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func BackupDB(namefile string) {

	conf := GetConfig()
	os.Setenv("PGPASSWORD", conf.DB_PASSWORD)

	outfile, err := os.Create(conf.BACKUP_SRC + "/" + namefile)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("pg_dump",
		"-h",
		conf.DB_HOST,
		"-U",
		conf.DB_USERNAME,
		conf.DB_NAME)

	var stderr bytes.Buffer
	cmd.Stdout = outfile
	cmd.Stderr = &stderr
	err2 := cmd.Run()
	if err2 != nil {
		fmt.Println(fmt.Sprint(err2) + ": " + stderr.String())
		return
	}
	cmd.Wait()

	defer outfile.Close()

}
