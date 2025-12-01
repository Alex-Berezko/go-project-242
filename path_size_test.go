package pz

import "testing"

func TestGetSize(t *testing.T) {
	//Arange -подготовка, expected (ожидание) результат, данные для функции
	size, err := GetSize("ooo.go")
	if err != nil {
		t.Fatalf("Ошибка, GetSize  %v", err)
	}

	//Act - выполняет функцию
	//Assert - проверка результата с expected (ожидание)
	expected := int64(541)
	if size != expected {
		t.Errorf("получили size: %d, ожидали expected: %d", size, expected)
	}
}
