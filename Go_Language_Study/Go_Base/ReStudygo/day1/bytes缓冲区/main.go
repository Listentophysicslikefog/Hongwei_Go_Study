package main

import (
    "bytes"
    "fmt"
)

func main() {
    byteSlice := make([]byte, 20) 
    byteSlice[0] = 1                                  // 将缓冲区第一个字节置1
    byteBuffer := bytes.NewBuffer(byteSlice)          // 创建20字节缓冲区 len = 20 off = 0
    c, _ := byteBuffer.ReadByte()                     // off+=1
    fmt.Printf("len:%d, c=%d\n", byteBuffer.Len(), c) // len = 20 off =1   打印c=1
    byteBuffer.Reset()                                // len = 0 off = 0
    fmt.Printf("len:%d\n", byteBuffer.Len())          // 打印len=0
    byteBuffer.Write([]byte("hello byte buffer"))     // 写缓冲区  len+=17
    fmt.Printf("len:%d\n", byteBuffer.Len())          // 打印len=17
    byteBuffer.Next(4)                                // 跳过4个字节 off+=4
    c, _ = byteBuffer.ReadByte()                      // 读第5个字节 off+=1
    fmt.Printf("第5个字节:%d\n", c)                    // 打印:111(对应字母o)    len=17 off=5
    byteBuffer.Truncate(3)                            // 将未字节数置为3        len=off+3=8   off=5
    fmt.Printf("len:%d\n", byteBuffer.Len())          // 打印len=3为未读字节数  上面len=8是底层切片长度
    byteBuffer.WriteByte(96)                          // len+=1=9 将y改成A
    byteBuffer.Next(3)                                // len=9 off+=3=8
    c, _ = byteBuffer.ReadByte()                      // off+=1=9    c=96
    fmt.Printf("第9个字节:%d\n", c)                    // 打印:96
}