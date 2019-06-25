# zeno

**zeno** is a math expression parser capable of:

- Generating a postfix equivalent to given expression
- Create a binary operation tree
- Calculate the result of the expression (without variables)
- Generate LaTeX representation of an operation tree

## Install

Use golang package manager to install zeno
````
$ go get -u github.com/xzebra/zeno
````

## Usage

````go
import "github.com/xzebra/zeno"

int main() {
    expression := "-2.1483*3.14^3+4"
    // Get a postfix representation of the expression
    postfix, _ := zeno.ToPostfix(expression)
    fmt.Println(postfix) // "-2.1483 3.14 3 ^ * 4 +"

    // Get the binary tree representation
    tree, _ := zeno.PostfixToTree(postfix)
    // You can also get it using ToTree
    tree, _ = zeno.ToTree(expression)

    // After that you can get the result (if the expression 
    // doesn't have variables) or its LaTeX representation.
    fmt.Println(tree.Calculate()) // -62.509529055200005
    fmt.Println(tree.LaTeX()) // "neg(2.1483)\cdot3.14^3+4"
}
````

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Remember to update tests according to the changes.

## License
[MIT](https://github.com/xzebra/zeno/blob/master/LICENSE)