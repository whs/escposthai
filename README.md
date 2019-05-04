# ESCPOS Thai printing

[![GoDoc](https://godoc.org/github.com/whs/escposthai?status.svg)](https://godoc.org/github.com/whs/escposthai)

Apparently none exists in open source world, and people just print images instead.

However, the ever-so-popular TM-U220 doesn't support github.com/kenshaw/escpos's image rasterizer, so I ended up writing this.

## Usage

```go
p := escpos.New(dst)
// Set character set to 20
p.WriteRaw([]byte{27, 116, 20, 13})
escposthai.PrintThai(p, "John วิญญูรู้ทฤษฎีน้ำแข็ง")
```

It's that easy.

## Caveats

This library assume that the input is a valid formed Thai sequence. Invalid Thai sequence might result in undefined behavior. Invalid sequences including:

- Writing upper letters in reverse order (eg. น -้ -ี)
- Text that starts with vowels
- Text that contains escape characters
- Text that contains unsupported character

## License

[MIT License](LICENSE)
