package hellococsup2021tw

import (
	"fmt"
	"net/http"
)

func HelloCocsup2021tw(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, COSCUP 2021 TW!")
}
