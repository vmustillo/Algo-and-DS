package tree

func IsSubtree(s *Node, t *Node) bool {
	if s == nil && t == nil {
		return true
	} else if t == nil {
		return true
	} else if s == nil && t != nil {
		return false
	} else {
		return isEqual(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
	}
}

func isEqual(s *Node, t *Node) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	} else {
		return s.Val == t.Val && isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
	}
}
