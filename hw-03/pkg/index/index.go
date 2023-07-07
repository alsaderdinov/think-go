// Package index предоставляет простую службу индексации
// целочисленные значения, связанные с конкретными словами.
package index

import (
	"strings"
)

// Service - служба индексации.
type Service struct {
	data map[string][]int
}

// New конструктор службы индексации.
func New() *Service {
	return &Service{
		data: make(map[string][]int),
	}
}

// Add добавляет заданное целочисленное значение к индексу для каждого слова в предоставленной строке.
// Разбивает входную строку на слова, преобразует их в строчные и связывает
// заданное целочисленное значение с каждым словом во внутреннем хранилище данных.
func (svc *Service) Add(s string, n int) {
	words := strings.Fields(strings.ToLower(s))
	for _, w := range words {
		svc.data[w] = append(svc.data[w], n)
	}
}

// Find извлекает целочисленные значения, связанные с заданным словом.
// Возвращается фрагмент, содержащий все целочисленные значения, связанные с данным словом.
// Если слово не найдено в индексе, возвращается пустой фрагмент.
func (svc *Service) Find(s string) []int {
	return svc.data[strings.ToLower(s)]
}
