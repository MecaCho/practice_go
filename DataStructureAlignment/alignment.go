package DataStructureAlignment

import (
	"fmt"
	"unsafe"
)

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

type Part2 struct {
	e byte
	c int8
	a bool
	b int32
	d int64
}

// unsafe.Alignof 来返回相应类型的对齐系数
func CompTwoStructAlign() {
	part1 := Part1{}
	part2 := Part2{}

	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	fmt.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
}

type S1 struct {
	a     int
	ID    int64
	Name  string
	Alias string
}

func StructAlign() {
	fmt.Printf("bool align: %d\n", unsafe.Alignof(bool(true)))
	fmt.Printf("int32 align: %d\n", unsafe.Alignof(int32(0)))
	fmt.Printf("int8 align: %d\n", unsafe.Alignof(int8(0)))
	fmt.Printf("int64 align: %d\n", unsafe.Alignof(int64(0)))
	fmt.Printf("byte align: %d\n", unsafe.Alignof(byte(0)))
	fmt.Printf("string align: %d\n", unsafe.Alignof("EDDYCJY"))
	fmt.Printf("map align: %d\n", unsafe.Alignof(map[string]string{}))

	fmt.Printf("struct align: %d\n", unsafe.Alignof(S1{1, 1, "test_name", "alias name"}))
}

// 成员对齐
// 第一个成员 e
// 类型为 byte
// 大小/对齐值为 1 字节
// 初始地址，偏移量为 0。占用了第 1 位
// 第二个成员 c
// 类型为 int8
// 大小/对齐值为 1 字节
// 根据规则1，其偏移量必须为 1 的整数倍。当前偏移量为 2。不需要额外对齐
// 第三个成员 a
// 类型为 bool
// 大小/对齐值为 1 字节
// 根据规则1，其偏移量必须为 1 的整数倍。当前偏移量为 3。不需要额外对齐
// 第四个成员 b
// 类型为 int32
// 大小/对齐值为 4 字节
// 根据规则1，其偏移量必须为 4 的整数倍。确定偏移量为 4，因此第 3 位为 Padding。而当前数值从第 4 位开始填充，到第 8 位。如下：ecax|bbbb
// 第五个成员 d
// 类型为 int64
// 大小/对齐值为 8 字节
// 根据规则1，其偏移量必须为 8 的整数倍。当前偏移量为 8。不需要额外对齐，从 9-16 位填充 8 个字节。如下：ecax|bbbb|dddd|dddd
// 整体对齐
// 符合规则 2，不需要额外对齐
//
// 结果
// Part2 内存布局：ecax|bbbb|dddd|dddd
//
// 1.5. 总结
// 通过对比 Part1 和 Part2 的内存布局，你会发现两者有很大的不同。如下：
//
// Part1：axxx|bbbb|cxxx|xxxx|dddd|dddd|exxx|xxxx
//
// Part2：ecax|bbbb|dddd|dddd
