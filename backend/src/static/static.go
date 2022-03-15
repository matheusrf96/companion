package static

import "net/http"

func CompanionJavascript(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../frontend/dist/cmp.js")
}
