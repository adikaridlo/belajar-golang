package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"models/user"
	"models/berita"
)

func helloWord(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Word !")
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWord).Methods("GET")
	myRouter.HandleFunc("/users", users.AllUser).Methods("GET")
	myRouter.HandleFunc("/create-user", users.NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", users.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{id}", users.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/hallo", users.Testing).Methods("PUT")
	myRouter.HandleFunc("/politik", artikel.Politik).Methods("GET")
	log.Fatal(http.ListenAndServe(":9999", myRouter))
}
func main()  {
	fmt.Println("Starting golang.. .")
	handleRequests()
}