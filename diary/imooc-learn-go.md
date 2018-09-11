# imooc learn go diary

<2018.09.11>
## 变量定义

### 使用var关键字

- `var a, b, c bool`
- `var s1, s2 string = "hello", "world"`
- 可放在函数内，或直接放在**包内**（看上去像全局变量，但不是全局变量）
- 使用`var()`集中定义变量

### 让编译器自动决定类型

- `var a, b, i, s1, s2 = true, false, 3, "hello", "world"`

### 使用`:=`定义变量

- `a, b, i, s1, s2 := true, false, 3, "hello", "world"`
- 只能在函数内使用

## 内建变量类型

- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
- byte, rune
- float32, float64, complex64, complex128

### 强制类型转换

- 类型转换是强制的（没有隐式转换）
- `var a, b int = 3, 4`
- `var c int = math.Sqrt(a*a + b*b)`    **wrong**
- `var c int = int(math.Sqrt(float64(a*a + b*b)))`  **right**
