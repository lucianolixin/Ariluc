## Linked List

###  特点

链表的内存是不连续的，随意不支持内存的随机访问,天然支持扩容。

- 头节点，记录链表的基地址，有了它就可以遍历得到整条链表。
- 尾节点，指针指向一个空地址null，表示这个链表上最后一个节点。

### 时间复杂度

- Access O(n)
- Insert O(1)
- Delete O(1)

### 循环链表

循环链表是一种特殊的单链表，唯一的区别就是尾节点指向头结点。


### 双向链表

双向链表不但有一个后继指针，还有一个前驱指针，指向之前节点。
优点：支持双向遍历。在插入删除等操作比单链表更简单，高效（删除链表中的节点，需要知道节点的前驱节点）。
缺点：双向链表，每个节点额外需要存储两个内存地址，所以更加占用存储空间。

### 代码技巧

a -> b ==> a -> x -> b

1.插入操作：一定注意操作顺序，要先将带插入节点x的next指针指向b,再把节点a的next指针指向节点x。这样才不会丢失指针，导致内存泄漏。

2.利用哨兵简化实现难度

单链表的插入和删除操作，如果我们的节点p后面插入一个新节点，只需要下面两行代码就可以搞定

```
new_node.next = p.next
p.next = new_node
```
但是我们要向一个空链表中插入第一个节点，刚刚的逻辑就不能用了。我们需要进行下面这样的特殊处理，其中head表示链表的头节点。所以，从这段代码我们能发现对于单链表的插入操作，第一个节点和其他节点是不同的。

```
if head == nil {
    head = new_node
}
```
如果我们要删除链表中的最后一个节点，也不能直接让被删除的前驱节点的后继指针直接指向被删除的后继指针。
我们可以看出，针对链表的插入、删除操作，需要对插入第一个节点和删除最后一个节点的情况进行特殊处理，**哨兵**是解决边界问题而被引入的，不直接参与业务逻辑。

如果引入哨兵节点，在任何时候，不管链表是不是空，head指针都会一直指向这个哨兵节点。我们把这种有哨兵节点的链表叫**带头链表。

3.注意边界条件

- 如果链表为空，代码能否正常运行
- 如果链表只包含一个节点，代码能否正常运行
- 如果链表只包含两个节点，代码能否正常运行
- 如果在处理链表的头节点和尾节点的时候，代码能否正常运行
