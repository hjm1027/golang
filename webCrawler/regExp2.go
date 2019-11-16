//一组测试
package main

import (
    "fmt"
    "regexp"
    "bytes"
)

func main() {
    //这个regexp包是go的标准正则表达式包，其中有很多函数，matchstring就是其中一个。测试模式是否匹配字符串，括号里面是至少有一个a-z之间的字符存在，如果是就返回ture
    match,_:=regexp.MatchString("p([a-z]+)ch","peachp")
    fmt.Println(match)
    //上面我们直接使用字符串匹配的正则表达式，但对于其他匹配任务要使用'compile'来使用一个优化过的正则对象。
    //这里的r声明出来，称为正则结构体，正则结构体有很多方法可以使用。
    r,_:=regexp.Compile("p([a-z]+)ch")//注意这里是至少一个，我可以是peach，也可以是peauch
    //这里的MatchString方法和上面的MatchString函数是一样的。
    fmt.Println(r.MatchString("peach"))
    //go双引号里面，就算有空格也只视为一个字符串。而这里的模板，也就是上面Comoile那行，正则表达式是按模板来，只会筛选符合条件的第一个字段，也就是说，只找p ch的，就算后面连了以一段字符串也不不会输出。然后，只找第一个，这里的punch，peach两个都满足，谁在前面谁输出。
    //这里有一个特殊情况，当两端字符都符合条件，并且中间没有空格连在一起的时候。两段都会输出（是因为前一个p和后一个ch看成了整体）
    fmt.Println(r.FindString("punchpaeach"))

    //下面这个是找索引，找起始索引和结束索引(0和5这里应该是符合前开后闭)
    //同样的，连在一起找第一个p和空格前最后一个ch。
    fmt.Println(r.FindStringIndex("peachpunch"))

    //这个方法会返回匹配的那一个字段，和具体匹配的字段（也就是p和ch中间那段）
    fmt.Println(r.FindStringSubmatch("peachpunch"))

    //这个方法返回符合要求的字段索引和这个字段的索引，前面是前闭后开，后面是两边都取到。
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    //这个方法返回所有匹配的字段（用空格隔开区分字段）前面的都是返回第一个匹配的字段。
    fmt.Println(r.FindAllString("peach punch pinch",-1))

    //这个是一样的，返回所有匹配字段的索引前闭后开，后面返回各自匹配字段的是两边都取。
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch",-1))

    //当然第二个参数也可以不是-1，如果是整数，就可以限制匹配次数（当然是从第一个开始匹配）如果是-1，那么就是匹配全部
    fmt.Println(r.FindAllString("peach punch pinch",2))

    //上面有matchstring的方法可以匹配字符串，也可以用match的方法来匹配字符切片
    fmt.Println(r.Match([]byte("peach")))

    //构建r，也就是正则结构体的时候，我们可以用MustCompile构建，因为这个只返回一个值，而compile返回两个值。
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    //regexp包也可以用来将字符串的一部分替换为其他的值
    //其实就是找p ch之间的字段，然后把整个匹配的字段全部换成第二个参数的字段，而且是全部匹配的都会换掉。这里的<>可有可无
    fmt.Println(r.ReplaceAllString("a peachpeach punch","<hello world>"))

    //这个方法可以将所有匹配的值转换为所需要的值，这里是用了bytes包的ToUpper方法
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in,bytes.ToUpper)
    fmt.Println(string(out))
}


