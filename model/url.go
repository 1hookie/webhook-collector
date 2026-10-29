package model

import "strings"

// Чистая функция: результат зависит только от аргументов,
// ничего не читает снаружи и ничего не меняет — поэтому её легко тестировать.
func HookURL(baseURL, id string) string {
	base := strings.TrimRight(baseURL, "/")
	id = strings.TrimLeft(id, "/")
	return base + "/wh/" + id
}

func (h Hook) URL(baseURL string) string {
	return HookURL(baseURL, h.ID)
}
