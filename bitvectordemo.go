package main
import(
	"fmt"
	"github.com/dropbox/godropbox/container/bitvector"

)


func main() {

	bitSlicer:=[]byte("Welcome to Bit Vecor")
	data:=bitvector.NewBitVector(bitSlicer,8)
	fmt.Println(data)
	data=bitvector.NewBitVector(bitSlicer,16)
	fmt.Println(data)
	data=bitvector.NewBitVector([]byte{0xF0, 0x0F}, 12)
	fmt.Println(data)
	var bitData int64 = 1<<32
	fmt.Println(bitData)
	var testData byte=123
	result:=append(bitSlicer, testData)
	fmt.Println(result)

}
