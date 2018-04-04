# Cracking the Coding Interview in Go

I bought [Cracking the Coding Interview, 6th Ed.](http://www.crackingthecodinginterview.com/)
as a refresher and study guide. The book is well worth the price for its succinct organization
and strategies and hints toward problem solving.

Seeing someone else's implementation is helpful can be helpful, especially when you get stumped,
but to truly benefit from the book, it is absolutely vital to write code yourself.

There is an existing [repo](https://github.com/careercup/ctci), but I have not looked at it. I
prefer to attempt to solve the problems on my own using the hints and strategies presented in
the book. I have definitely needed to refer to the book's solutions, however, which are written
in Java.

I have implemented the solutions as tests using the Go [testing](https://golang.org/pkg/testing/)
package. I start each implementation by writing the test cases that I would like to pass, and
then work toward writing the solution.

Each test can be run like this:

    $ go test chapter-01/urlify_test.go


I've only just begun, so this is definitely a work in progress.
