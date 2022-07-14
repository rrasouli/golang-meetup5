#!/usr/local/bin/python3.7
import sys
sys.setrecursionlimit(30000)

class BinaryTreeNode:
    def __init__(self, data):
        self.data = data
        self.leftChild = None
        self.rightChild = None


class BinaryTree(object):
    def __init__(self, path):
        self.path = path
        self.arr = list()
        self.read_file_and_create_arr()

    def read_file_and_create_arr(self):
        f = open(self.path, "r")
        input = f.read()
        ls_input = list(input)
        for ch in ls_input:
            self.arr.append(ord(ch))
        root = self.insert(None, self.arr[0])
        self.root1 = root
        for x in self.arr[1:]:
            self.insert(root, x)
        self.preorder(root)

    def insert(self, root, newValue):
        # if binary search tree is empty, make a new node and declare it as root
        if root is None:
            root = BinaryTreeNode(newValue)
            return root
        # binary search tree is not empty, so we will insert it into the tree
        # if newValue is less than value of data in root, add it to left subtree and proceed recursively
        if newValue < root.data:
            root.leftChild = self.insert(root.leftChild, newValue)
        else:
            # if newValue is greater than value of data in root, add it to right subtree and proceed recursively
            root.rightChild = self.insert(root.rightChild, newValue)
        return root

    def preorder(self, root):
        # if root is None return
        if root == None:
            return
        # traverse root
        # print(root.data)
        # traverse left subtree
        self.preorder(root.leftChild)
        # traverse right subtree
        self.preorder(root.rightChild)

    def search_value(self, value):
        value = ord(value)
        node_obj = self.root1
        while node_obj != None:
            if node_obj.data < value:
                node_obj = node_obj.rightChild
            elif node_obj.data > value:
                node_obj = node_obj.leftChild
            else:
                print(f"we found '{chr(value)}'")
                return
        print(f"Not found '{chr(value)}'")


a = BinaryTree(path="charsFile.txt")
a.search_value(value="ד")