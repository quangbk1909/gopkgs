package httpx

func IsSuccessStatusCode(code int) bool {
	return code >= 200 && code < 300
}
