package main

import (
	"fmt"
	"log"
	"net/http"
)

func returnMeta(w http.ResponseWriter, req *http.Request) {
	metaData := "{\"name\": \"Polygon KOR Edu\", \"description\": \"This NFTs are created by Antonio Kim\", " +
		"\"image\": \"http://localhost:8081/00.png\", \"attributes\": [{\"trait_type\": \"Background\", \"value\": " +
		"\"white\"}, {\"trait_type\": \"Body\", \"value\": \"maroon\"}, {\"trait_type\": \"Eyes\", \"value\": \"standard\"}, " +
		"{\"trait_type\": \"Head Gear\", \"value\": \"std_lord\"}, {\"trait_type\": \"Held Item\", \"value\": \"nut\"}, " +
		"{\"trait_type\": \"Hands\", \"value\": \"standard\"}]}"
	fmt.Fprintf(w, metaData)
}

func main() {
	fmt.Println("Server Starts...")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/metaNFTs", returnMeta)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
