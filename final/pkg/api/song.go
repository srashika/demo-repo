package api

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

//name, singer, track_id
type Song struct{
	//defining the song
	Name string
	Singer string
	ID string
}

//ToJSON to be used for marshalling of SONG type.
func (s Song) ToJSON() []byte{
	ToJSON, err :=json.Marshal(s)
	if err !=nil{
		panic(err)

	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Song {
	song := Song{}
	err := json.Unmarshal(data, &song)
	if err != nil {
		panic(err)
	}
	return song
}

var songs = map[string]Song{
	"01":Song{Name: "Cold Mess", Singer: "Prateek Kuhad", ID: "01"},
	"02":Song{Name: "Papercut", Singer: "Linkin Park", ID: "02"},
	"03":Song{Name: "IFLY", Singer: "Bazzi", ID: "03"},
}
// AllBooks returns a slice of all books
func AllSongs() []Song {
	values := make([]Song, len(songs))
	idx := 0
	for _, song := range songs {
		values[idx] = song
		idx++
	}
	return values
}

func SongsHandleFunc(w http.ResponseWriter, r *http.Request){
	//for allsongs: assuming we have a get request for api/songs
	switch method := r.Method; method {
	case http.MethodGet:
		songs := AllSongs()
		writeJSON(w, songs)

	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		song := FromJSON(body)
		id, created := CreateSong(song)
		if created {
			w.Header().Add("Location", "/api/song/"+id)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
    default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
    }
}
func SongHandleFunc(w http.ResponseWriter, r *http.Request){
	id := r.URL.Path[len("/api/song/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		song, found := GetSong(id)
		if found {
			writeJSON(w, song)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		song := FromJSON(body)
		exists := UpdateSong(id, song)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteSong(id)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	s, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(s)
}

func CreateSong(song Song) (string, bool) {
	_, exists := songs[song.ID]
	if exists {
		return "", false
	}
	songs[song.ID] = song
	return song.ID, true
}

func GetSong(id string) (Song, bool) {
	song, found := songs[id]
	return song, found
}

func UpdateSong(id string, song Song) bool {
	_, exists := songs[id]
	if exists {
		songs[id] = song
	}
	return exists
}

func DeleteSong(id string) {
	delete(songs, id)
}