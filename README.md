# golang-meetup5 

This repo exclusively created for the 5th Go meetup in Israel. 
The repo has a 2 scripts 1 in golang and the second in python testing binary trees search which its input source
is a random 256 characters file.

### **binary_tree.go**
Usage:

```go run binary_tree.go ```

A python code with a preorder traversal binary tree search.

**readFile**

A function that input a file with characters and return a string
**InsertElement**

This function construct the binary tree by an input of a single integer.
In order to search character the input should be modified to ascii and from ascii to integer
Usage 

```tree.insertElement(int, int) ```

**searchNode**

This boolean function searches for a single integer in a binary tree, retrun true or false.
usage ```tree.searchNode(int value)
### **binary_tree.py**
Usage:

```./binary_tree.go ```

A python code with a preorder traversal binary tree search.

** BinaryTree(File path) **

Creating new instance of binary tree object input file path

Usage:

``` foo = BinaryTree(path='bar.txt') ```

** searchValue('char') **

This function doing the search of a single character within the binary tree object

Usage:

``` foo.searchValue('s') ```

### createRandomChars.sh ###

A bash script that output a random 256 characters, recommended to redirect the output to a file

Usage:

``` ./createRandomChars.sh >> file.txt ```

### charsFile.txt ###

This file created by the **createRandomChars.sh** a file that contains random characters.
There is a single Hebrew letter that added 'ד' which entered for a complex binary search.





