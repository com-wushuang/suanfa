![bin_tree_example](https://github.com/com-wushuang/suanfa/blob/main/image/bin_tree_example.webp)

# 递归遍历
## 先序遍历
根左右。a，b，d，e，c，f，g
```go
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preOrder(root.Left)
	preOrder(root.Right)
}
```

## 中序遍历
左根右。d，b，e，a，f，c，g
```go
func inOrder(root *Node) {
	if root != nil {
		return
	}
	inOrder(root.Left)
	fmt.Print(root.Val, " ")
	inOrder(root.Right)
}
```

## 后序遍历
左右根。d，e，b，f，g，c，a
```go
func postOrder(root *Node) {
	if root != nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Print(root.Val, " ")
}
```

# 非递归遍历
## 先序遍历
- 不需要入栈，每次遍历到"左"节点，立即输出即可。
- 遍历到最左下的节点时，实际上输出的已经不再是实际的"根"节点，而是实际的"左"节点。这符合先序的定义。
- 因为已经访问过所有"左"节点，现在只需要将这些没用的节点出栈，然后转向到"右"节点。于是“右”节点也变成了“左”节点，后续处理同上。
```
private List<Integer> dfsPreOrder(TreeNode root) {
	ArrayList<Integer> results = new ArrayList<>();
	Stack<TreeNode> stack = new Stack<>();

	TreeNode cur = root;
	while (cur != null || !stack.empty()) {
		while (cur != null) {
			results.add(cur.val);
			stack.push(cur);
			cur = cur.left;
		}
		cur = stack.pop();
		// 转向
		cur = cur.right;
	}
	return results;
}
```

## 中序遍历
- 基于对先序的分析，先序与中序的区别只在于对"左"节点的处理上，我们调整一行代码即可完成中序遍历。
- 我们在出栈之后才访问这个节点。因为先序先访问实际根，后访问实际左，而中序恰好相反。
- 相同的是，访问完根+左子树（先序）或左子树+根（中序）后，都需要转向到"右"节点，使"右"节点称为新的"左"节点。
```
private List<Integer> dfsInOrder(TreeNode root) {
    List<Integer> results = new ArrayList<>();
    Stack<TreeNode> stack = new Stack<TreeNode>();
    TreeNode cur = root;
    while (cur != null || !stack.empty()) {
        while (cur != null) {
            stack.push(cur);
            cur = cur.left;
        }
        cur = stack.pop();
        results.add(cur.val);
        cur = cur.right;
    }
    return results;
}
```

## 后续遍历
- 入栈顺序不变，我们只需要考虑第3点的变化（合适时机转向）。出栈的对象一定都是"左"节点（"右"节点会在转向后称为"左"节点，然后入栈），也就是实际的左或根；实际的左可以当做左右子树都为null的根，所以我们只需要分析实际的根。
- 对于实际的根，需要保证先后访问了左子树、右子树之后，才能访问根。实际的右节点、左节点、根节点都会成为"左"节点入栈，所以我们只需要在出栈之前，将该节点视作实际的根节点，并检查其右子树是否已被访问即可。如果不存在右子树，或右子树已被访问了，那么可以访问根节点，出栈，并不需要转向；如果还没有访问，就转向，使其"右"节点成为"左"节点，等着它先被访问之后，再来访问根节点。
- 所以，我们需要增加一个标志，记录右子树的访问情况。由于访问根节点前，一定先紧挨着访问了其右子树，所以我们只需要一个标志位。
```
private List<Integer> dfsPostOrder(TreeNode root) {
    List<Integer> results = new ArrayList<>();
    Stack<TreeNode> stack = new Stack<>();
    
    TreeNode cur = root;
    TreeNode last = null;
    while(cur != null || !stack.empty()){
        while (cur != null) {
            stack.push(cur);
            cur = cur.left;
        }
        cur = stack.peek();
        if (cur.right == null || cur.right == last) {
            results.add(cur.val);
            stack.pop();
            // 记录上一个访问的节点
            // 用于判断“访问根节点之前，右子树是否已访问过”
            last = cur;
            // 表示不需要转向，继续弹栈
            cur = null;
        } else {
            cur = cur.right;
        }
    }
    return results;
}
```