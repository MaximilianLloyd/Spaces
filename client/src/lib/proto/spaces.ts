/* eslint-disable */

export const protobufPackage = "spaces_package";

export interface Vec3 {
  x: number;
  y: number;
  z: number;
}

export interface Geometry {
  size: Vec3 | undefined;
  position: Vec3 | undefined;
}

export interface Material {
  color: string;
}

export interface Mesh {
  geometry: Geometry | undefined;
  material: Material | undefined;
}

export interface MeshSequence {
  meshes: Mesh[];
  id: string;
}

export const SPACES_PACKAGE_PACKAGE_NAME = "spaces_package";
