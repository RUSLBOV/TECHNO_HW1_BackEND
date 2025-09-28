package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"uniq" // замените на актуальный путь к пакету uniq
)

func main() {
	var opts uniq.Options

	// Флаги
	flag.BoolVar(&opts.Count, "c", false, "Count occurrences")
	flag.BoolVar(&opts.Duplicates, "d", false, "Show duplicate lines")
	flag.BoolVar(&opts.Unique, "u", false, "Show unique lines")
	flag.BoolVar(&opts.IgnoreCase, "i", false, "Ignore case")
	flag.IntVar(&opts.SkipFields, "f", 0, "Skip N fields")
	flag.IntVar(&opts.SkipChars, "s", 0, "Skip N characters")
	flag.Parse()

	// Проверка конфликтующих флагов
	if (opts.Count && opts.Duplicates) || (opts.Count && opts.Unique) || (opts.Duplicates && opts.Unique) {
		fmt.Fprintln(os.Stderr, "Error: options -c, -d, -u are mutually exclusive")
		os.Exit(1)
	}

	args := flag.Args()
	var inputFile, outputFile string
	if len(args) > 0 {
		inputFile = args[0]
	}
	if len(args) > 1 {
		outputFile = args[1]
	}

	// Чтение входных данных
	reader := os.Stdin
	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		reader = f
	}

	lines, err := readLines(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Обработка
	output := uniq.UniqLines(lines, opts)

	// Запись результата
	writer := os.Stdout
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		writer = f
	}

	err = writeLines(writer, output)
	if err != nil {
		log.Fatal(err)
	}
}

// readLines считывает все строки из io.Reader
func readLines(reader *os.File) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines записывает строки в io.Writer
func writeLines(writer *os.File, lines []string) error {
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}
	return nil
}
