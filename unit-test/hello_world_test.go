package unit_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -> basic to run unit test
// go test -v  (verbose) -> to see what function that we running for test
// go test -v -run TestNameFunction -> to run specific test
// go test ./... -> running all test inside all packages
// go test -v -bench=. -> running seluruh benchmark di 1 module (unit test di run jg)
// go test -v -run=NotMathUnitTest (nama unit test asal) -bench=. (run all benchmark without unit test) 
// go test -v -run=NotMathUnitTest (nama unit test asal) -bench=BenchA (only run specific benchmark func)
// go test -v -bench=./... -> running all inside all packages

// failing a unit test
// t.Fail() -> menggagalkan unit test tapi tetap proses test selanjutnya dalam 1 func
// t.FailNow() -> menggagalkan unit test saat itu jg dan tidak memproses lebih lanjut test lainnya
// t.Error(args...) -> log print error + kemudian melanjutkan dengan t.Fail()
// t.Fatal(args...) -> sama + kemudia melanjutkan dengan t.FailNow()


func TestHelloWorld(t *testing.T) {
	result := HelloWorld("we")
	if result != "Hello world" {
		t.Error("ini bukan hello world")
	}
	fmt.Println("Test done 1")
}

func TestHelloAl(t *testing.T) {
	result := HelloWorld("al")
	fmt.Println(result)
	if result != "Hello al" {
		panic("Result is not hello al")
	}
	fmt.Println("Test done 2")
}

// assertion
// lib / cara untuk validasi cmn ga ada built in dari golang -> harus nambah library/framework
// testify salah satunya 
// assert klo gagal dia call fail
// require klo gagal dia call fail now

func TestHelloAld(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Unit test cannot work on windows")
	}

	result := HelloWorld("ald")
	assert.Fail(t, "Hello ald", result, "Result must be Hello ald")
}

// testing.M
// required to use TestMain with param testing.M -> hanya bisa dieksekusi once per one package
// go will automatically execute this function each time we running unit test 

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit test")

	m.Run()

	fmt.Println("Sesudah unit test")
}

// sub test
// pembuatan func unit test di dlm unit test
// jika hanya ingin run sub test go test -run TestNamaSubTest/NamaSubTest or go test -run /NamaSubTest

func TestSubTest(t *testing.T) {
	t.Run("Aldrich", func(t *testing.T) {
		result := HelloWorld("Aldrich")
		require.Equal(t, "Hello Aldrich", result)
	})

	t.Run("Dono", func(t *testing.T) {
		result := HelloWorld("Dono")
		require.Equal(t, "Hello Aldrich", result)
	})
}

// table test
// di golang kita bisa membuat test dengan konsep table test
// table test yaitu dimana kita menyediakan data berupa slice yg berisi params dan expected result dari unit test (slice ini kita iterasi menggunakan sub test)

func TestTableSubTest(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Aldrich",
			request: "Aldrich",
			expected: "Hello Aldrich",
		},
		{
			name: "Dono",
			request: "Dono",
			expected: "Hello Dono",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}


// mock -> obj yg sudah kita program dengan ekspetasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yg sudah kita program dari awal
// mock adalah salah satu teknik dalam unit test dmn kita bisa buat object dari suatu object yg sulit ditest
// misal kita punya kode program yg perlu memanggil third party service . ini sulit karena kita harus selalu panggil service tsb dan responsenya jg blm tentu sesuai dengan yg kita inginkan

// Testify mock
// pastikan pembuatan desain kode program kita baik karena jika tidak itu akan sulit untuk di buat mock

// case study -> query ke db 
// kita tidak mau running db 
// buat layer service sebagai business logic dan layer repository as gateway to interact with DB
// agar mudah di test kita wajib menggunakan interface (jgn langsung function ke db)


// benchmark
// mekanisme menghitung perfomance code kita
// caranya golang akan melakukan iterasi kode yg kita panggil berkali2 sampai waktu tertentu (jmlh iterasi dan lamanya ditentukan oleh go sendiri lwt testing.B)
// testing.B mirip dengan testing.T jadi ada func Fail(), FailNow(), Error(), etc yg membedakan func/attribute tambahan untuk benchmarking
// salah satunya attribute N (total iterasi sebuah benchmark)
// cara kerjanya simple kita hanya perlu membuar looping sejumlah N attribute sisanya go yg akan mengurus
// benchmark di golang sm kyk unit test jdi sudah ditentukan nama func nya seperti func BenchmarkA(b *testing.B) dengan params *testing.B (tdk return apapun) bisa digabung dengan unit test



func BenchmarkHelloWorld(b *testing.B) {
	for i :=0; i < b.N; i++ {
		HelloWorld("Aldrich")
	}
}