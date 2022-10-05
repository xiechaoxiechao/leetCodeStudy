package Solutions

  type Node struct {
      Val int
      Prev *Node
      Next *Node
     Child *Node
  }

func flatten(root *Node) *Node {
    if root==nil{
        return nil
    }
	unfold(root)
	return root
}

func unfold(root *Node) *Node {
	if root.Child != nil {
        next := root.Next
		root.Child.Prev = root
        root.Next = root.Child
		childListTile := unfold(root.Child)
        
		root.Child = nil
        if next!=nil{
            childListTile.Next = next
		    next.Prev = childListTile
            return unfold(next)
        }else{
          
            return childListTile
        }
	
	}else{
        if root.Next!=nil{
            return unfold(root.Next)
        }else{
           
            return root
        }
    }
		
}
