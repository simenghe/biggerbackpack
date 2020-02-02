package routes
import(
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	bagController "Backpack/controllers"
)
// HandleReq handles our main route
func HandleReq() {
	fmt.Println("Handling the requests!")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", bagController.FetchTestBags)
	myRouter.HandleFunc("/health",bagController.Health)
	myRouter.HandleFunc("/getbags",bagController.FetchBags)
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}