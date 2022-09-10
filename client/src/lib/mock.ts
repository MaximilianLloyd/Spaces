import { MeshSequence, Material, Geometry, Mesh } from "./proto/spaces";

const mat: Material = {
	color: "#FF0000",
};

const meshes: Mesh[] = [];

for (let i = 0; i < 10; i++) {
	const geo: Geometry = { position: { x: i, y: i, z: i, },
		size: {
			x: 1,
			y: 1,
			z: 1,
		},
	};

	meshes.push({
		geometry: geo,
		material: mat,
	});
}

export const mockSequence: MeshSequence = {
	id: "stupid-line",
	meshes,
};
