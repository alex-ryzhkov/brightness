# brightness.go
A simple utility written in Go to change laptop's backlight level (or any other property which is operated via single byte value in `sysfs`).

## Installation
``` bash
go build brightness.go
sudo chown root brightness
sudo chmod +s brightness
sudo mv brightness /bin
```

## Usage
```
 brightness -inc|-dec VALUE FILE
 ```

## Why?
A brightness utility in my DE had 2 problems:

1. It changed the backlight value with steps that were too big, so it was difficult to fine-tune a proper level of backlight.
2. At the minimal value it would turn off the display completely.

 So I wrote a simple shell script for changing the brightness directly by echoing values into `/sys/class/backlight`. Having a shell script enabled in `/etc/sudoers` (you [can't](http://www.faqs.org/faqs/unix-faq/faq/part4/section-7.html) use `SUID` for files that start with `#!`) triggered my OCR so I rewrote it in a compiled language thus enabling usage of `SUID`.

