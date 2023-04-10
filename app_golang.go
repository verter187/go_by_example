package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Проверяем, соответствует ли шаблон строке
	match, _ := regexp.MatchString("p([a-z]+)ch", "panch")
	fmt.Println(match)

	// Можно скомпилировать оптимизированную структуру Regexp
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Множество методов доступны для этой структуры. Вот
	// тест на совпадение, который мы видели ранее.
	fmt.Println(r.MatchString("peach"))

	// Этот метод находит соответствие для регулярного выражения.
	fmt.Println(r.FindString("punch peach"))

	// Этот метод также находит первое совпадение, но возвращает
	// начальный и конечный индексы совпадения вместо текста.
	fmt.Println(r.FindStringIndex("punch peach"))

	// Варианты Submatch включают в себя информацию как о совпадениях с полным шаблоном, так и о совпадениях с частями шаблона.
	// Например, эта конструкция вернет информацию как для p([a-z]+)ch, так и для ([a-z]+).
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Точно так же это возвратит информацию об индексах совпадений и подсовпадений.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Метод All применяется ко всем совпадениям на входе, а не только к первому.
	// Например, чтобы найти все совпадения для регулярного выражения.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Этот метод All доступен и для других функций, которые мы видели выше.
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Указание неотрицательного целого числа в качестве второго аргумента для этих функций ограничит количество совпадений.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// В примерах выше были строковые аргументы и использовались такие имена, как MatchString.
	// Мы также можем предоставить []byte аргументы и удалить String из имени функции.
	fmt.Println(r.Match([]byte("peach")))

	// При создании констант с регулярными выражениями вы можете использовать MustCompile, как аналог Compile.
	// Обычный Compile не будет работать для констант, потому что он возвращает 2 значения.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Пакет regexp также можно использовать для замены подмножеств строк другими значениями.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
