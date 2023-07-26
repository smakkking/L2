package pattern

import "fmt"

type Compressor interface {
	compress(file_name string)
}

type ZIP_Compression struct{}

func (c *ZIP_Compression) compress(file string) {
	fmt.Println("ZIP compression")
}

type ARJ_Compression struct{}

func (c *ARJ_Compression) compress(file string) {
	fmt.Println("ARJ compression")
}

type RAR_Compression struct{}

func (c *RAR_Compression) compress(file string) {
	fmt.Println("RAR compression")
}

// класс, который будет выбирать "стратегию"
type CompressionClass struct {
	c Compressor
}

func (c *CompressionClass) compress(file string) {
	c.compress(file)
}

func NewCompressionClass(local_c Compressor) *CompressionClass {
	return &CompressionClass{
		c: local_c,
	}
}

func (comp_c *CompressionClass) ChangeStrategy(c Compressor) {
	comp_c.c = c
}

func MainFunc() {
	p := NewCompressionClass(new(ZIP_Compression))
	p.compress("file.txt")
	p.ChangeStrategy(new(RAR_Compression))
	p.compress("file.txt")
}
