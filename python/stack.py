#!/usr/bin/env python3
# coding=utf-8

#########################################################################
#   Copyright (C) 2019 All rights reserved.
#   
#   FILE: stack.py
#   AUTHOR: Max Xu
#   MAIL: xuhuan@live.cn
#   DATE: 2019.04.22 16:45:34
#
#########################################################################


import os
import sys

class Node():
    def __init__(self, data, next=None):
        self._data = data
        self._next = next


class LinkedStack():
    def __init__(self):
        self._top: Node = None

    def push(self, data):
        print("Push", data)
        new_top = Node(data)
        new_top._next = self._top
        self._top = new_top
    
    def pop(self) -> object:
        if self._top:
            data = self._top._data
            self._top = self._top._next
            print("Pop:", data)
            return data
        else:
            return None

    def __repr__(self) -> str:
        s = ""
        top = self._top
        while top != None:
            s += "->" + str(top._data)
            print(s)
            top = top._next
        return s 


def main(args):
    ls = LinkedStack()
    ls.push(12)
    ls.push(14)
    ls.push(16)
    print(ls)

    print(ls.pop())
    print(ls.pop())
    print(ls.pop())
    print(ls.pop())


if __name__ == "__main__":
    main(sys.argv[1:])
