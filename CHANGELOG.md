# Changelog

## [0.0.2] - 2023-06-17

### Fix

- 修复 `func()(qTry error)` 命名返回不被转换的问题

### Improve

- 错误判断集中在一个 if 里面, 避免注入的错误检查行数过多