package clio

func SetHeader (key, value string) {
    Context().ResponseWriter.Header().Set(key, value)
}

func SetCookie (key, value string) {}

func DeleteCookie (key string) {}
