package main
import(
	"fmt"
	"net/http"
	"github.com/KabukaK/pipeboard/internal/config"

)

func main(){
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"status": "ok"}`)
})

cfg := config.Load()
addr:= ":" + cfg.Port

fmt.Println("API listening on", addr)
err:= http.ListenAndServe(addr, nil)
if err != nil {
	panic(err)
}
}