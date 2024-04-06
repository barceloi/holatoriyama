package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Character struct {
	Id         int
	Name       string
	Race       string
	Image      string
	isSelected bool
}

var (
	characters = map[string][]Character{
		"Characters": {
			{Id: 1, Name: "Karin", Race: "Gato", Image: "/images/karin.png", isSelected: false},
			{Id: 2, Name: "Goku", Race: "Saiyajin", Image: "images/goku.png", isSelected: false},
			{Id: 3, Name: "Camara del tiempo", Race: "Templo", Image: "images/camara.jpg", isSelected: false},
			{Id: 4, Name: "Esfera", Race: "Esfera", Image: "images/esfera.jpg", isSelected: false},
			{Id: 5, Name: "Upa", Race: "Humano", Image: "images/upa2.png", isSelected: false},
			{Id: 6, Name: "Goten", Race: "50% Humano, 50% Saiyajin", Image: "images/goten.webp", isSelected: false},
			{Id: 7, Name: "Trunks", Race: "50% Humano, 50% Saiyajin", Image: "images/trunks.webp", isSelected: false},
		},
	}
	lastID = 7
)

func main() {
	fmt.Println("hola toriyama")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.Execute(w, characters)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		name := r.PostFormValue("name")
		race := r.PostFormValue("race")
		image := r.PostFormValue("image")

		lastID++
		newCharacter := Character{Name: name, Race: race, Image: image, Id: lastID, isSelected: false}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "character", newCharacter)

	}

	handler3 := func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/delete/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid character ID", http.StatusBadRequest)
			return
		}

		for i, character := range characters["Characters"] {
			if character.Id == id {
				characters["Characters"] = append(characters["Characters"][:i], characters["Characters"][i+1:]...)
				return
			}
		}

		http.Error(w, "Character not found", http.StatusNotFound)
	}

	handler4 := func(w http.ResponseWriter, r *http.Request) {
		selectedCharacters := r.URL.Query()["selected_characters"]
		ballCount := 0

		gokuSelected := false
		camaraSelected := false
		gotenSelected := false
		trunksSelected := false
		for _, character := range selectedCharacters {
			if strings.ToLower(character) == "goku" {
				gokuSelected = true
				fmt.Println("goku selected")
			} else if strings.ToLower(character) == "camara del tiempo" {
				fmt.Println("chamber selected")
				camaraSelected = true
			} else if strings.ToLower(character) == "goten" {
				gotenSelected = true
			} else if strings.ToLower(character) == "trunks" {
				trunksSelected = true
			} else if strings.HasPrefix(strings.ToLower(character), "esfera") {
				ballCount++
			}
		}

		if gokuSelected && camaraSelected {
			htmlResponse := "<img src='/images/g-super.png' alt='Super Image'>"

			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, htmlResponse)
			return
		} else if gotenSelected && trunksSelected {
			htmlResponse := "<img src='/images/gotenks.webp' alt='Gotenks'>"

			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, htmlResponse)
			return

		}

		if ballCount == 7 {
			htmlResponse := "<img src='/images/dragon.webp' alt='Super Image'>"

			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, htmlResponse)
			return
		}

		fmt.Fprintf(w, "")
	}

	http.HandleFunc("/", handler1)
	http.HandleFunc("/add/", handler2)
	http.HandleFunc("/delete/", handler3)
	http.HandleFunc("/mix/", handler4)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	log.Fatal(http.ListenAndServe(":7000", nil))
}
