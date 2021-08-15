package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	res := []string{}

	// 1. Разделим строку на массив слов и отсортируем
	words := getSortedSliceFromStr(str)

	// 2. Создать структуру: слово, частота
	type wordsAndCountStruct struct {
		word  string
		count int
	}

	wordsAndCountSlice := []wordsAndCountStruct{}

	currentWord := ""
	currentCount := 0
	for _, value := range words {
		if value != currentWord {
			// Запишем текущее слово и количество и обнулим
			if currentWord != "" {
				wordsAndCountSlice = append(wordsAndCountSlice, wordsAndCountStruct{currentWord, currentCount})
			}

			currentWord = value
			currentCount = 1
		} else {
			currentCount++
		}
	}
	// Запишем последнюю итерацию
	if currentWord != "" {
		wordsAndCountSlice = append(wordsAndCountSlice, wordsAndCountStruct{currentWord, currentCount})
	}

	// 3. Отсортировать структуру по количеству, если количество одинаковое - по слову
	sort.Slice(wordsAndCountSlice, func(i, j int) bool {
		if wordsAndCountSlice[i].count == wordsAndCountSlice[j].count {
			return wordsAndCountSlice[i].word < wordsAndCountSlice[j].word
		}
		return wordsAndCountSlice[i].count > wordsAndCountSlice[j].count
	})

	// 4. Получить из слайса структуры слайс результата
	for _, wordsStruct := range wordsAndCountSlice {
		res = append(res, wordsStruct.word)
	}

	// 5. Если слов меньше 10
	limit := 10
	if len(res) < 10 {
		limit = len(res)
	}

	return res[0:limit]
}

func getSortedSliceFromStr(str string) []string {
	words := strings.Fields(str)

	// Отсортируем слова
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	return words
}
