package advance

import "fmt"

type DevLanguage interface {
	study()
}

type CLanguage struct {
	book string
}

// 等同于　func setBook(c *CLanguage, book string)
func (c *CLanguage) setBook(book string) {
	c.book = book
}

// func study(c CLanguage) ()
func (c CLanguage) study() {
	fmt.Printf("study c language, read %s\n", c.book)
}

/*
	type JavaLanguage struct {
		book string
	}

	func (java JavaLanguage) study() {
		fmt.Printf("study java language, read %s\n", java.book)
	}
*/
func Test_receiver() {
	c := CLanguage{
		book: "Ｃ语言设计与实现",
	}
	c.setBook("Ｃ语言设计与实现2")
	c.study()

	/*
		java := JavaLanguage{
			book: "Java语言设计与实现",
		}
		java.study()
	*/
}
