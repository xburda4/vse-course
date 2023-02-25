package vanilla

import "net/http"

func Router() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/users", HandleUsers)
	r.HandleFunc("/users/", HandleUser)

	return r
}