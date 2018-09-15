let world = {
	addNode(xpos, ypos) {
		let div = document.createElement("div");
		div.className = "node";
		div.onclick = function() {world.addNode(23,32);};
		div.innerHTML = '';
		div.style.left = (xpos-8)+'px';
		div.style.top = (ypos-8)+'px';
		document.getElementById("container").appendChild(div);
	},
	addEdge(x0, y0, x1, y1) {
		let line = "<line x1='"+x0+"px' y1='"+y0+"px' x2='"+x1+"px' y2='"+y1+"px' stroke='black' stroke-width='2' />";
		let bgsvg = 'url("'+"data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' version='1.1'>"+line+'</svg>")';
		document.getElementById("container").style.backgroundImage=bgsvg;
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
				world.addNode(node.xpos, node.ypos);
			}
			let edges = message.payload.edges;
			for (var idx in edges) {
				let edge = edges[idx];
				world.addEdge(edge.x0, edge.y0, edge.x1, edge.y1);
			}
		})
	}
}
