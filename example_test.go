package num2persian_test

import (
	"fmt"
	"math/big"

	"github.com/pinkorca/num2persian"
)

func ExampleConvert() {
	fmt.Println(num2persian.Convert(0))
	fmt.Println(num2persian.Convert(21))
	fmt.Println(num2persian.Convert(1234))
	fmt.Println(num2persian.Convert(-500))
	// Output:
	// صفر
	// بیست و یک
	// هزار و دویست و سی و چهار
	// منفی پانصد
}

func ExampleConvert_largeNumbers() {
	fmt.Println(num2persian.Convert(1000000))       // million
	fmt.Println(num2persian.Convert(1000000000))    // billion
	fmt.Println(num2persian.Convert(1000000000000)) // trillion
	// Output:
	// یک میلیون
	// یک میلیارد
	// یک تریلیون
}

func ExampleConvertBigInt() {
	n := new(big.Int)
	n.SetString("1000000000000000000000000000000000", 10) // 10^33 (decillion)
	fmt.Println(num2persian.ConvertBigInt(n))
	// Output:
	// یک دسیلیون
}

func ExampleConvertFloat() {
	fmt.Println(num2persian.ConvertFloat(12.5, 1))
	fmt.Println(num2persian.ConvertFloat(3.14, 2))
	fmt.Println(num2persian.ConvertFloat(99.99, 2))
	// Output:
	// دوازده ممیز پنج
	// سه ممیز چهارده
	// نود و نه ممیز نود و نه
}

func ExampleConvertString() {
	result, _ := num2persian.ConvertString("1234567")
	fmt.Println(result)

	result2, _ := num2persian.ConvertString("42.5")
	fmt.Println(result2)
	// Output:
	// یک میلیون و دویست و سی و چهار هزار و پانصد و شصت و هفت
	// چهل و دو ممیز پنج
}

func ExampleConvertOrdinal() {
	fmt.Println(num2persian.ConvertOrdinal(1))
	fmt.Println(num2persian.ConvertOrdinal(2))
	fmt.Println(num2persian.ConvertOrdinal(3))
	fmt.Println(num2persian.ConvertOrdinal(21))
	fmt.Println(num2persian.ConvertOrdinal(100))
	// Output:
	// اول
	// دوم
	// سوم
	// بیست و یکم
	// صدم
}

func ExampleToToman() {
	fmt.Println(num2persian.ToToman(1000))
	fmt.Println(num2persian.ToToman(1500000))
	// Output:
	// هزار تومان
	// یک میلیون و پانصد هزار تومان
}

func ExampleToRial() {
	fmt.Println(num2persian.ToRial(10000))
	fmt.Println(num2persian.ToRial(15000000))
	// Output:
	// ده هزار ریال
	// پانزده میلیون ریال
}
