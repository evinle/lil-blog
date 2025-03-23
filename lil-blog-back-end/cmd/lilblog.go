package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("welcome"))
        hwMap  := map[string]string{"hello": "world"}
        // jsonHwMap, error := json.MarshalIndent(hwMap, "", " ")
        jsonHwMap, error := json.Marshal(hwMap)

        if(error != nil){
            panic(error)
        } 

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonHwMap)

	})

	http.ListenAndServe(":3000", r)
}
