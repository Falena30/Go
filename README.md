# Go
this repository about my studying Golang

progress from day 1
Installing Golang and make first Golang file
 -Print
 -Variable String
 -variable type
 -comment
 -shorthand variable
 -operand math
 -constant
 -convert string to int or int to string
 
 day 2
using github to store my progess
func main in golang only can use in one file so remember
 -function and return
 -function operand math
 -function using parameter
 -function multipel return
 -loop
 -loop penjumlahan deret
 -if else
- else if
 -function with if else
 -switch case and default
 -variabel room
 -pointer
 -array 
 -array multidimensi
 -slice 
-slice dengan make
 -copy and append
 -key value dengan map
 -check dan delete key  value
 -struct
 -import local file
 -using method
 -interface adlah sekumpulan method kosong yang ntninya digunakan untuk mempermudah dalam pemanggilannya biasanya dipasangkan dengan struct

Hari 3
day off karena listrik mati dan semangat hilang

Hari 4
 -http postman
 -first request dengan postman

-----------------------------------------------
new lesson

Hari 1
- untuk melakukan pencaarian dapat menggunakan string.Contain dimana merupakan bagian dari package string dan fungsinya untuk mengetahui mulai dari satu kata atau kalimat, dan keluarnya berupa boolean true atau false

- pointer nilai defaultnya adalah nil, dan juga tidak bisa asal diisi karena sifatnya yang refrence jadi jika ingin disi harus menggunakan &untuk merubah dari refrence ke isinya, dan untuk mencetak pointer harus memakaiin * jika tidak yang akan muncul adalah refrencenya, jika pointer diisi dengan variabel lain da variable lain itu diganti isinya makan nilai pointer akan berubah juga

-struct boleh diisi salah satunya tapi tidak akan bisa dicetak keduanya secara bersamaan, embed struct digunakan untuk membedakan dan mempermudah struct, jika embed struct memiliki kesamaan variabel antara parent dan childnya maka untuk pemanggilan harus lebih spesifik yaitu dengan menambahkan nama parentya terlebih dahulu, unttuk menggunsksn slice atau array cukup menambahkan [] pada saat deklarasinya, jika dia bertipe data var maka nilainya harus disimpan terlebih dahulu baru bisa ditamplkan

-perbedaan method dan func adalah cara deklarasinya, yaitu di sela2 func dan nama functionya

-untuk memanggil private properti pada file lain harus masukkan kedalam fungsi public jika tidak pasti akan terjadi error, atau membuatnya menjadi public yaitu mengawali namanya menjadi huruf besar

CATATAN PENTING
-jika ada beberapa file di folder main dan ingin memasukkan isinya ke main, bisa langsung di pakai akan tetapi jika ingin menjalankan harus menjalankan kedua filenya contoh go run main.go antoher.go

-init() adalah fungsi khusus dimana fungsi ini akan dijalankan terlebih dahulu bahkan lebih dulu dari main()

-Goroutine adalah sebuah fitur pada golang, dimana fitur ini berfungsi untuk membuat proses menjadi lebih baik atau menjadi lebih cepat sehingga tidak berjalan bersamaan, misal ada dua varibale yang dicetak dan salah satunya menggunakan Goroutine maka keduanya tidak akan berjalan bersamaan, untuk menggunakan Goroutine pertama harus masukkan package runtime dan menggunakan 'go' sebelum method yang dipanggil

