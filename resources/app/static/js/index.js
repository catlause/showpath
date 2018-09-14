let index = {
	addNode(xpos, ypos) {
		let div = document.createElement("div");
		div.className = "node";
		div.onclick = function() {index.addNode(23,32);};
		div.innerHTML = '';
		div.style.left = xpos+'px';
		div.style.top = ypos+'px';
		document.body.appendChild(div);
	},
	init: function() {
		document.addEventListener('astilectron-ready', function() {
			index.getgraph();
		})
	},
	getgraph: function() {
		let message = {"name":"getgraph"};

		astilectron.sendMessage(message, function(message) {
			if (message.name == "error") {
				asticode.notifier.error(message.payload);
				return
			}

			let coords = message.payload.coords;
			for (var idx in message.payload.coords) {
				let coord = coords[idx];
				index.addNode(coord.xpos, coord.ypos);
			}
		})
	}
}
