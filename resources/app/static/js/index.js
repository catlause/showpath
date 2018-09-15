let world = {
	addNode(xpos, ypos, color) {
		let svg_graph = document.getElementById("svg_graph");
		let svgNS = svg_graph.namespaceURI;
		let node = document.createElementNS(svgNS, "circle");
		node.setAttribute("cx", xpos);
		node.setAttribute("cy", ypos);
		node.setAttribute("r", 8);
		node.setAttribute("stroke", color);
		node.setAttribute("stroke-width", 2);
		node.setAttribute("fill", color);
		node.setAttribute("fill-opacity", 0.5);
		node.onclick = function() {world.addNode(23,32,"green");};
		svg_graph.appendChild(node);
	},
	addEdge(x0, y0, x1, y1, color) {
		let svg_graph = document.getElementById("svg_graph");
		let svgNS = svg_graph.namespaceURI;
		let edge = document.createElementNS(svgNS, "line");
		edge.setAttribute("x1", x0);
		edge.setAttribute("y1", y0);
		edge.setAttribute("x2", x1);
		edge.setAttribute("y2", y1);
		edge.setAttribute("stroke", color);
		edge.setAttribute("stroke-width", 2);
		svg_graph.appendChild(edge);
	},
	init: function() {
		document.addEventListener('astilectron-ready', function() {
			world.getgraph();
		})
	},
	getgraph: function() {
		let message = {"name":"getgraph"};

		astilectron.sendMessage(message, function(message) {
			if (message.name == "error") {
				asticode.notifier.error(message.payload);
				return
			}

			let nodes = message.payload.nodes;
			for (var idx in nodes) {
				let node = nodes[idx];
				world.addNode(node.xpos, node.ypos, node.color);
			}
			let edges = message.payload.edges;
			for (var idx in edges) {
				let edge = edges[idx];
				world.addEdge(edge.x0, edge.y0, edge.x1, edge.y1, edge.color);
			}
		})
	}
}
