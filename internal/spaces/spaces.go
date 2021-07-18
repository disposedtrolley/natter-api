package spaces

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Space struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type spaceResponse struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("read request body: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var space Space
		err = json.Unmarshal(body, &space)
		if err != nil {
			log.Printf("unmarshal request body: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = db.QueryRow(
			"INSERT INTO spaces (name, owner) VALUES ($1, $2) RETURNING space_id",
			space.Name, space.Owner).Scan(&space.ID)
		if err != nil {
			log.Printf("insert new space record: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		uri := fmt.Sprintf("/spaces/%d", space.ID)

		resp := spaceResponse{
			Name: space.Name,
			URI:  uri,
		}

		w.Header().Add("Location", uri)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}
