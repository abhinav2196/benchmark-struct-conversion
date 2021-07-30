package main

//import (
//"fmt"
//"github.com/jinzhu/copier"
//)
//
//type User struct {
//	Name        string
//	Role        string
//	Age         int32
//	EmployeCode int64 `copier:"EmployeNum"` // specify field name
//
//	// Explicitly ignored in the destination struct.
//	Salary   int
//}
//
//func (user *User) DoubleAge() int32 {
//	return 2 * user.Age
//}
//
//// Tags in the destination Struct provide instructions to copier.Copy to ignore
//// or enforce copying and to panic or return an error if a field was not copied.
//type Employee struct {
//	// Tell copier.Copy to panic if this field is not copied.
//	Name      string `copier:"must"`
//
//	// Tell copier.Copy to return an error if this field is not copied.
//	Age       int32  `copier:"must,nopanic"`
//
//	// Tell copier.Copy to explicitly ignore copying this field.
//	Salary    int    `copier:"-"`
//
//	DoubleAge int32
//	EmployeId int64 `copier:"EmployeNum"` // specify field name
//	SuperRole string
//}
//
//func (employee *Employee) Role(role string) {
//	employee.SuperRole = "Super " + role
//}
//
//func main() {
//	var (
//		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000}
//		users     = []User{{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 100000}, {Name: "jinzhu 2", Age: 30, Role: "Dev", Salary: 60000}}
//		employee  = Employee{Salary: 150000}
//		employees = []Employee{}
//	)
//
//	copier.Copy(&employee, &user)
//
//	fmt.Printf("%#v \n", employee)
//	// Employee{
//	//    Name: "Jinzhu",           // Copy from field
//	//    Age: 18,                  // Copy from field
//	//    Salary:150000,            // Copying explicitly ignored
//	//    DoubleAge: 36,            // Copy from method
//	//    EmployeeId: 0,            // Ignored
//	//    SuperRole: "Super Admin", // Copy to method
//	// }
//
//	// Copy struct to slice
//	copier.Copy(&employees, &user)
//
//	fmt.Printf("%#v \n", employees)
//	// []Employee{
//	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Admin"}
//	// }
//
//	// Copy slice to slice
//	employees = []Employee{}
//	copier.Copy(&employees, &users)
//
//	fmt.Printf("%#v \n", employees)
//	// []Employee{
//	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Admin"},
//	//   {Name: "jinzhu 2", Age: 30, Salary:0, DoubleAge: 60, EmployeId: 0, SuperRole: "Super Dev"},
//	// }
//
//	// Copy map to map
//	map1 := map[int]int{3: 6, 4: 8}
//	map2 := map[int32]int8{}
//	copier.Copy(&map2, map1)
//
//	fmt.Printf("%#v \n", map2)
//	// map[int32]int8{3:6, 4:8}
//}

import (
	"BenchmarkStructConversion/Approaches"
	"BenchmarkStructConversion/Models"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)



func main(){

	wg := sync.WaitGroup{}
	wg.Add(1)

	startTime := time.Now()

	go printStats()
	//s1 := S1{
	//	f1 : "fx",
	//}
	//s2 := S2{
	//	s1: s1,
	//	f2: "xyz",
	//}
 //ip := Input{
 //	name : "abhinav",
 //	s1: s1,
 //	s2 : s2,
 //	f2: "xx",
 //	f1: "yy",
 //}
	//res1D := &response1{
	//	Page:   1,
	//	Fruits: "x",
	//	Z: &response2{X:1,Y:"r"},
	//}

	res2D := Models.Response1{
		Page:   1,
		Fruits: "x",
		Z: &Models.Response2{X:1,Y:"r"},
	}

 for i := 0; i < 10000000; i++ {
	  //var output response1
	  var outputX Models.Response1

	  //copier.Copy(&output,res1D)
	 //err := FillIt(res1D,&output)
	 //err := FillItJsoniter(res1D,&output)
	 //err := FillItEasyJson(res2D,&output)
	 err := Approaches.ConvertWithEasyJson(res2D,&outputX)

	 //manualMapping(res1D,&output)
	if err != nil{
		//fmt.Printf("Recieved an error")
		break
 }

	//fmt.Printf("Output printing= %v\n", output)

 }

	totalTimeTaken := time.Since(startTime)

	fmt.Println("Total time taken",totalTimeTaken)

	//runtime.GC()

 wg.Wait()

}


func manualMappingByTypeCasting(input interface{}, output interface{}) error{

	//output,ok := input.(Input)
	//if !ok {
	//	return errors.New("Type casting failed")
	//}
	return nil
}

func manualMapping(input *response1,output *response1) {
output.Page =  input.Page
output.Fruits = input.Fruits
//var newStruct *response2
//	newStruct.X = input.Z.X
//	newStruct.Y = input.Z.Y
newStruct := &response2{input.Z.X,input.Z.Y}
output.Z = newStruct
//output.Z.Y = input.Z.Y
//output.Z.X = input.Z.X

}

func print_heap_info() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("env:%v, heapsys:%d,heapalloc:%d,heapidel:%d,heapreleased:%d,heapinuse:%d\n",
		os.Getenv("GODEBUG"), m.HeapSys, m.HeapAlloc, m.HeapIdle, m.HeapReleased, m.HeapInuse)
}

func printStats(){
	for {
		print_heap_info()
		fmt.Printf("No of live Goroutines at time%d= %d\n",time.Now(), runtime.NumGoroutine())
		PrintMemUsage()
		time.Sleep(2*time.Second)
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func FillIt(input interface{}, output interface{}) error {
	//fmt.Println("Input recieved body= ",input)
//&Input{"oyoyo","rrr","f",S1{"x"},S2{S1{"f"},"t"}}
	body, err := json.Marshal(input)
	if err != nil {
		return err
	}
	//fmt.Println("Printing marshalled body= ",string(body))
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}
	return nil
}

//func FillItEasyJson(input response1,output *response1) error {
//	//fmt.Println("Input recieved body= ",input)
//	//&Input{"oyoyo","rrr","f",S1{"x"},S2{S1{"f"},"t"}}
//	body, err := input.MarshalJSON()
//	if err != nil {
//		return err
//	}
//	//fmt.Println("Printing marshalled body= ",string(body))
//	err = output.UnmarshalJSON(body)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func FillItJsoniter(input interface{}, output interface{}) error {
	//fmt.Println("Input recieved body= ",input)
	//&Input{"oyoyo","rrr","f",S1{"x"},S2{S1{"f"},"t"}}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//var json = jsoniter.ConfigFastest

	body, err := json.Marshal(input)
	//body, err := json.Marshal(input)
	if err != nil {
		return err
	}
	//fmt.Println("Printing marshalled body= ",string(body))
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}
	return nil
}

//type response1 struct {
//	Page   int `json:"page",copier:"must"`
//	Fruits []string `json:"fruits",copier:"must"`
//	Z *response2 `json:"z",copier:"must"`
//}
//
//type response2 struct {
//	X   int `json:"x",copier:"must"`
//	Y string `json:"y",copier:"must"`
//}

type response1 struct {
	Page   int `copier:"must",json:"page"`
	Fruits string `copier:"must",json:"fruits"`
	Z *response2 `copier:"must",json:"z"`
}



type response2 struct {
	X   int `copier:"must",json:"x"`
	Y string `copier:"must",json:"y"`
}

// MarshalJSON supports json.Marshaler interface
func (v response1) abhinavTest() int {
	fmt.Println("abhinav test")
	return 2
}
