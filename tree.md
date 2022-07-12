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

## 统计二叉树节点个数
```
public int countNodes(TreeNode root) {
    if (root == null) return 0;
    return 1 + countNodes(root.left) + countNodes(root.right);
}
```

## 二叉树的最大深度
**方法1**
- 所谓最大深度就是根节点到「最远」叶子节点的最长路径上的节点数。
- 遍历一遍二叉树，用一个外部变量记录每个节点所在的深度，取最大值就可以得到最大深度。
- 在前序位置增加 depth，在后序位置减小 depth。
- 因为前序位置是进入一个节点的时候，后序位置是离开一个节点的时候，depth 记录当前递归到的节点深度，你把 traverse 理解成在二叉树上游走的一个指针，所以当然要这样维护。
- 至于对 res 的更新，你放到前中后序位置都可以，只要保证在进入节点之后，离开节点之前（即 depth 自增之后，自减之前）就行了。


```
// 记录最大深度
int res = 0;
// 记录遍历到的节点的深度
int depth = 0;

// 主函数
int maxDepth(TreeNode root) {
	traverse(root);
	return res;
}

// 二叉树遍历框架
void traverse(TreeNode root) {
	if (root == null) {
		return;
	}
	// 前序位置
	depth++;
    if (root.left == null && root.right == null) {
        // 到达叶子节点，更新最大深度
		res = Math.max(res, depth);
    }
	traverse(root.left);
	traverse(root.right);
	// 后序位置
	depth--;
}
```
**方法2**
```
int maxDepth(TreeNode root) {
	if (root == null) {
		return 0;
	}
	// 利用定义，计算左右子树的最大深度
	int leftMax = maxDepth(root.left);
	int rightMax = maxDepth(root.right);
	// 整棵树的最大深度等于左右子树的最大深度取最大值，
    // 然后再加上根节点自己
	int res = Math.max(leftMax, rightMax) + 1;

	return res;
}
```
## 二叉树的直径
- 求二叉树的最长直径
- 二叉树的「直径」长度，就是任意两个结点之间的路径长度。最长「直径」并不一定要穿过根结点。
- 每一条二叉树的「直径」长度，就是一个节点的左右子树的最大深度之和。
- 思路就是遍历整棵树中的每个节点，然后通过每个节点的左右子树的最大深度算出每个节点的「直径」，最后把所有「直径」求个最大值即可。
```
// 记录最大直径的长度
int maxDiameter = 0;

public int diameterOfBinaryTree(TreeNode root) {
    maxDepth(root);
    return maxDiameter;
}

int maxDepth(TreeNode root) {
    if (root == null) {
        return 0;
    }
    int leftMax = maxDepth(root.left);
    int rightMax = maxDepth(root.right);
    // 后序位置，顺便计算最大直径
    int myDiameter = leftMax + rightMax;
    maxDiameter = Math.max(maxDiameter, myDiameter);

    return 1 + Math.max(leftMax, rightMax);
}
```
## 二叉树层序遍历(广度优先搜索)
- 我们需要借助一个队列，由于队列是先进先出的，所以可以先让根入队，在根出队的同时打印根，并让根的左孩子入队，再让右孩子入队，这样一来，左孩子结点就存储在队头的位置，可以首先被访问；
- 被访问之后，左孩子出队的同时打印左孩子，并让左孩子的孩子入队；此时队列的下一个元素就是右孩子，右孩子出队的同时打印右孩子，并让右孩子的孩子入队；
- 以此类推，直到队列为空。
```
// 输入一棵二叉树的根节点，层序遍历这棵二叉树
void levelTraverse(TreeNode root) {
    if (root == null) return;
    Queue<TreeNode> q = new LinkedList<>();
    q.offer(root);

    // 从上到下遍历二叉树的每一层
    while (!q.isEmpty()) {
        int sz = q.size();
        // 从左到右遍历每一层的每个节点
        for (int i = 0; i < sz; i++) {
            TreeNode cur = q.poll();
            // 将下一层节点放入队列
            if (cur.left != null) {
                q.offer(cur.left);
            }
            if (cur.right != null) {
                q.offer(cur.right);
            }
        }
    }
}
```
## 翻转二叉树
**遍历思维**
```
// 主函数
TreeNode invertTree(TreeNode root) {
    // 遍历二叉树，交换每个节点的子节点
    traverse(root);
    return root;
}

// 二叉树遍历函数
void traverse(TreeNode root) {
    if (root == null) {
        return;
    }

    /**** 前序位置 ****/
    // 每一个节点需要做的事就是交换它的左右子节点
    TreeNode tmp = root.left;
    root.left = root.right;
    root.right = tmp;

    // 遍历框架，去遍历左右子树的节点
    traverse(root.left);
    traverse(root.right);
}
```
**分解问题思维**
```
// 定义：将以 root 为根的这棵二叉树翻转，返回翻转后的二叉树的根节点
TreeNode invertTree(TreeNode root) {
    if (root == null) {
        return null;
    }
    // 利用函数定义，先翻转左右子树
    TreeNode left = invertTree(root.left);
    TreeNode right = invertTree(root.right);

    // 然后交换左右子节点
    root.left = right;
    root.right = left;

    // 和定义逻辑自恰：以 root 为根的这棵二叉树已经被翻转，返回 root
    return root;
}
```

