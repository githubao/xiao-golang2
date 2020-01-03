// 节点
// author: baoqiang
// time: 2019/3/1 下午3:01
package impl

import "fmt"

type KValue struct {
	Key   string
	Value string
}

type Node struct {
	Data     KValue
	NextNode *Node
}

func CreateHead(data KValue) *Node {
	var head = &Node{Data: data, NextNode: nil}
	return head
}

func AddNode(data KValue, node *Node) *Node {
	var newNode = &Node{data, nil}
	node.NextNode = newNode
	return newNode
}

func ShowNode(head *Node) {
	node := head
	for {
		if node.NextNode != nil {
			fmt.Println(node.Data)
			node = node.NextNode
		} else {
			break
		}
	}
	fmt.Println(node.Data)
}

func TailNode(head *Node) *Node {
	node := head
	for {
		if node.NextNode == nil {
			return node
		} else {
			node = node.NextNode
		}
	}

}

func FindValueByKey(key string, head *Node) string {
	node := head
	for {
		if node.NextNode != nil {
			if node.Data.Key == key {
				return node.Data.Value
			}
			node = node.NextNode
		} else {
			break
		}
	}
	return node.Data.Value
}
