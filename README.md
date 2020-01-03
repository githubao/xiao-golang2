# xiao-go-lang
&lt;go programming language> src code

1. 字符串格式化类型
- %d %x %o %b: 十进制整数，十六进制，八进制，二进制
- %f，%g, %e: 浮点数，自动浮点数，科学计数法
- %t: 布尔值
- %v: 内置格式的任何值
- %T: 任何值的类型
- %q: 带双引号的字符串或者带单引号的字符

2. map可以初始化为对应类型的初始值，而不需要类似python的defaultdict类似的东东。

3. 跟C语言不同，go只能使用i++，而且i++在go语言中等价于 i += 1

4. 跟C语言不同，switch,case语句不会默认的fallthrough

5. Go语言的关键字

   | 关键词  | value                               |
   | ------ | ----------------------------------- |
   | 1      | break default func interface select |
   | 2      | case defer go map struct            |
   | 3      | chan else goto package switch       |
   | 4      | const fallthrough if range type     |
   | 5      | continue for import return var      |

6. Go语言的预定义常量(&类型&函数)

   | 预定义常量(&类型&函数)    |                                           |
   | --------------------- | ----------------------------------------- |
   | 常量                   | true false iota nil                       |
   | 类型                   | int int8 int16 int32 int64                |
   | 类型                   | uint uint8 uint16 uint32 uint64 uintptr   |
   | 类型                   | float32 float64 complex128 complex64      |
   | 类型                   | bool byte rune string error               |
   | 函数                   | make len cap new append copy close delete |
   | 函数                   | complex real imag                         |
   | 函数                   | panic recover                             |

7. go语言中有零值保障机制，未初始化的变量会自动赋值为0

8. 调用new函数，返回的是对象的指针， var p *int = new(int)

9. 与c语言不同，函数里面返回变量的指针是安全的

10. 两个变量的类型不携带任何信息且是零值，例如struct{} 或者 []int，那么目前的实现里面，他们有相同的地址。针对new函数来讲。

11. 把函数里面的局部变量的地址赋值给全局变量，称作局部变量的逃逸，尽量避免这种情况，因为会阻碍垃圾回收。

12. map[], <-ch，x.(T),返回一个布尔型的变量，通常命名为ok，表示右边的操作有没有成功。

13. 类型通常会声明String()方法，表示该类型的人类可能形式，但是不要在里面使用fmt.Sprint等一系列相关的方法都会造成递归调用导致程序堆栈满异常。

14. Go语言目录：语言基础，表达式，函数，数据，方法，接口，并发，包，进阶

15. [极客学院Golang学习资料](http://wiki.jikexueyuan.com/project/the-go-study-notes-fourth-edition/language.html)

16. go lang keywords:
- break default func interface select
- case defer go map struct
- chan else goto package switch
- const fallthrough if range type
- continue for import return var

17. what's defer chan fallthrough?
- defer: finally run code
- chan: FIFO queue
- fallthrough: switch case fallthrough

18. predefined identifier:
- append bool type cap close complex complex64 complex128 uint16
- copy false float32 float64 imag int int8 int16 unit32
- int32 int64 iota len make new nil panic uint64
- print println real recover string true uint uint8 uintptr

19. type: rune(unicode)=int32 chan unitptr

20. 结构体，接口，函数，大小写区分可见性。大写可见

21. 常量结构体里面，如果后面一行没有初始化，那么就会使用上面一行的运算结果值

22. go语言中，数组是值类型，结构是值类型，切片是引用类型

23. copy(dest,src)

24. 当赋值类型到接口时，会发生拷贝，所以对原类型的修改，对接口无效

25. 使用resp作为response的缩写

26. 目录使用中划线命名，包名去掉中划线

27. 函数命名应该是做什么，而不是怎么做

28. 函数入参不多于5个，出参不多于2个

29. chan的参数，根据最小可见原则，可以设置只读或者只写

30. 错误分组定义，没有或者只有一个失败的原因时不使用error

31. error或者log信息不要以大写开头，不要以点号结尾

32. 每个协程里面需要defer

33. fmt.Sprintf相关的函数内部都用到了反射，追求性能的话考虑使用strconv.FormatInt 或者 strconv.Itoa

34. iota的默认值是0

35. 注意在函数中的简式变量声明引起的额外后果，弄清楚是使用的原来的变量还是声明了一个新的变量

36. bytes.Equal认为空slice和长度为零的相同，但是reflect.DeepEqual认为两种情况不同

37. 如果用type重新定义了一个类型的别名，是不会继承原来类型对应的方法的

38. 从for switch 和 for select中跳出，跳出的是switch或者select并不是for

39. 在循环中使用defer，需要用一个函数把逻辑代码块包起来

40. map的value元素是不能够被取地址的，如果想更新它的值只能重新赋值到另外一个变量

41. runtime.GOMAXPROCS()表示os的线程数，默认为1，一般设置为cpu的个数，最大值可设置为256

42. 多个go协程不能保证对同一个变量操作的先后顺序，如果需要顺序的保证应该使用chan

43. 使用runtime.Gosched()使得调度器让出当前协程的资源占用

44. 元素在通道中传递是浅层拷贝

45. waitgroup在函数之间传递的时候需要传递指针

46. goarch(386, amd64, arm), goos(darwin, freebsd, linux, windows)

47. 对于java：堆上的变量在外面能找到是new出来的(需要被gc)，函数内部的变量在栈上分配(退出作用域自动被回收)

48. chan由写的一方关闭

49. 分批处理字节流上传文件

50. sync.Pool包提供的连接池复用机制

51. 目前最火的20种自由职业技能：
区块链、TensorFlow、亚马逊dynamoDB系统、配音、字幕
美术设计、内容策划、计算机视觉、微软Power BI、增强现实
聊天机器人、React Native、流量采购、Go语言开发、信息安全
Scala开发、Instagram API、Adobe Premiere、机器学习、AngularJS开发

52. 在反射中使用断言的时候，一定要使用指针进行断言

53. too many open files, 首先查看是否有连接未关闭，然后查看操作系统限制，用户态限制，进程限制三个