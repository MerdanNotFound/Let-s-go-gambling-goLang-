package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	str1 := scanner.Text()
	_ = scanner.Scan()
	str2 := scanner.Text()
	_ = scanner.Scan()
	str3 := scanner.Text()
	_ = scanner.Scan()
	str4 := scanner.Text()
	fmt.Println(str2, str1, str3, str1, str4)
}

/* Кастомный разделитель
Напишите программу, которая считывает строку-разделитель и три строки, а затем выводит указанные строки через разделитель.

Формат входных данных
На вход программе подаётся строка-разделитель и три строки, каждая на отдельной строке.

Формат выходных данных
Программа должна вывести введённые три строки через разделитель.

Sample Input 1:

*
Раз
Два
Три
Sample Output 1:

Раз*Два*Три
Sample Input 2:

$
I
like
Go
Sample Output 2:

I$like$Go
*/
