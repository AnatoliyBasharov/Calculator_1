package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


func romanToArabic(roman string) (int, error) {
    romanNumerals := map[rune]int{
       'I': 1,
       'V': 5,
       'X': 10,
       'L': 50,
       'C': 100,
       'D': 500,
       'M': 1000,
    }

    var result int
    prevValue := 0

    for i := len(roman) - 1; i >= 0; i-- {
       value := romanNumerals[rune(roman[i])]
       if value < prevValue {
          result -= value
       } else {
          result += value
       }
       prevValue = value
    }

    return result, nil
}
=
func calculate(a, b int, operator string) (int, error) {
    switch operator {
    case "+":
       return a + b, nil
    case "-":
       return a - b, nil
    case "*":
       return a * b, nil
    case "/":
       if b == 0 {
          return 0, fmt.Errorf("division by zero")
       }
       return a / b, nil
    default:
       return 0, fmt.Errorf("unsupported operator: %s", operator)
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Консольный калькулятор")
    fmt.Println("Для выхода введите 'exit'")

    for {
       fmt.Print("Введите выражение (например, 2 + 3): ")
       input, _ := reader.ReadString('\n')
       input = strings.TrimSpace(input)

       if input == "exit" {
          break
       }

       parts := strings.Fields(input)
       if len(parts) != 3 {
          fmt.Println("Неверный формат ввода. Попробуйте еще раз.")
          continue
       }

       num1 := parts[0]
       operator := parts[1]
       num2 := parts[2]


       isRoman1 := strings.ContainsAny(num1, "IVXLCDM")
       isRoman2 := strings.ContainsAny(num2, "IVXLCDM")

       if isRoman1 && isRoman2 {
          fmt.Println("Ошибка: нельзя смешивать римские и арабские цифры в одном выражении.")
          continue
       }

       var a, b int
       var err error

       if isRoman1 {
          a, err = romanToArabic(num1)
          if err != nil {
             fmt.Println("Ошибка при чтении первого числа:", err)
             continue
          }
       } else {
          a, err = strconv.Atoi(num1)
          if err != nil {
             fmt.Println("Ошибка при чтении первого числа:", err)
             continue
          }
       }

       if isRoman2 {
          b, err = romanToArabic(num2)
          if err != nil {
             fmt.Println("Ошибка при чтении второго числа:", err)
             continue
          }
       } else {
          b, err = strconv.Atoi(num2)
          if err != nil {
             fmt.Println("Ошибка при чтении второго числа:", err)
             continue
          }
       }

       result, err := calculate(a, b, operator)
       if err != nil {
          fmt.Println("Ошибка:", err)
          continue
       }

       if isRoman1 || isRoman2 {
          fmt.Printf("Результат: %s\n", arabicToRoman(result))
       } else {
          fmt.Printf("Результат: %d\n", result)
       }
    }
}

func arabicToRoman(arabic int) string {

    return ""
}