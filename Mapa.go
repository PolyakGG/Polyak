myMap := make(map[int]int) // Создаем пустой словарь (map).

for i := 0; i < 10; i++ {
	var val int
	_, err := fmt.Scan(&val)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		continue // В случае ошибки продолжаем цикл.
	}

	// Проверяем наличие значения в map-е.
	if result, exists := myMap[val]; exists {
		// Если значение есть, выводим его.
		fmt.Printf("%d ", result)
	} else {
		// Если значения нет, вызываем функцию work и записываем результат в map.
		result := work(val)
		myMap[val] = result
		fmt.Printf("%d ", result)
	}
}
// Check commit