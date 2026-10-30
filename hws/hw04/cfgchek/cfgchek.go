// Задача 5. «Парсинг флагов» — пакет flag через NewFlagSet, чтобы функцию можно было тестировать.
package cfgcheck

import (
	"flag"
	"fmt"
	"io"
)

const defaultConfigFile = "./config.json"

type Options struct {
	ConfigFile string
	Pretty     bool
	Quiet      bool
}

// parseFlags разбирает args (это os.Args[1:], без имени программы).
// Ошибка возвращается вызывающему, а не завершает процесс.
func parseFlags(args []string) (Options, error) {
	var opts Options

	fs := flag.NewFlagSet("cfgcheck", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // не печатать usage в вывод тестов

	fs.StringVar(&opts.ConfigFile, "config", defaultConfigFile, "путь к файлу конфигурации")
	fs.BoolVar(&opts.Pretty, "pretty", false, "печатать отформатированный JSON")
	fs.BoolVar(&opts.Quiet, "quiet", false, "при успехе ничего не печатать")

	if err := fs.Parse(args); err != nil {
		return Options{}, err
	}
	// Всё, что осталось после флагов, — лишние позиционные аргументы.
	if fs.NArg() > 0 {
		return Options{}, fmt.Errorf("unexpected arguments: %v", fs.Args())
	}
	return opts, nil
}
