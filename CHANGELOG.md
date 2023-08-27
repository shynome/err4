# Changelog

## [0.0.7] - 2023-08-27

### Fix

-fix: skip ierr in if stat

## [0.0.6] - 2023-08-27

### Fix

- support `*_test.go` file
- fix: wrong change `args.err4` var

## [0.0.5] - 2023-08-24

### Change

- build tag 更改为 `ierr`
- 生成的文件后缀名修改 `_ierr.go`
- 检测改成探测以 i 开头的 `error` 类型变量了

## [0.0.4] - 2023-06-19

### Change

- 文件开头必须为 `//go:build err4`, 更为严格精确了

## [0.0.3] - 2023-06-19

### Fix

- 添加 nil 检查

## [0.0.2] - 2023-06-17

### Fix

- 修复 `func()(qTry error)` 命名返回不被转换的问题

### Improve

- 错误判断集中在一个 if 里面, 避免注入的错误检查行数过多
