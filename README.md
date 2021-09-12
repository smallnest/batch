# batch

[![License](https://img.shields.io/:license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/smallnest/batch?status.png)](http://godoc.org/github.com/smallnest/batch)  [![travis](https://travis-ci.org/smallnest/batch.svg?branch=master)](https://travis-ci.org/smallnest/batch) [![Go Report Card](https://goreportcard.com/badge/github.com/smallnest/batch)](https://goreportcard.com/report/github.com/smallnest/batch) [![coveralls](https://coveralls.io/repos/smallnest/batch/badge.svg?branch=master&service=github)](https://coveralls.io/github/smallnest/batch?branch=master) 


> 仅支持Go 1.18及以上版本, 因为泛型的原因。

批处理一组元素，或者等待一定的时间。

一旦元素达到设定的阈值，就进行处理，否则调用者就阻塞等待。

如果没有充足的元素需要处理，那么会把既有的元素处理完毕后继续等待.

如果设置了timeout,即使还没有达到设定的阈值，也会进行批处理。

使用这个库既可以避免CPU的无谓空转，也可以有效的进行成批数据的处理。


## 使用方法

```
    in := make(chan int, 100)
	go func() {
        for {
            // 往in中塞数据
            ......
        }
		
	}()

	
	Batch(in, 80, func(items []int) {
		// 处理一批数据
	})
```


