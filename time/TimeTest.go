package main

import (
	"fmt"
	"time"
)

func main() {
	//时间操作
	//创建时间
	//Golang中时间操作位于 time 包中，常见操作有：
	//当前时间
	nowTime := time.Now()
	fmt.Printf("当前时间为：%T\n", nowTime) //其类型是time.Time
	fmt.Println(nowTime)

	//自定义时间
	customTime := time.Date(2020, 3, 3, 0, 0, 0, 0, time.Local)
	fmt.Println(customTime)

	//时间格式化与解析，Go的时间格式化必须传入Go的生日：Mon Jan 2 15:04:05 -0700 MST 2006
	//stringTime := nowTime.Format("2020年3月3日 15:15:15")//解析的时间是错误的
	//stringTime := nowTime.Format("2006年1月2日 15:04:05")//解析成功，且格式与参数格式一致
	stringTime := nowTime.Format("2006-1-2 15:04:05")
	fmt.Println(stringTime)

	//Go的时间解析
	stringTime = "2020-01-01 01:01:01"
	objTime, err := time.Parse("2006-01-02 15:04:05", stringTime)
	if err != nil {
		fmt.Println("Parse err:", err)
		return
	}
	fmt.Println(objTime)
	//注意：这些方法的参数模板必须与时间一一对应，否则报错！
	//比如上面的stringTime = "2020/01/01 01:01:01" 与参数模板不一致，转换出错

	//获取年月日
	year, month, day := nowTime.Date()
	fmt.Println(year, month, day)
	//获取时分秒
	hour, min, sec := nowTime.Clock()
	fmt.Println(hour, min, sec)

	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(nowTime.Day())
	fmt.Println(nowTime.Hour())

	//今年一共过了多少天
	fmt.Println(nowTime.YearDay())

	//时间戳，计算时间距离1970年1月1日的秒（其他单位）数
	fmt.Println(nowTime.Unix())     //秒
	fmt.Println(nowTime.UnixNano()) //纳秒

	//时间间隔
	fmt.Println(nowTime.Add(time.Second * 10)) //10秒后
	fmt.Println(nowTime.AddDate(1, 0, 0))      //一年后
	//传入负数则往前计算
	//Sub()函数用来计算两个时间的差值
	fmt.Println(nowTime.Add(time.Second * 10).Sub(nowTime)) //10s

	//时间睡眠
	//time.Sleep(time.Second * 3) //程序睡眠3秒钟

	//时间中的通道操作（定时器）
	//标准库中的Timer可以让用户自定义一个定时器，在用对select处理多个channel的超时、单channel读写的超时等情形时很方便：
	timer := time.NewTimer(time.Second * 3)
	//ch := timer.C     //timer内部包含一个通道
	//fmt.Println(<-ch) //3秒后，通道内有了数据，可以取出
	//配合协程
	go func() {
		<-timer.C
		fmt.Println("timer结束")
	}()
	time.Sleep(time.Second * 5) //睡眠5秒防止主线程退出
	flag := timer.Stop()        //取消定时器
	fmt.Println(flag)

	//time.After()函数的使用：
	ch := time.After(time.Second * 3) //底层其实是new Timer(d).C
	newTime := <-ch                   //阻塞3秒
	fmt.Println(newTime)
}
