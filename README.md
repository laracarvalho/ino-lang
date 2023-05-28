# INO
INO: Interpreter in G(o). Toy language to learn more about lexers, intepreters, and lang-making with Go, using the [Writing An Interpreter In Go](https://interpreterbook.com/) book as reference.

The WAIIG book creates Monkey lang for its examples, and I'm taking creative and educational liberty of changing the lang and expanding until I reach the limit of my comprehension of the subject.

Syntax example:

```ino
  var five = 5;
  var ten = 10;

  var add = fun(x, y) {
    x + y;
  };

  var result = add(five, ten);
```


* This project is NOT aimed at production value.