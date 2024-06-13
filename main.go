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
			{Id: 1, Name: "Karin", Race: "Gato", Image: "/images/karin.png"},
			{Id: 2, Name: "Goku", Race: "Saiyajin", Image: "images/goku.png"},
			{Id: 3, Name: "Camara del tiempo", Race: "Templo", Image: "images/camara.jpg"},
			{Id: 4, Name: "Esfera", Race: "Esfera", Image: "images/esfera.jpg"},
			{Id: 5, Name: "Upa", Race: "Humano", Image: "images/upa2.png"},
			{Id: 6, Name: "Goten", Race: "50% Humano, 50% Saiyajin", Image: "images/goten.webp"},
			{Id: 7, Name: "Trunks", Race: "50% Humano, 50% Saiyajin", Image: "images/trunks.webp"},
		},
	}
	lastID = 7
)

func main() {
	fmt.Println("hola toriyama")

	http.HandleFunc("/", handlerCharactersGet)
	http.HandleFunc("/add/", handlerCharactersCreate)
	http.HandleFunc("/delete/", handlerCharactersDelete)
	http.HandleFunc("/mix/", handler4)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	log.Fatal(http.ListenAndServe(":7000", nil))
}

func handlerCharactersGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, characters)
}

func handlerCharactersCreate(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	name := r.PostFormValue("name")
	race := r.PostFormValue("race")
	image := r.PostFormValue("image")

	lastID++
	newCharacter := Character{Name: name, Race: race, Image: image, Id: lastID, isSelected: false}

	characters["Characters"] = append(characters["Characters"], newCharacter)

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.New("character").Parse(`
		<li style="display: flex; align-items: center; gap: 1rem;" class="character" id="{{ .Id }}">
			<input id="select-{{ .Id }}" name="selected_characters" value="{{ .Name }}" type="checkbox">
			<img style="width: 4rem;" src={{ .Image }} alt="" />
			<div>
				<p>{{ .Name }}</p>
				<p>{{ .Race }}</p>
				<button hx-delete="/delete/{{ .Id }}" hx-target="closest .character" hx-swap="outerHTML" hx-confirm="Estás seguro?">Eliminar</button>
			</div>
		</li>
	`))
	tmpl.Execute(w, newCharacter)
}

func handlerCharactersDelete(w http.ResponseWriter, r *http.Request) {
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

func handler4(w http.ResponseWriter, r *http.Request) {
	selectedCharacters := r.URL.Query()["selected_characters"]
	ballCount := 0

	gokuSelected := false
	chamberSelected := false
	gotenSelected := false
	trunksSelected := false
	for _, character := range selectedCharacters {
		if strings.ToLower(character) == "goku" {
			gokuSelected = true

		} else if strings.ToLower(character) == "camara del tiempo" {

			chamberSelected = true
		} else if strings.ToLower(character) == "goten" {
			gotenSelected = true
		} else if strings.ToLower(character) == "trunks" {
			trunksSelected = true
		} else if strings.HasPrefix(strings.ToLower(character), "esfera") {
			ballCount++
		}
	}

	if gokuSelected && chamberSelected {
		htmlResponse := "<img src='/images/g-super.png' alt='Super Saiyan'>"

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
		htmlResponse := "<img src='/images/dragon.webp' alt='Shenlong'>"

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, htmlResponse)
		return
	}

	fmt.Fprintf(w, "")
}
