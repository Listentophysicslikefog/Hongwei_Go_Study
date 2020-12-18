package main
import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

/*   map最详细解释 ： http://www.coder55.com/article/7759 
var m map[string]int
    // https://blog.golang.org/go-maps-in-action
    // Map types are reference types, like pointers or slices,  and so the value of m above is nil;
    // it doesn't point to an initialized map. A nil map behaves like an empty map when reading, 
    // but attempts to write to a nil map will cause a runtime panic; don't do that. 
    // To initialize a map, use the built in make function
    // Map类型 是引用类型，因此 m 的值为 nil
    fmt.Printf("type of un-initialized map reference %vn", reflect.TypeOf(m))
    fmt.Printf("type of the pointer of map reference %vn", reflect.TypeOf(&m))
    fmt.Printf("The un-initialized map reference %pn", m)
    fmt.Printf("The pointer of map reference %pn", &m)
    // The make function allocates and initializes a hash map data structure 
    // and returns a map value that points to it. The specifics of that data structure
    // are an implementation detail of the runtime and are not specified by the language itself.
    // make函数将会分配并初始化一个底层hash map结构，然后返回一个 map 值，该值指向底层的hash map结构
    m = make(map[string]int)
    fmt.Printf("The initialized map reference %pn", m)
    fmt.Printf("type of initialized map reference %vn", reflect.TypeOf(m))
    fmt.Printf("The pointer of map reference %pn", &m)
*/



func main(){

	//map ,map类型的变量默认的初始值为nil，map是引用类型所以使用的时候需要先使用make函数来分配内存
	//map初始化方法make(map[关键字类型]值类型,容量)
	var m1 map[string]int
    m1= make(map[string]int,9)  //一定要make申请空间，map初始化
	m1["hello"]=666
	m1["he"]=789
	fmt.Println(m1)

	//判断是否有该关键字对应的值
	value,ok:=m1["hello"]
	if !ok{
fmt.Print("没有该关键字对应的值")
	}else{
		fmt.Println(value)
	}
	

	//map的遍历
	for key,val:=range m1{
		fmt.Println(key,val)
	}


	//map的删除

	delete(m1,"he")   //根据关键字删除
	fmt.Println(m1)

	delete(m1,"hel")  //删除没有的元素那么就不会操作




	//对map进行排序后输出
	rand.Seed(time.Now().UnixNano())  //初始化随机数种子
    var scormap=make(map[string]int,200)
    for i:=0;i<100;i++{
		key:=fmt.Sprintf("stu%02d",i) //生成stu开头的字符串
		value:=rand.Intn(100)  //生成0~99的随机数
		scormap[key]=value
	}

	//取出map中的所有key存入切片keys
	var keys=make([]string,0,200)
	for key:=range scormap{
     keys=append(keys,key)
	}

	//对切片进行排序
	sort.Strings(keys)

	//按照排好序的key遍历map
	for _,key:=range keys{
		fmt.Println(key,scormap[key])
	}

	//如果去获取一个map里不存在的key对应的value值，那么就会返回对应类型的0值


	//元素类型为map的切片
var s1=make([]map[int]string,9,10)  //这里是初始化切片，没有初始化map所以map里没有元素

s1[0]=make(map[int]string,1)  //需要对map初始化，这里初始化了一个，不然后面会报错
s1[0][2]="A"  //表示切片第一个元素，并且该元素还是map，对该map赋值
fmt.Println(s1)

//值为切片类型的map
var m11=make(map[string][]int,10)
m11["北京"]=[]int{10,20,30}
fmt.Println(m11)
}
