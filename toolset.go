package clio

func SetHeader (key, value string) {
    Context().ResponseWriter.Header().Set(key, value)
}

func SetCookie (key, value string) {}

func DeleteCookie (key string) {}

func Cookies () map[string]string {
  cookies := make(map[string]string)
  return cookies
}

// vim: noai:ts=4:sw=4
