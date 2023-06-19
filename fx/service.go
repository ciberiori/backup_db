package fx

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/option"

	drive "google.golang.org/api/drive/v3"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// ServiceAccount : Use Service account
func ServiceAccount(credentialFile string) *http.Client {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		log.Fatal(err)
	}
	var c = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &c)
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(oauth2.NoContext)
	return client
}

func Pushfiles(source string) {
	conf := GetConfig()
	filename := conf.BACKUP_SRC + "/" + source
	drivefile := source                         // Filename
	baseMimeType := "text/plain"                // MimeType
	client := ServiceAccount("credential.json") // Please set the json file of Service account.

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fileInf, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	f := &drive.File{Name: drivefile, Parents: []string{"1W8san5dfzCMMg0d8EroB4xNQDPNN_8r6"}}
	res, err := srv.Files.
		Create(f).
		ResumableMedia(context.Background(), file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()
	if err != nil {
		log.Fatalln(err)
	}
	os.Remove(filename)
	fmt.Printf("%s\n", res.Id)
}

func ListFiles() {
	srv, err := drive.NewService(context.Background(), option.WithCredentialsFile("credential.json"))
	if err != nil {
		log.Fatal("Unable to access Drive API:", err)
	}
	r, err := srv.Files.List().PageSize(100).Fields("nextPageToken, files").Do()
	if err != nil {
		log.Fatal("Unable to list files:", err)
	}
	fmt.Println("Files:")
	for _, i := range r.Files {
		fmt.Printf("%v (%v) %v %v\n", i.Name, i.Id, i.MimeType, i.Parents)
	}
}
