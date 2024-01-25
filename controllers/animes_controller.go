package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type AnimeTitle struct {
	Link string `json:"link"`
	Text string `json:"text"`
}
type Anime struct {
	Studio      string     `json:"studio"`
	Genres      []string   `json:"genres"`
	Hype        int        `json:"hype"`
	Description string     `json:"description"`
	Title       AnimeTitle `json:"title"`
	Start_date  string     `json:"start_date"`
	Image       string     `json:"image"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ListAnimes(w, r)
	case http.MethodPost:
		CreateAnime(w, r)
	case http.MethodDelete:
		DeleteAnime(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func DeleteAnime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func CreateAnime(w http.ResponseWriter, r *http.Request) {
	var anime Anime = Anime{
		Studio:      r.FormValue("studio"),
		Genres:      []string{r.FormValue("genres")},
		Hype:        0,
		Title:       AnimeTitle{Link: r.FormValue("link"), Text: r.FormValue("title")},
		Description: r.FormValue("description"),
		Start_date:  r.FormValue("start_date"),
		Image:       r.FormValue("image")}

	var animes []Anime
	file, err := os.ReadFile("data/data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not read file data/data.json"))
		return
	}

	err = json.Unmarshal(file, &animes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not read file data/data.json"))
		return
	}
	animes = append(animes, anime)
	html := fmt.Sprintf(`<div class='anime-card'><div class='anime-div-image'><img class='anime-image' src=%s alt='Imagem do anime'></div><div class='anime-info'><p class='anime-title'><a href=%s>%s</a></p><p class='anime-studio'>Studio:%s</p><p class='anime-start-date'>Started At:%s</p><p class='anime-description'>%s</p></div></div>`, anime.Image, anime.Title.Link, anime.Title.Text, anime.Studio, anime.Start_date, anime.Description)
	tmpl, _ := template.New("anime-card").Parse(html)
	tmpl.Execute(w, nil)
}
func ListAnimes(w http.ResponseWriter, r *http.Request) {
	var animes []Anime
	file, err := os.ReadFile("data/data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not read file data/data.json"))
		return
	}

	err = json.Unmarshal(file, &animes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not read file data/data.json"))
		return
	}

	template, err := template.ParseFiles("templates/list.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could parse HTMX Templates"))
	}
	template.Execute(w, animes)
}
