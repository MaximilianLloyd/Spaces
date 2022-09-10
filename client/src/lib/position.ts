import * as THREE from 'three'


// Assume a 10x10 grid
interface PlaceInGridOptions {
  x: number;
  y: number;
  z: number;
  mesh: THREE.Mesh;
}

export function placeInGrid({ x, y, z, mesh }: PlaceInGridOptions): THREE.Vector3 {
  // grid.geometry.computeBoundingBox()
  mesh.geometry.computeBoundingBox()

  // const gridSize = grid.geometry.boundingBox.getSize(new THREE.Vector3());
  const meshSize = mesh.geometry.boundingBox.getSize(new THREE.Vector3());

  const offsetVec = new THREE.Vector3(meshSize.x / 2, meshSize.y / 2, meshSize.z / 2)

  // const { x: gridWidth, y: gridHeight, z: gridDepth } = gridSize;

  const vec = new THREE.Vector3(x, y, z)
  vec.add(offsetVec)
  return vec;
}
