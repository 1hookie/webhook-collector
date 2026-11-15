package envcfg

import "strconv"

type App struct {
	Host string
	Port int
}

// FillFromEnv переопределяет поля cfg значениями из env.
// Пустые, отсутствующие и некорректные значения молча игнорируются.
func FillFromEnv(cfg *App, env map[string]string) {
	if host := env["HC_HOST"]; host != "" {
		cfg.Host = host
	}

	// Для отсутствующего ключа map вернёт "", и Atoi вернёт ошибку —
	// одна проверка закрывает и «нет», и «пусто», и «не число».
	if port, err := strconv.Atoi(env["HC_PORT"]); err == nil && port >= 1 && port <= 65535 {
		cfg.Port = port
	}
}
