import * as THREE from "three";
import gsap from 'gsap'
import { placeInGrid } from "./position";
import { MeshSequence } from "./proto/spaces";

export function renderMeshSequence(meshSequence: MeshSequence): THREE.Group {
	const group = new THREE.Group();

	group.name = meshSequence.id;


	// For animations
	for (const [i, mesh] of meshSequence.meshes.entries()) {
		const threeGeometry = new THREE.BoxGeometry(
			mesh.geometry.size.x,
			mesh.geometry.size.y,
			mesh.geometry.size.z
		);

		const threeMaterial = new THREE.MeshBasicMaterial({
			color: new THREE.Color(mesh.material.color),
			transparent: true,
		});

		const threeMesh = new THREE.Mesh(threeGeometry, threeMaterial);

		threeMesh.position.copy(
			placeInGrid({
				x: mesh.geometry.position.x - 1,
				y: mesh.geometry.position.y - 1,
				z: mesh.geometry.position.z - 1,
				mesh: threeMesh,
			})
		);

		threeMesh.scale.set(0, 0, 0)
		gsap.to(threeMesh.scale, { x: 1, y: 1, z: 1, delay: i * 0.1 })

		group.add(threeMesh);
	}

	return group;
}

export interface AppState {
	server: {
		MeshSequences: MeshSequence[];
	};
	drawnIds: string[];
}

export function renderAppState(state: AppState, threeScene: THREE.Scene) {
	const { server } = state;
	if (!server.MeshSequences) return;

	for (const sequence of server.MeshSequences) {
		const { id } = sequence;

		const shouldDraw = !state.drawnIds.includes(id);

		if (shouldDraw) {
			console.log(`Drawing ${id}`);
			const group = renderMeshSequence(sequence);
			threeScene.add(group);
  
			state.drawnIds.push(id)
		}
	}
	// Check if should remove
	threeScene.traverse((object) => {
		const isGroup = object.type === "Group";


		if (isGroup) {
			console.log(object.name, state.drawnIds)
			const shouldRemove = !state.drawnIds.includes(object.name);

			if (shouldRemove) {
				console.log('REMOVING', object.name)
				threeScene.remove(object);
			}
		}
	});


}
