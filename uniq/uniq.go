package uniq

import (
	"fmt"
	"strings"
)

// Options хранит параметры обработки строк
type Options struct {
	Count      bool // -c
	Duplicates bool // -d
	Unique     bool // -u
	IgnoreCase bool // -i
	SkipFields int  // -f num_fields
	SkipChars  int  // -s num_chars
}

// UniqLines обрабатывает входной слайс строк и возвращает результат в соответствии с опциями
func UniqLines(lines []string, opts Options) []string {
	var result []string
	if len(lines) == 0 {
		return result
	}

	// Функция обработки строки с учетом опций
	process := func(line string) string {
		processed := line

		// Пропустить первые SkipFields полей
		if opts.SkipFields > 0 {
			fields := strings.Fields(processed)
			if len(fields) > opts.SkipFields {
				processed = strings.Join(fields[opts.SkipFields:], " ")
			} else {
				processed = ""
			}
		}

		// Пропустить первые SkipChars символов
		if opts.SkipChars > 0 {
			if len(processed) > opts.SkipChars {
				processed = processed[opts.SkipChars:]
			} else {
				processed = ""
			}
		}

		// Игнорировать регистр
		if opts.IgnoreCase {
			processed = strings.ToLower(processed)
		}

		return processed
	}

	prev := process(lines[0])
	count := 1

	for i := 1; i <= len(lines); i++ {
		var current string
		if i < len(lines) {
			current = process(lines[i])
		}

		if i == len(lines) || current != prev {
			// Решаем, выводить ли prev в результат
			if (opts.Duplicates && count > 1) ||
				(opts.Unique && count == 1) ||
				(!opts.Duplicates && !opts.Unique) {

				lineOut := lines[i-count]
				if opts.Count {
					lineOut = fmt.Sprintf("%d %s", count, lineOut)
				}
				result = append(result, lineOut)
			}

			if i < len(lines) {
				prev = current
				count = 1
			}
		} else {
			count++
		}
	}

	return result
}
