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

### 变量定义要点

- 变量类型写在变量名之后
- 编译器可推测变量类型
- 没有`char`，只有`rune`
- 原生支持复数类型

## 常量的定义

- `const filename = "abc.txt"`
- `const`数值可以作为各种类型使用
- `const a, b = 3, 4`
- `var c int = int(math.Sqrt(a*a + b*b))`，**a和b不用再进行强制类型转换为float**

### 特殊的常量-枚举类型

- 普通枚举类型
- 自增值枚举类型（使用`iota`）

## 条件语句

### if

- `if`的条件里不需要括号
- `if`的条件里可以赋值，但是作用于在`if`语句内

### switch

- `switch`会自动`break`，除非使用**fallthrough**
- `switch`后可以不跟表达式

<2018.09.12>
### `for`

- `for`的条件里不要括号
- `for`的条件里可以省略初始条件，结束条件，递增表达式
- go语言中没有`while`

### 基本语法要点回顾

- `for`和`if`的条件没有括号
- `if`条件里也可以定义变量
- 没有`while`
- `switch`不需要`break`，也可以直接`switch`多个条件