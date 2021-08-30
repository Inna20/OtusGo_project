package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// Структура: слово, частота.
type wordsAndCountStruct struct {
	word  string
	count int
}

// Коллекция, которая реализует интерфейс sort.Interface (которую можно отсортировать с помощью sort.Sort).
type WordsAndCountSlice []wordsAndCountStruct

func (s WordsAndCountSlice) Len() int {
	return len(s)
}

func (s WordsAndCountSlice) Less(i, j int) bool {
	if s[i].count == s[j].count {
		return s[i].word < s[j].word
	}
	return s[i].count > s[j].count
}

func (s WordsAndCountSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Top10(str string) []string {
	// 1. Разделим строку на массив слов и отсортируем
	words := getSortedSliceFromStr(str)

	// 2. Слайс структур слово+количество
	wcSlice := getWordsAndCountSlice(words)

	// 3. Отсортировать структуру по количеству, если количество одинаковое - по слову
	sort.Sort(wcSlice)

	// 4. Получить из слайса структуры слайс результата
	return formatResult(wcSlice)
}

func getWordsAndCountSlice(words []string) WordsAndCountSlice {
	wcSlice := WordsAndCountSlice{}

	currentWord := ""
	currentCount := 0
	for _, value := range words {
		if value != currentWord {
			// Запишем текущее слово и количество и обнулим
			if currentWord != "" {
				wcSlice = append(wcSlice, wordsAndCountStruct{currentWord, currentCount})
			}

			currentWord = value
			currentCount = 1
		} else {
			currentCount++
		}
	}
	// Запишем последнюю итерацию
	if currentWord != "" {
		wcSlice = append(wcSlice, wordsAndCountStruct{currentWord, currentCount})
	}

	return wcSlice
}

func getSortedSliceFromStr(str string) []string {
	words := strings.Fields(str)

	// Отсортируем слова
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	return words
}

func formatResult(wcSlice WordsAndCountSlice) []string {
	res := []string{}

	for _, wordsStruct := range wcSlice {
		res = append(res, wordsStruct.word)
	}

	// Если слов меньше 10
	limit := 10
	if len(res) < 10 {
		limit = len(res)
	}

	return res[0:limit]
}
