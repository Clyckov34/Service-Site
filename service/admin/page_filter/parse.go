package page_filter

import (
	"strconv"
	"strings"
)

// ParserNumberTask распарсивает данные для работы с ним
func ParserNumberTask(numberTask string) (Parser, error) {
	if numberTask != "" {
		data := strings.Split(numberTask, "-")

		_, err := strconv.Atoi(data[1])
		if err != nil {
			return Parser{}, nil
		}
	
		return Parser{
			Key:    strings.ToUpper(data[0]),
			Number: data[1],
		}, nil
	}

	return Parser{}, nil
}
