func detectCycle(head *ListNode) *ListNode {
    var slow *ListNode = head
    var fast *ListNode = head
    cycleFound := false

    // let's find the meeting point first
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            cycleFound = true
            break
        }
    }

    if !cycleFound { return nil }

    // the distance to cycle start from both the 
    // head and meeting point are same
    for slow != head {
        head = head.Next
        slow = slow.Next
    }
    return head
}
