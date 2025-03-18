package main

import (
	"os"

	"iap/thirdparty/emicklei/dot"
	"iap/thirdparty/emicklei/dot/dotx"
)

func main() {
	savePath := "/mnt/c/Users/eristocrates/notes/graphviz/"

	g := dot.NewGraph(dot.Directed)
	g.Node("test")

	ex := dotx.NewComposite(savePath+"example", g, dotx.ExternalGraph)

	ex.Export(func(g *dot.Graph) {
		c1 := ex.Node("component")

		sub := dotx.NewComposite(savePath+"subsystem", ex.Graph, dotx.ExternalGraph)
		sub.Input("in1", c1)
		sub.Input("in2", c1)
		sub.Output("out2", c1)

		sub.Export(func(g *dot.Graph) {
			sc1 := sub.Node("subcomponent 1")
			sc2 := sub.Node("subcomponent 2")
			sub.Input("in1", sc1)
			sub.Input("in2", sc2)
			sub.Output("out2", sc2)
			sc1.Edge(sc2)

			sub2 := dotx.NewComposite(savePath+"subsystem2", sub.Graph, dotx.ExternalGraph)
			sub2.Export(func(g *dot.Graph) {
				sub2.Input("in3", sc1)
				sub2.Output("out3", sc2)
				sub3 := sub2.Node("subcomponent 3")
				sub2.Input("in3", sub3)
			})
		})
	})

	os.WriteFile(savePath+"g.dot", []byte(g.String()), os.ModePerm)
	// os.WriteFile("/mnt/c/Users/eristocrates/notes/TestExampleSubsystemSameGraph.dot", []byte(g.String()), os.ModePerm)
}
