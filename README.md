### A implemention for golang Error-Handling

for this proposal [proposal: Go 2: Error-Handling Paradigm with !err Grammar Sugar](https://github.com/golang/go/issues/62253)

和提案中不同的是, 这里的 ierr 需要提前声明为 error 类型

### how to use

#### 安装

```
go install github.com/shynome/err4/cmd/err4gen@v0.0.9
```

#### 创建文件

```go
// err4.go
package main

//go:generate err4gen .
```

#### 编写代码

```go
//go:build ierr
// main.go

package main

import (
	"errors"
)

func main() {
	var ierr error
	var ierr2 error
	_, _ = ierr, ierr2

	ierr = errors.New("e")

	var e1, e2 error
	_, ierr, ierr2 = 7, e1, e2
}

func fn() (ierr error) {
	ierr = errors.New("e")
	return
}

func fn2() (ierr error) {
	ierr = errors.New("e")
	return
}

```

#### 生成代码

运行 `go generate .` 后会生成 `main_ierr.go` 文件

```go
//go:build !ierr

// Code generated by github.com/shynome/err4 DO NOT EDIT

package main

import (
	"errors"
)

func main() {
	var ierr error
	var ierr2 error
	_, _ = ierr, ierr2

	ierr = errors.New("e")
	if ierr != nil {
		return
	}

	var e1, e2 error
	_, ierr, ierr2 = 7, e1, e2
	if ierr != nil || ierr2 != nil {
		return
	}
}

func fn() (ierr error) {
	ierr = errors.New("e")
	if ierr != nil {
		return
	}
	return
}

func fn2() (ierr error) {
	ierr = errors.New("e")
	if ierr != nil {
		return
	}
	return
}

```

### For VSCode User

you need add `"go.buildTags": "ierr"` to `settings.json`

```json
{
  "go.buildTags": "ierr"
}
```
