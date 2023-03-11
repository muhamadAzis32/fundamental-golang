// 1
i := 21
j := true
russia := 'Я'
k := 123.456

fmt.Printf("Nilai integer i: %d\n", i)
fmt.Printf("Tipe data dari variabel i: %T\n", i)
fmt.Printf("Nilai boolean j: %t\n", j)
fmt.Printf("Tanda persen: %%\n")
fmt.Printf("Unicode Russia: %c\n", russia)
fmt.Printf("Nilai base 10: %d\n", i)
fmt.Printf("Nilai base 8: %o\n", i)
fmt.Printf("Nilai base 16 (lowercase): %x\n", i)
fmt.Printf("Nilai base 16 (uppercase): %X\n", i)
fmt.Printf("Unicode karakter Я: %U\n", russia)
fmt.Printf("Float: %f\n", k)
fmt.Printf("Float scientific: %e\n", k)

// 2
var i int = 21
var j bool = true
var k float64 = 123.456

fmt.Printf("%v \n", i)
fmt.Printf("%T \n", i)
fmt.Printf("%% \n")
fmt.Printf("%t \n \n", j)
fmt.Printf("%b \n", i)
fmt.Printf("%c \n", 10000101111)
fmt.Printf("%d \n", i)
fmt.Printf("%o \n", i)
fmt.Printf("%x \n", 15)
fmt.Printf("%X \n", 15)
fmt.Printf("%U \n \n", 'Я')
fmt.Printf("%f \n", k)
fmt.Printf("%e \n", k)

https: //prnt.sc/k5buvEoB3YSC


// Rusia
fmt.Printf("%c\n", '\u042F')