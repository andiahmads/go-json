package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

/*
golang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu dengan menggunakan
function json.Marshal(interface{})
karena parameternya interface{}, kita bisa masukkan tipe data apapapun kedalam function Marshal

*/

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestLogJson(t *testing.T) {
	LogJson("andi")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"foo", "bar"})

}

/*
jika mengikuti standar dari json.org data json bentuknya adalah object dan array
pada golang object direpresentasikan dengan tipe data struct
dimana tap attribute di JSON Object merupakan attribut struct.
*/

type Customer struct {
	FirstName string
	MiddName  string
	LastName  string
	Hobbies   []string
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "andi",
		MiddName:  "ahmad",
		LastName:  "saputra",
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

/* Decode JSON
untuk melakukan konversi dari JSON ke tipe data DECODE, kita bisa menggunakan functin json.Unmarshal(byte[], interface{})
dimana byte[] adalah data dari jsonnya, sedangkan interface{} adalah tempat menyimpan hasil konversinya, bisa berupa pointer.
*/

func TestDecodeJson(t *testing.T) {
	jsonString := `{"fist_name":"andi","mid_name":"ahmad","last_name":"saputra"}`
	jsonBytes := []byte(jsonString)

	//ambil struct
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

/* JSON ARRAY
Pada golang json array direpresentasikan dalam bentuk slice
konversi dari json atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice
*/
func TestJsonArray(t *testing.T) {
	new_customer := Customer{
		FirstName: "andi",
		MiddName:  "ahmad",
		LastName:  "saputra",
		Hobbies:   []string{"a", "b", "c"},
	}
	bytes, _ := json.Marshal(new_customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"fist_name":"andi","mid_name":"ahmad","last_name":"saputra","Hobbies":["a","b","c"]}`
	byteJson := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(byteJson, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONARRAYComplex(t *testing.T) {
	customerWithAddress := &Customer{
		FirstName: "Endi",
		MiddName:  "ahmad",
		LastName:  "Saputra",
		Addresses: []Address{
			{
				Street:     "jl.hidupbaru",
				Country:    "indonesia",
				PostalCode: "123",
			},
			{
				Street:     "jl.hidupbaru 2",
				Country:    "brazil",
				PostalCode: "1234",
			},
		},
	}
	bytes, _ := json.Marshal(customerWithAddress)
	fmt.Println(string(bytes))
}

func TestDecodeJsonArrayComplex(t *testing.T) {
	jsonString := `{"fist_name":"Endi","mid_name":"ahmad","last_name":"Saputra","Hobbies":null,"address":[{"street":"jl.hidupbaru","country":"indonesia","postal_code":"123"},{"street":"jl.hidupbaru 2","country":"brazil","postal_code":"1234"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer.Addresses)
}

/* DECODE JSON ARRAY
selain menggunakan array pada attribute di object
kita juga bisa melakukan encode atau decode lansung json arraynya
Encode dan decode JSON array bisa menggunakan tipe data slice
*/

func TestDecodeJsonArrayUsingSlices(t *testing.T) {
	jsonArray := `[{"Street":"jl.hidupbaru","Country":"indonesia","PostalCode":"123"}]`
	jsonbytes := []byte(jsonArray)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonbytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)

}

func TestEncodeJsonArray(t *testing.T) {
	addresses := []Address{
		{
			Street:     "jl.hidupbaru",
			Country:    "indonesia",
			PostalCode: "123",
		},
		{
			Street:     "jl.hidupbaru 2",
			Country:    "brazil",
			PostalCode: "1234",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}

/* MAP
saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSON nya dynamic
atributnya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap.
pada kasus seperti ini menggunakan struct bukan opsi yg tepat.
untuk kasus seperti ini kita bisa menggunakan tipe data map[string]interface{}
secara otomatis atribut akan menjadi key di map. dan value menjadi value dimap.
namun karena value berupa interface, maka kita harus melakukan konversi secara manual jika ingin mengambil value.
dan tipe data map tidak  mendukung json tag.
*/

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"2112","name":"macbook pro m1","price":20000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)
	fmt.Println(result)

}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    12,
		"name":  "apple mac book",
		"price": 2000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}

/*
STREAM DECODER
terdapat suatu case yg mana data json tersebut barasal dari i.Reader(File,Network,Request.Body)
nah pada case seperti ini kita bisa saja meyimpan data io.Reader kedalam sebuah variable lalu konversi ke json.
tetapi ada cara yg lebih sederhana, karena package json memiliki fitur untuk membaca data dari stream.

untuk membuat json.Descoder kita bisa menggunakan function json.NewDecoder(reader)
selanjutnya membaca isi input reader dan konversi lansung dengan menggunakan Decode(interface{})

*/
func TestStreamDecoder(t *testing.T) {
	file, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	customer := &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
}

/*
STREAM ENCODER
package json juga mendukung membuat Encoder yang bisa digunakan untuk menulis lansung JSON nya ke io.Writer.
dengan begitu kita tidak perlu lagi menyimpan JSON datanya terlebih dahulu kedalam variabel string atau byte.
kita bisa tulis lansung ke io.Writer

untuk membuat encoder kita bisa menggunakan function json.NewEncoder(writer)
dan untuk menulis data sebagai json lansung ke writer, kita bisa gunakan function Encode(interface{})
*/

func TestStreamEncoder(t *testing.T) {
	write, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(write)

	customer := &Customer{
		FirstName: "Andi",
		MiddName:  "ahmad",
		LastName:  "saputra",
	}
	encoder.Encode(customer)
}
