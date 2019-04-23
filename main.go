package main
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
)
var port = flag.Int("port", 3333, "static server port")
var dir = flag.String("dir", "./", "static server directory")
func Run() {
	flag.Parse()
	portstr := fmt.Sprintf(":%d", *port)
	log.Println("Config", portstr, *dir)
	err := http.ListenAndServe(portstr, handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(*dir))))
	log.Fatal(err)
}
func main() {
	Run()
}
