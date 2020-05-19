package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type FilesList struct {
	Name      []string
	PathValue string
}

func showFolder(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	path := r.FormValue("pathboxvalue")
	fmt.Println(path)

	reg1 := regexp.MustCompile("[/]$")
	res := reg1.Match([]byte(path))
	fmt.Println(res)

	if !res {
		if path != "" {
			path = path + "/"
		}
	}

	extention := "."
	fmt.Println("path:", path, "+", extention)

	file_names := make([]string, 0)
	var file_data FilesList

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	files, err := ioutil.ReadDir(path + extention)
	if err != nil {
		if path != "" {
			file_data.PathValue = path
		} else {
			file_data.PathValue = pwd
		}
		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}

		err = t.Execute(w, file_data)
		if err != nil {
			panic(err)
		}

		fmt.Println(err)

	} else {

		reg := regexp.MustCompile("^[.]")

		for _, f := range files {
			fmt.Println(f.Name(), "---", f.Mode())
			if !reg.Match([]byte(f.Name())) {
				file_names = append(file_names, f.Name())
			} else {
				fmt.Println(f.Name())
			}
		}

		file_data.Name = file_names
		if path != "" {
			file_data.PathValue = path
		} else {
			file_data.PathValue = pwd
		}

		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}

		err = t.Execute(w, file_data)
		if err != nil {
			panic(err)
		}
	}

}

func addfile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("in addFile function...")
	r.ParseForm()

	path := r.FormValue("pathboxvalue")
	fmt.Println(path)

	reg1 := regexp.MustCompile("[/]$")
	res := reg1.Match([]byte(path))
	fmt.Println(res)

	if !res {
		if path != "" {
			path = path + "/"
		}
	}

	extention := "."
	fmt.Println("path:", path, "+", extention)

	file_names := make([]string, 0)
	var file_data FilesList

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	files, err := ioutil.ReadDir(path + extention)
	if err != nil {
		if path != "" {
			file_data.PathValue = path
		} else {
			file_data.PathValue = pwd
		}
		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}

		err = t.Execute(w, file_data)
		if err != nil {
			panic(err)
		}

		fmt.Println(err)

	} else {

		reg := regexp.MustCompile("^[.]")

		for _, f := range files {
			fmt.Println(f.Name(), "---", f.Mode())
			if !reg.Match([]byte(f.Name())) {
				file_names = append(file_names, f.Name())
			} else {
				fmt.Println(f.Name())
			}
		}

		file_data.Name = file_names
		if path != "" {
			file_data.PathValue = path
		} else {
			file_data.PathValue = pwd
		}

		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}

		err = t.Execute(w, file_data)
		if err != nil {
			panic(err)
		}
	}

}

func addfolder(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/home", showFolder).Methods("Get")
	r.HandleFunc("/home", showFolder).Methods("Post")
	r.HandleFunc("/addfile", addfile).Methods("Get")
	r.HandleFunc("/addfolder", addfolder).Methods("Get")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))

	http.ListenAndServe(":3000", r)

}