## 将二叉树展开为链表
![tree_to_node](https://github.com/com-wushuang/suanfa/blob/main/image/tree_to_node.png)
```
// 虚拟头节点，dummy.right 就是结果
TreeNode dummy = new TreeNode(-1);
// 用来构建链表的指针
TreeNode p = dummy;

void traverse(TreeNode root) {
    if (root == null) {
        return;
    }
    // 前序位置
    p.right = new TreeNode(root.val);
    p = p.right;

    traverse(root.left);
    traverse(root.right);
}
```

## 构造最大二叉树
![max_tree](https://github.com/com-wushuang/suanfa/blob/main/image/max_tree.png)
- 每个二叉树节点都可以认为是一棵子树的根节点，对于根节点，首先要做的当然是把想办法把自己先构造出来，然后想办法构造自己的左右子树。
- 我们要遍历数组把找到最大值 maxVal，从而把根节点 root 做出来，然后对 maxVal 左边的数组和右边的数组进行递归构建，作为 root 的左右子树。
```
/* 主函数 */
TreeNode constructMaximumBinaryTree(int[] nums) {
    return build(nums, 0, nums.length - 1);
}

// 定义：将 nums[lo..hi] 构造成符合条件的树，返回根节点
TreeNode build(int[] nums, int lo, int hi) {
    // base case
    if (lo > hi) {
        return null;
    }

    // 找到数组中的最大值和对应的索引
    int index = -1, maxVal = Integer.MIN_VALUE;
    for (int i = lo; i <= hi; i++) {
        if (maxVal < nums[i]) {
            index = i;
            maxVal = nums[i];
        }
    }

    // 先构造出根节点
    TreeNode root = new TreeNode(maxVal);
    // 递归调用构造左右子树
    root.left = build(nums, lo, index - 1);
    root.right = build(nums, index + 1, hi);
    
    return root;
}
```

## 通过前序和中序遍历结果构造二叉树
- 题目：
![pre_in_rebuild](https://github.com/com-wushuang/suanfa/blob/main/image/pre_in_rebuild.png)
- 思路：
![rebuild](https://github.com/com-wushuang/suanfa/blob/main/image/rebuild.jpeg)
```
// 存储 inorder 中值到索引的映射
HashMap<Integer, Integer> valToIndex = new HashMap<>();

public TreeNode buildTree(int[] preorder, int[] inorder) {
    for (int i = 0; i < inorder.length; i++) {
        valToIndex.put(inorder[i], i);
    }
    return build(preorder, 0, preorder.length - 1,
                 inorder, 0, inorder.length - 1);
}

TreeNode build(int[] preorder, int preStart, int preEnd, 
               int[] inorder, int inStart, int inEnd) {
        
    if (preStart > preEnd) {
        return null;
    }

    // root 节点对应的值就是前序遍历数组的第一个元素
    int rootVal = preorder[preStart];
    // rootVal 在中序遍历数组中的索引rebuik
    int index = valToIndex.get(rootVal);

    int leftSize = index - inStart;

    // 先构造出当前根节点
    TreeNode root = new TreeNode(rootVal);
    // 递归构造左右子树
    root.left = build(preorder, preStart + 1, preStart + leftSize,
                      inorder, inStart, index - 1);

    root.right = build(preorder, preStart + leftSize + 1, preEnd,
                       inorder, index + 1, inEnd);
    return root;
}
```
## 通过后序和中序遍历结果构造二叉树
- 思路：
![in_post_rebuild](https://github.com/com-wushuang/suanfa/blob/main/image/in_post_rebuild.jpeg)
```
// 存储 inorder 中值到索引的映射
HashMap<Integer, Integer> valToIndex = new HashMap<>();

TreeNode buildTree(int[] inorder, int[] postorder) {
    for (int i = 0; i < inorder.length; i++) {
        valToIndex.put(inorder[i], i);
    }
    return build(inorder, 0, inorder.length - 1,
                 postorder, 0, postorder.length - 1);
}

TreeNode build(int[] inorder, int inStart, int inEnd,
               int[] postorder, int postStart, int postEnd) {

    if (inStart > inEnd) {
        return null;
    }
    // root 节点对应的值就是后序遍历数组的最后一个元素
    int rootVal = postorder[postEnd];
    // rootVal 在中序遍历数组中的索引
    int index = valToIndex.get(rootVal);
    // 左子树的节点个数
    int leftSize = index - inStart;
    TreeNode root = new TreeNode(rootVal);
    // 递归构造左右子树
    root.left = build(inorder, inStart, index - 1,
                        postorder, postStart, postStart + leftSize - 1);
    
    root.right = build(inorder, index + 1, inEnd,
                        postorder, postStart + leftSize, postEnd - 1);
    return root;
}
```
## 通过后序和前序遍历结果构造二叉树
```go
class Solution {
    // 存储 postorder 中值到索引的映射
    HashMap<Integer, Integer> valToIndex = new HashMap<>();

    public TreeNode constructFromPrePost(int[] preorder, int[] postorder) {
        for (int i = 0; i < postorder.length; i++) {
            valToIndex.put(postorder[i], i);
        }
        return build(preorder, 0, preorder.length - 1,
                    postorder, 0, postorder.length - 1);
    }

    // 定义：根据 preorder[preStart..preEnd] 和 postorder[postStart..postEnd]
    // 构建二叉树，并返回根节点。
    TreeNode build(int[] preorder, int preStart, int preEnd,
                   int[] postorder, int postStart, int postEnd) {
        if (preStart > preEnd) {
            return null;
        }
        if (preStart == preEnd) {
            return new TreeNode(preorder[preStart]);
        }

        // root 节点对应的值就是前序遍历数组的第一个元素
        int rootVal = preorder[preStart];
        // root.left 的值是前序遍历第二个元素
        // 通过前序和后序遍历构造二叉树的关键在于通过左子树的根节点
        // 确定 preorder 和 postorder 中左右子树的元素区间
        int leftRootVal = preorder[preStart + 1];
        // leftRootVal 在后序遍历数组中的索引
        int index = valToIndex.get(leftRootVal);
        // 左子树的元素个数
        int leftSize = index - postStart + 1;

        // 先构造出当前根节点
        TreeNode root = new TreeNode(rootVal);
        // 递归构造左右子树
        // 根据左子树的根节点索引和元素个数推导左右子树的索引边界
        root.left = build(preorder, preStart + 1, preStart + leftSize,
                postorder, postStart, index);
        root.right = build(preorder, preStart + leftSize + 1, preEnd,
                postorder, index + 1, postEnd - 1);

        return root;
    }
}
```

## 二叉树的序列化
```
String SEP = ",";
String NULL = "#";

/* 主函数，将二叉树序列化为字符串 */
String serialize(TreeNode root) {
    StringBuilder sb = new StringBuilder();
    serialize(root, sb);
    return sb.toString();
}

/* 辅助函数，将二叉树存入 StringBuilder */
void serialize(TreeNode root, StringBuilder sb) {
    if (root == null) {
        sb.append(NULL).append(SEP);
        return;
    }

    /****** 前序遍历位置 ******/
    sb.append(root.val).append(SEP);
    /***********************/

    serialize(root.left, sb);
    serialize(root.right, sb);
}
```

## 二叉树的反序列化
```
/* 主函数，将字符串反序列化为二叉树结构 */
TreeNode deserialize(String data) {
    // 将字符串转化成列表
    LinkedList<String> nodes = new LinkedList<>();
    for (String s : data.split(SEP)) {
        nodes.addLast(s);
    }
    return deserialize(nodes);
}

/* 辅助函数，通过 nodes 列表构造二叉树 */
TreeNode deserialize(LinkedList<String> nodes) {
    if (nodes.isEmpty()) return null;

    /****** 前序遍历位置 ******/
    // 列表最左侧就是根节点
    String first = nodes.removeFirst();
    if (first.equals(NULL)) return null;
    TreeNode root = new TreeNode(Integer.parseInt(first));
    /***********************/

    root.left = deserialize(nodes);
    root.right = deserialize(nodes);

    return root;
}
```

# 满二叉树和完全二叉树
**满二叉树**
- 通俗定义: 对于满二叉树，除最后一层无任何子节点外，每一层上的所有结点都有两个子结点二叉树。
- 严谨定义: 个二叉树，如果每一个层的结点数都达到最大值，则这个二叉树就是满二叉树。也就是说，如果一个二叉树的层数为 `K` ，且结点总数是 `(2^k) -1` ，则它就是满二叉树。

**完全二叉树**
- 通俗定义: 完全二叉树是效率很高的数据结构，完全二叉树是由满二叉树而引出来的。对于深度为K的，有n个结点的二叉树，当且仅当其每一个结点都与深度为K的满二叉树中编号从1至n的结点一一对应时称之为完全二叉树。
- 严谨定义: 若设二叉树的深度为 `h` ，除第 `h` 层外，其它各层 `(1～h-1)` 的结点数都达到最大个数，第 `h` 层所有的结点都连续集中在最左边，这就是完全二叉树。

#  二叉搜索树
二分搜索树（`Binary Search Tree`），也称为 二叉查找树 、二叉搜索树 、有序二叉树或排序二叉树。满足以下几个条件：
- 若它的左子树不为空，左子树上所有节点的值都小于它的根节点。
- 若它的右子树不为空，右子树上所有的节点的值都大于它的根节点。
它的左、右子树也都是二分搜索树。

## BST 的中序遍历结果是有序的
- 输入一棵 BST，以下代码可以将 BST 中每个节点的值升序打印出来：
```
void traverse(TreeNode root) {
    if (root == null) return;
    traverse(root.left); // 变成右指针，就是降序遍历
    // 中序遍历代码位置
    print(root.val);
    traverse(root.right); // 变成左指针，就是降序遍历
}
```

## 寻找第 K 小的元素
- 给定一个二叉搜索树，编写函数来查找其中第 `k` 个元素。
- 一个直接的思路就是升序排序，然后找第 k 个元素。
```
int kthSmallest(TreeNode root, int k) {
    // 利用 BST 的中序遍历特性
    traverse(root, k);
    return res;
}

// 记录结果
int res = 0;
// 记录当前元素的排名
int rank = 0;
void traverse(TreeNode root, int k) {
    if (root == null) {
        return;
    }
    traverse(root.left, k);
    /* 中序遍历代码位置 */
    rank++;
    if (k == rank) {
        // 找到第 k 小的元素
        res = root.val;
        return;
    }
    /*****************/
    traverse(root.right, k);
}
```

## BST 转化累加树
- ![bst_sum](https://github.com/com-wushuang/suanfa/blob/main/image/bst_sum.png)
- 其实，正确的解法很简单，还是利用 BST 的中序遍历特性。
```
TreeNode convertBST(TreeNode root) {
    traverse(root);
    return root;
}

// 记录累加和
int sum = 0;
void traverse(TreeNode root) {
    if (root == null) {
        return;
    }
    traverse(root.right);
    // 维护累加和
    sum += root.val;
    // 将 BST 转化成累加树
    root.val = sum;
    traverse(root.left);
}
```