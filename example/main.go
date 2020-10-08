package main

import (
	"fmt"

	"os"

	"net/url"

	wallpaper ".."
)

func exists(name string) (bool, error) {
	// https://stackoverflow.com/a/22467409/8608146
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func isURL(str string) bool {
	// https://stackoverflow.com/a/55551215/8608146
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func main() {
	args := os.Args[1:]

	background, err := wallpaper.Get()

	if err != nil {
		panic(err)
	}

	// more than one arg => set wallpaper
	if len(args) > 0 {
		fmt.Println(args[0])
		fmt.Print("Checking if it's a file ...")
		valid, err := exists(args[0])
		if err == nil && valid {
			erre := wallpaper.SetFromFile(args[0])
			if erre == nil {
				fmt.Println("\nSet wallpaper successfully")
			}
			return
		}
		fmt.Println(" Apparently not.")
		fmt.Print("Checking if it's a valid url ...")
		if isURL(args[0]) {
			erre := wallpaper.SetFromURL(args[0])
			if erre == nil {
				fmt.Println("\nSet wallpaper successfully")
			}
		}
		fmt.Println(" Apparently not.\nPlease check the args")
		return
	}
	// else print the wallpaper
	fmt.Print(background)
}
