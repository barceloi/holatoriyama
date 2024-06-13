![qr (1).png](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/d1d47922-bd92-48d2-bda4-a841a90a0957/qr_(1).png)

## Introducción

Vengo del mundo del arte.

Entré a la programación por un meetup de creative coding.

Comencé a trabajar como programador pre-junior

Ahora uso algo de programación, pero principalmente uso herramientas low code (webflow) y mas enfocado al marketing. SEO, Landings, Home, Analytics

Me topé con una comunidad llamada SmallBets. Se resume en hacer pequeñas apuestas, distribuir el riesgo. Pequeños experimentos que tengan el potencial de crecer.

Y viendo a indie hackers como Marc Lou, Arvid Kahl, Jordan O’Connor y Lane Wagner, decidi profundizar en el stack de golang y htmx. Como una forma de crecer técnicamente y tambénd e construir pequeñas herramientas de una forma que me da la sensacion de robustez, agilidad y que usa un paradigma que me resulta un poco más intuitivo que el que se está usando en las aplicaciones modernas, donde javascript es el amo y señor.

## Esta es una pequeña introducción a este stack.

### Sobre HTMX

HTMX se basa en el paradigma HATEOAS

Hypermedia as the Engine of Application State

podemos usar htmx con cualquier lenguaje de backend

### Sobre Golang

![Untitled](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/34501487-9f3b-40d3-aea5-b68e646a8aa4/Untitled.png)

## Requisitos

1. Tener un editor instalado p.ej: https://code.visualstudio.com/
    1. si usas vs code, instalar extensión oficial de go
    
    ![Untitled](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/f1f37d76-5eca-407e-b22e-d17ae24512ca/Untitled.png)
    

1. Tener una terminal basada en unix
- si estás en windows.. primero agregar WSL https://learn.microsoft.com/en-us/windows/wsl/install
    - instalar la distribución que viene por defecto (Ubuntu)

- Linux y Mac están ok

1. Tener instalado Golang

Dos opciones para instalarlo

1. Instalación oficial - https://go.dev/doc/install 
2. Webi - https://webinstall.dev/golang/

asegurarte que tienes instalada al menos la versión 1.20 de go

lo puedes ver ingresando este comando en la terminal

```bash
go version
```

si no funciona, prueba cerrar y volver a abrir la terminal

si sigues teniendo problemas para instalar go, intenta googleando o chatgptiando el error que te aparezca.

## **Sección 1: Introducción y Preparación**

Primero, tener instalado golang

1. crear directorio

```bash
mkdir holatoriyama
cd holatoriyama
```

1. Crear proyecto

```bash
go mod init github.com/barceloi/holatoriyama

cat go.mod
```

si no tienes github usar cualquier nombre example.com/username/holatoriyama

1. crear archivo main.go

```bash
touch main.go
```

1. abrir vs code

```bash
code .
```

1. escribir en el archivo main.go

```go
package main

import "fmt"

func main() {
	fmt.Println("hola toriyama")
}
```

1. correr el programa

```go
go run .
```

6b. podemos compilar y ejecutar

```go
go build && ./holatoriyama
```

si aparecen errores raros, usar go mod tidy

1. Agregar air

instalar air - para que no haya que reiniciar el servidor a cada momento

```bash
go install [github.com/cosmtrek/air@latest](http://github.com/cosmtrek/air@latest)
```

guia de como instalar air https://blog.stackademic.com/setting-up-air-for-live-reload-in-golang-project-c92e6a32bb6f

crear archivo configuracion air

```bash
touch .air.toml
```

```toml
# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

root = "."
tmp_dir = "tmp"

[build]
  # cmd = "make build-api"
  cmd = "go build -o ./tmp/main main.go"
  bin = "tmp/main"
  full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html", "css", "js", "env", "yaml"]
  exclude_dir = ["tmp", "assets", "vendor", "bin", "build", "deploy"]
  include_dir = []
  exclude_regex = ["_test.go"]
  exclude_file = []
  exclude_unchanged = true
  log = "air.log"
  args_bin = []
  stop_on_error = true
  send_interrupt = false
  delay = 1000
  kill_delay = 500

[log]
  time = false

[color]
  main = "magenta"
  watcher = "cyan"
  build = "yellow"
  runner = "green"

[misc]
  clean_on_exit = true
```

probar

```bash
air
```

![Untitled](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/7a358717-2c3f-4313-9289-09a57ea233ab/Untitled.png)

Si no funciona, crear el alias en 

```
nano ~/.zshrc
or
nano ~/.bashrc
```

agregar al final el alias

```
alias air='$(go env GOPATH)/bin/air'
```

ctr+o (grabar) y luego ctrl+x (salir)

## **Sección 2: Estructura del Proyecto**

Aquí, describe la estructura de archivos del proyecto, incluidos los archivos Go y HTML. Explica la relación entre el servidor y el cliente, y cómo interactúan entre sí.

Creamos el servidor y cliente. Cómo se relacionan

Cliente>

![Untitled](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/e6f36f0b-195c-4556-b38a-8201076f0636/Untitled.png)

Servidor>

Vamos a crear un server con 4 handlers, uno para listar, otro para crear, otro para eliminar, y otro para mezclar.

### **Sección 3: Creación del Servidor HTTP**

1. Primero creamos a los personajes, hacemos un struct y luego un map.

```go
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
```

1. Creamos el primer handler para enviar el listado, el endpoint, y un handler para servir las imágenes

```go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	// lastID = 7
)

func main() {
	fmt.Println("hello toriyama")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.Execute(w, characters)
	}

	http.HandleFunc("/", handler1)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

```

