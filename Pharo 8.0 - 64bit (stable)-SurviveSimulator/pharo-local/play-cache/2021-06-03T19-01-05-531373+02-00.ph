v := RTView new.	v @ RTDraggableView.		"We set a dynamic spring layout"	stepping := RTSpringLayoutStepping new view: v.	stepping after: [ v canvas camera focusOnCenter ].	v addAnimation: stepping.	"Add a new circle when pressing on a button"	v canvas addMenu: 'Add circle' callback: [ 		| el |		el := (RTEllipse new size: 20; color: (Color blue alpha: 0.4)) element.		el @ RTDraggable.		el translateTo: 5 atRandom @ 5 atRandom.		v add: el.		stepping addNode: el.		v signalUpdate.	].	"Add a new edge when pressing on a button"	v canvas addMenu: 'Add connected circle' callback: [ 		| el edge |		el := (RTEllipse new size: 20; color: (Color blue alpha: 0.4)) element.		el @ RTDraggable.		el translateTo: 5 atRandom @ 5 atRandom.		v add: el.		edge := RTLine edgeFrom: el to: v elements atRandom.		v add: edge.		stepping addNode: el; addEdge: edge.		v signalUpdate.	].	v 