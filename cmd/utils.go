package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func mimeType(arg string) string {
	// ? KDE Menu Editor puts edited files into ~/.local/share/applications
	// ? so files might not be in /usr/share/applications, should add some way to change it or parse that folder too?
	input := fmt.Sprintf("/usr/share/applications/%s.desktop", arg)
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "MimeType") {
			return scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Fatal("This application file doesn't contain a MimeType entry!")
	return ""
}

func processMime(mime string) []string {
	// The .desktop file of an application lists what file types it supports.
	// Luckily, this list is easily parsable.
	s := strings.Replace(mime, "MimeType=", "", -1)
	arr := strings.Split(s, ";")
	arr = arr[:len(arr)-1]
	return arr
}

func associate(mimes []string, app string, confirm bool, xdgbinlocation string) {
	if !confirm {
		fmt.Println("This will override your currently set default app for files that use a file type!")
		fmt.Println("Press Enter to confirm.")
		fmt.Scanln() // ? Is there a better way to do this?
	}
	for _, mime := range mimes {
		cmd := exec.Command(xdgbinlocation, "default", app+".desktop", mime)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

	}
}
