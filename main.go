package main
import (
    "io"
    "net/http"
)
func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)
    err := http.ListenAndServe(":3333", nil)
    checkerr(err)
}
func checkerr(err error) {
    panic("unimplemented")
}
func getRoot(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "This is a website\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello HTTP\n")
}