### **Sección 4: Creación del Cliente +** Mostrar personajes

1. creamos el index.html

agregamos la estructura base de html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>Hola Toriyama</title>
</head>
<body>
    <h1>Hola Toriyama</h1>
    
</body>
</html>
```

1. Hacemos la estructura del html, para el listado y para el form

```html
<body>
    <div class="mt-10 mx-auto max-w-7xl  ">
        <h1 class="flex justify-center">Hola Toriyama</h1>
        <div class="mt-6 grid grid-cols-2 gap-4">
            <div>Personajes
                
            </div>
            <div>Agregar</div>
        </div>
    </div>
    
</body>
```

1. Creamos el listado para que se muestren los personajes (sin htmx por ahora)

```html
	<div>
        <h2>Personajes</h2>
        <ul role="list" id="character-list">
          {{ range .Characters }}
            {{ block "character" .}}
            <li style="display: flex; align-items: center; gap: 1rem;" class="character" id="{{ .Id }}">
              <input id="select-{{ .Id }}" name="selected_characters" value="{{ .Name }}" type="checkbox">
              <img style="width: 4rem;" src={{ .Image }} alt="" />
              <div>
                <p>
                  {{ .Name}}
                </p>
                <p>
                  {{ .Race }}
                </p>
                <button>Eliminar</button>
              </div>
            </li>
            {{ end }}
          {{ end }}
        </ul>

      </div>
```

1. agregamos la carpeta con imágenes

```html
mkdir images
```

link we transfer para imágenes: https://we.tl/t-pdKj5IlIpg

código QR

![qr (1).png](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/b0b9f0d4-f986-4603-8478-6f7124fa9fea/qr_(1).png)

Con eso ya se muestran los personajes que estamos mandando desde el server al template.

![Untitled](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/f2dcf717-daed-43d4-aa57-bb65e586c667/Untitled.png)

### Hacer el handler y endpoint para agregar un personaje y luego conectarlo con htmx

1. Creamos el form para crear personajes en el cliente

```html
<div>
        <h2>Agregar Personaje</h2>
        <form hx-post="/add/" hx-target="#character-list" hx-swap="beforeend" hx-indicator="#spinner">

            <div>
              <div>
                <label for="name">Nombre</label>
                <div>
                  <input type="text" name="name" id="name" />
                </div>
              </div>
            
              <div>
                <label for="race">Especie</label>
                <div>
                  <input type="text" name="race" id="race"/>
                </div>
              </div>

              <div>
                <label for="image">Imagen</label>
                <div>
                  <input type="text" name="image" id="image"/>
                </div>
              </div>
            </div>

          <div>
            <button type="button">
              Cancelar
            </button>
            <button type="submit">
              Agregar
            </button>
          </div>

        </form>

      </div>
```

1. agregamos htmx al head

```html
<script src="https://unpkg.com/htmx.org@1.9.11" integrity="sha384-0gxUXCCR8yv9FM2b+U3FDbsKthCI66oH5IA9fHppQq9DDMHuMauqq1ZHBpJxQ0J0" crossorigin="anonymous"></script
```

1. agregamos htmx al form
- hx-post: a que endpoint vamos a hacer el POST
- hx-target: donde vamos a poner la respuesta de nuestro request
- hx-swap: que tipo de swap vamos a hacer. El default es innerHTML. beforeend es para agregar al final de los otros elementos.
- opcional> poner el hx-indicator, por si quieren hacer el sidequest de un spinner

```html
 <form hx-post="/add/" hx-target="#character-list" hx-swap="beforeend" hx-indicator="#spinner">
```

1. hacer el handler y endpoint en el server

```go
	handler2 := func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(2 * time.Second)
		name := r.PostFormValue("name")
		race := r.PostFormValue("race")
		image := r.PostFormValue("image")

		lastID++
		newCharacter := Character{Name: name, Race: race, Image: image, Id: lastID, isSelected: false}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "character", newCharacter)

	}
	
	...
	http.HandleFunc("/add/", handler2)
	...
```

con eso ya podemos agregar nuevos personajes o elementos

![Untitled](https://prod-files-secure.s3.us-west-2.amazonaws.com/d02f837a-8329-4fca-906d-a147c70aecf9/aa9e3ccf-f542-4adf-9dfb-4d4aac857a85/Untitled.png)

### Eliminar un personaje

1. Crear el handler y endpoint en el server

```go
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
```

1. agregar htmx al botón de eliminar de los personajes
- hace un delete al endpoint, enviando el id.
- el target más cercano con la clase .character, es sí mismo.

```html
<button hx-delete="/delete/{{ .Id }}" hx-target="closest .character" hx-swap="outerHTML">Eliminar</button>
```

con eso se borran los personajes

### Mezclar personajes

1. envolver el listado de personajes en un form

```html
<form>
	<ul>
	...
	</ul>
</form>
```

1. agregar htmx al form
- hacer un get al endpoint mix
- la respuesta que vaya a un div con id mix-response
- que vaya poniendo las respuestas al final de los elementos en ese div
- opcional> un spinner

```html
<form hx-get="/mix/" hx-target="#mix-response" hx-swap="beforeend" hx-indicator="#spinner">
```

1. crear el endpoint para las mezclas

```go
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
```

Side quests

- Agregar un spinner para cuando creamos un personaje
    - https://htmx.org/attributes/hx-indicator/

- ponerle tailwind al head y amononar

```html
 <script src="https://cdn.tailwindcss.com"></script>
```

- Crear más fusiones
- Jugar con algún random
- revivir a Bora