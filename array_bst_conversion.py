from collections import deque

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution(object):
    def sortedArrayToBST(self, nums):
        """
        :type nums: List[int]
        :rtype: TreeNode
        """
        return self.constructBST(nums, 0, len(nums)-1)
    
    def constructBST(self, nums, startIdx, endIdx):
        if startIdx > endIdx:
            return None
        
        midIdx = (startIdx + endIdx) // 2
        bst = TreeNode(nums[midIdx])
        bst.left = self.constructBST(nums, startIdx, midIdx - 1)
        bst.right = self.constructBST(nums, midIdx + 1, endIdx)
        return bst
    
    def printBST(self, root):
        queue = deque([(root, 0)])
        nodes = []
        while queue:
            currentNode, currentLevel = queue.popleft()
            if len(nodes) == currentLevel:
                nodes.append([currentNode.val])
            else:
                nodes[currentLevel].append(currentNode.val)

            if currentNode.left is not None:
                queue.append((currentNode.left, currentLevel+1))
            if currentNode.right is not None:
                queue.append((currentNode.right, currentLevel+1))
        return nodes
    
    def bstToArray(self, bst):
        order = []
        self.inOrderTraversal(bst, order)
        return order
    

    def inOrderTraversal(self, node, order):
        if node is None:
            return None
        
        self.inOrderTraversal(node.left, order)
        order.append(node.val)
        self.inOrderTraversal(node.right, order)



if __name__ == "__main__":
    solution = Solution()
    array = [-5, 1, 5, 8, 12, 16, 19, 22, 66]
    bst = solution.sortedArrayToBST(array)
    print("bst", solution.printBST(bst))
    print("sorted array", solution.bstToArray(bst))

'''
output:
('bst', [[12], [1, 19], [-5, 5, 16, 22], [8, 66]])
('sorted array', [-5, 1, 5, 8, 12, 16, 19, 22, 66])
'''
