package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
	"net/http"
	"github.com/codegangsta/negroni"
)

func playSE(se string) {
	_, err := exec.Command("afplay", "sound/" + se).Output()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	list, err := ioutil.ReadDir("sound")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	mux := http.NewServeMux()
	for _, finfo := range list {
		soundName := strings.Split(finfo.Name(), ".")[0]
		mux.HandleFunc("/" + soundName, func(w http.ResponseWriter, req *http.Request) {
			playSE(soundName + ".mp3")
		})
	}

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}



