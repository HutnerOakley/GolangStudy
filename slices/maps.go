package slices

import "fmt"

/*
Go 的 `map[K]V`，是**哈希表**（hash table），key 唯一，键值对集合。
类比 Python 的 `dict`，但底层、key 限制、传参行为有不少差异。

类型声明：`map[KeyType]ValueType`

*/

func modify(m map[string]int) {
	m["test"] = 999 // 修改map内容,外部生效
}

func TestMap() {

	// 方式 1：字面量初始化
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m)

	//方式 2：直接声明
	//`make(map[K]V, 初始容量hint)`，第二个参数是**预估容量**，
	//不是固定上限，只是减少扩容。
	var test1 map[string]string // nil map，没有分配哈希表底层空间
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	test1 = make(map[string]string, 10)
	test1["one"] = "1"
	test1["two"] = "2"
	test1["three"] = "3"
	fmt.Println(test1)

	//方式 3：make创建（推荐，空 map）
	test2 := make(map[string]string)
	test2["four"] = "4"
	test2["five"] = "5"
	test2["six"] = "6"
	fmt.Println(test2)

	// 增、删、改、查
	mps := make(map[string]int)

	//新增 / 修改：同一个语法，key存在就是更新，不存在新增
	mps["zhangsan"] = 10

	// 查询
	age := mps["zhangsan"]
	fmt.Println(age)

	//【重要】map读取会返回两个值：value, exists
	//如果只写单返回值：key 不存在返回**value 类型的零值**。
	//比如 `map[string]int`，找不到 key 返回`0`；很容易误判，业务代码推荐**用 ok 判断是否存在**
	age, ok := m["list"]
	if ok {
		fmt.Println("存在，age=", age)
	} else {
		fmt.Println("key不存在")
	}

	// 删除
	delete(mps, "zhangsan")

	// 遍历map
	// **Go map 遍历是无序的！**
	//每次 range 输出顺序都不一样，Go 故意随机化遍历顺序，防止程序员依赖顺序。
	//如果需要有序，把 key 放到 slice 里排序，再循环 slice 取 map
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// map 本身是一个**hmap 结构体的指针包装**。
	//函数传递 map，依然是**值传递**，拷贝这个指针。
	//所以：**函数内部修改 map 里的键值对，外面原 map 能看到变化！**
	m3 := make(map[string]int)
	modify(m3)
	fmt.Println(m3) //map[test:999]

}
