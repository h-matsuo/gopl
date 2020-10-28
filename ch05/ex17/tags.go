package tags

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return visit(nil, doc, name...)
}

func visit(nodes []*html.Node, n *html.Node, name ...string) []*html.Node {
	if n == nil {
		return nodes
	}
	if n.Type == html.ElementNode {
		for _, tag := range name {
			if tag == n.Data {
				nodes = append(nodes, n)
				break
			}
		}
	}
	nodes = visit(nodes, n.FirstChild, name...)
	nodes = visit(nodes, n.NextSibling, name...)
	return nodes
}
