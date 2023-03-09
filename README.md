# Linked-List-Cycle-II

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Example 1:

![circularlinkedlist](https://user-images.githubusercontent.com/88260025/223997097-02de67fa-7a90-4820-ab63-2df6b679d2c3.png)

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:

![circularlinkedlist_test2](https://user-images.githubusercontent.com/88260025/223997222-2d2688a4-dc2d-4083-8688-791fe300405c.png)

Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:

![circularlinkedlist_test3](https://user-images.githubusercontent.com/88260025/223997429-a802498d-d64b-46b4-b9a4-ed3cf839fc7f.png)

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
 

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.
