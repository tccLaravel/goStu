package main

import (
	"bytes"
	"fmt"
	"os"
)
/**
创建缓冲器
1.bytes.NewBuffer
2.bytes.NewBufferString
 */
func createBuf()  {
	buf1 := bytes.NewBuffer([]byte("hello"))
	buf2 := bytes.NewBuffer([]byte{'h','e','l','l','o'})
	buf3 := bytes.NewBufferString("hello")
	// 以上三者等效,输出//hello
	buf4 := bytes.NewBuffer([]byte{})
	buf5 := bytes.NewBufferString("")
	//以上两者等效,输出//""
	fmt.Println(buf1.String(),buf2.String(),buf3.String(),buf4,buf5,1) //hello hello hello   1
}
/**
写入到缓冲器
1.buf.Write([]byte) 将一个byte类型的slice放到缓冲器的尾部
2.buf.WriteString(string) 把一个字符串放到缓冲器的尾部
3.buf.WriteByte(byte) 将一个byte类型的数据放到缓冲器的尾部
4.buf.WriteRune(Rune) 将一个rune类型的数据放到缓冲器的尾部
 */
func writeToBuf()  {
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())
	buf.WriteString(" world")
	fmt.Println(buf.String())
	s := []byte(" tcc ")
	buf.Write(s)
	fmt.Println(buf.String())
	var b byte = '!'
	buf.WriteByte(b)
	fmt.Println(buf.String())
}

//把buf数据写入 IO
func bufWriteToFile()  {
	file,_ := os.Create("text.txt")
	buf := bytes.NewBufferString("hello word tcc")
	buf.WriteTo(file)
	//或者使用写入
	fmt.Fprintf(file,buf.String())
}

/**
读入缓冲器
把IO 数据 读入 buf
 */
func readFromBuffer()  {
	file,_ := os.Create("text.txt")
	buf := bytes.NewBufferString("hi ")
	buf.ReadFrom(file)
	fmt.Println(buf.String())
}

func readBuffer()  {
	//readBuf()
	//readByte()
	//readRune()
	readBytesAndStrings()
	readFromBuffer()
}


//ReadBytes ,ReadString 方法，需要一个byte作为分隔符，读的时候从缓冲器里找出第一个出现的分隔符，缓冲器头部开始到分隔符之间的byte返回。
//可以用于处理字符的分割,只是返回值不同一个返回 byte  一个返回 string
func readBytesAndStrings()  {
	var d byte = 'e'
	buf := bytes.NewBufferString("你好esmieth")
	fmt.Println(buf.String()) //你好esmieth
	b,_ := buf.ReadBytes(d)  //读到分隔符，并返回给b
	fmt.Println(buf.String())  //smieth
	fmt.Println(string(b)) //你好e
	s,_  := buf.ReadString(d)   //读取到分隔符，并返回给b
	fmt.Println(buf.String())  //th
	fmt.Println(s) //smie
}

func readRune()  {
	buf := bytes.NewBufferString("我i唐绵绵")
	fmt.Println(buf.String())
	r, size, _ := buf.ReadRune()
	fmt.Println(buf.String())
	fmt.Println(string(r))
	fmt.Println(size) // "我"作为utf8存储占3个byte
	r, size, _ = buf.ReadRune()
	fmt.Println(buf.String())
	fmt.Println(string(r))
	fmt.Println(size) //
}

/**
一个byte一个byte的读
*/
func readByte()  {
	buf := bytes.NewBufferString("i am tcc") // 注意这里如果是汉字的话,每一个汉字占用的是3byte,那么解析出来的就是单个byte代表的字符,而不是一个汉字
	fmt.Println(buf.String())
	b, _ := buf.ReadByte() // i
	_, _ = buf.ReadByte() // ""
	c, _ := buf.ReadByte()// a
	fmt.Println(string(b))
	fmt.Println(string(c))
}

/**
从buf里面取数据的通用方法
 */
func readBuf()  {
	s1 := []byte("read")
	buff := bytes.NewBuffer(s1)
	s2 := []byte(" tcc Buffer")
	buff.Write(s2)
	fmt.Println(buff.String()) //read From Buffer

	s3 := make([]byte,3)
	_, _ = buff.Read(s3)       //把buff的内容读入到s3，s3的容量为3，读了3个过来
	fmt.Println(buff.String()) //lo world
	fmt.Println(string(s3))    //hel
	_, _ = buff.Read(s3)       //继续读入3个，原来的被覆盖

	fmt.Println(buff.String())     //rom Buffer
	fmt.Println(string(s3))    //"d F"
}

func nextBuf()  {
	buf := bytes.NewBufferString("study golang")
	b := buf.Next(3)
	fmt.Println(buf.String())
	fmt.Println(string(b))
}

/**
buffer 基本操作
 */
func main() {
	createBuf()
	writeToBuf()
	bufWriteToFile()
	readBuffer()
	nextBuf()
}
