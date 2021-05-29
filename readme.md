Package log10 calculates log base 10 of an integer, fast.

It is inspired by [Daniel Lemire's blog post on this topic](// https://lemire.me/blog/2021/05/28/computing-the-number-of-digits-of-an-integer-quickly).

TODO:

* Add implementations for other bit widths.
* [Teach the compiler some tricks](https://github.com/golang/go/issues/46444) to make this faster yet.
