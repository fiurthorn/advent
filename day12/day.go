package day12

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/fiurthorn/advent/lib"
)

type Day struct {
	tree *Tree
}

//go:embed example.txt
var dayExample string

//go:embed data.txt
var dayData string

func (d Day) Run() {
	log.Printf("Example:  %v", d.process(dayExample))
	log.Printf("Solution: %v", d.process(dayData))
}

func (d Day) process(data string) string {
	lines := lib.Lines(data)
	return fmt.Sprintf("1:%v 2:%v",
		d.process1(lines),
		d.process2(lines),
	)
}

type TreeNode struct {
	Name     string
	Foreward []*TreeNode
}

func (tn *TreeNode) String() string {
	names := []string{}
	for _, name := range tn.Foreward {
		names = append(names, name.Name)
	}
	return fmt.Sprintf("%s -> %v", tn.Name, names)
}

type Tree struct {
	Start *TreeNode
	End   *TreeNode
	Index map[string]*TreeNode

	Solutions  []*lib.StringSet
	Solutions2 []string
}

func (t *Tree) String() string {
	return fmt.Sprintf("Start: %v\nEnd  : %v\n%s", t.Start, t.End, t.Index)
}

func (t *Tree) add(a, b string) {
	aNode, bNode := t.Index[a], t.Index[b]

	aNode.Foreward = append(aNode.Foreward, bNode)
	bNode.Foreward = append(bNode.Foreward, aNode)

	if a == "start" && t.Start == nil {
		t.Start = aNode
	}
	if a == "end" && t.End == nil {
		t.End = aNode
	}

	if b == "start" && t.Start == nil {
		t.Start = bNode
	}
	if b == "end" && t.End == nil {
		t.End = bNode
	}
}

func (d Day) process1(lines []string) string {
	d.tree = &Tree{Index: map[string]*TreeNode{}}

	for _, line := range lines {
		nodes := strings.Split(line, "-")
		aName, bName := nodes[0], nodes[1]

		if _, has := d.tree.Index[aName]; !has {
			d.tree.Index[aName] = &TreeNode{Name: aName}
		}
		if _, has := d.tree.Index[bName]; !has {
			d.tree.Index[bName] = &TreeNode{Name: bName}
		}

		d.tree.add(aName, bName)
	}

	visited := lib.NewStringSetWith("start")
	d.traverse(d.tree.Start, visited)

	return fmt.Sprintf("%v", len(d.tree.Solutions))
}

func (d Day) traverse(tn *TreeNode, v *lib.StringSet) {
	if tn.Name == "end" {
		d.tree.Solutions = append(d.tree.Solutions, v.Clone())
		return
	}
	for _, node := range tn.Foreward {
		if d.isUpper(node.Name[0]) || !v.Has(node.Name) {
			d.traverse(node, v.CloneWith(node.Name))
		}
	}

}

func (d Day) traverse2(tn *TreeNode, v *lib.StringSet, twice string, way string) {
	if tn.Name == "end" {
		d.tree.Solutions2 = append(d.tree.Solutions2, way)
		return
	}
	for _, node := range tn.Foreward {
		if d.isUpper(node.Name[0]) || !v.Has(node.Name) {
			d.traverse2(node, v.CloneWith(node.Name), twice, way+"-"+node.Name)
		} else if v.Has(node.Name) && twice == "" && node.Name != "start" && node.Name != "end" {
			d.traverse2(node, v.CloneWith(node.Name), node.Name, way+"-"+node.Name)
		}
	}
}

func (d Day) isUpper(r byte) bool {
	return r >= 'A' && r <= 'Z'
}

func (d Day) process2(lines []string) string {
	d.tree = &Tree{Index: map[string]*TreeNode{}}

	for _, line := range lines {
		nodes := strings.Split(line, "-")
		aName, bName := nodes[0], nodes[1]

		if _, has := d.tree.Index[aName]; !has {
			d.tree.Index[aName] = &TreeNode{Name: aName}
		}
		if _, has := d.tree.Index[bName]; !has {
			d.tree.Index[bName] = &TreeNode{Name: bName}
		}

		d.tree.add(aName, bName)
	}

	visited := lib.NewStringSetWith("start")
	d.traverse2(d.tree.Start, visited, "", "start")

	return fmt.Sprintf("%v", len(d.tree.Solutions2))
}
