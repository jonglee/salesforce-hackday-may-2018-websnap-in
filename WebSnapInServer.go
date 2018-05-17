package main
import (
	"net/http"
	"os"
)


func main() {
	port := os.Getenv("PORT")
	http.Handle("/", http.FileServer(http.Dir("./html")))
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		panic(err)
	}
}