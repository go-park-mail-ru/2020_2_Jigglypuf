package cookie

import "net/http"

func checkCookie(r *http.Request) bool{
	_, err := r.Cookie("session_id")

	if err != nil{
		return false
	}
	return true
}
