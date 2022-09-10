import "./style.css";
import * as THREE from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls.js";
import { BoxGeometry } from "three";

import { placeInGrid } from "./lib/position";
import { mockSequence, initHandlers } from "./lib";
import { AppState, renderAppState, renderMeshSequence } from "./lib/draw";
import axios from "axios";
import { MeshSequence } from "./lib/proto/spaces";
import { initTimelines } from "./lib/animations";

/**
 * Base
 */
// Canvas
const canvas: any = document.querySelector("canvas.webgl");

// Scene
const scene = new THREE.Scene();

/**
 * Sizes
 */
const sizes = {
	width: window.innerWidth,
	height: window.innerHeight,
};

window.addEventListener("resize", () => {
	// Update sizes
	sizes.width = window.innerWidth;
	sizes.height = window.innerHeight;

	// Update camera
	camera.aspect = sizes.width / sizes.height;
	camera.updateProjectionMatrix();

	// Update renderer
	renderer.setSize(sizes.width, sizes.height);
	renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
});

/**
 * Camera
 */
// Base camera
const camera = new THREE.PerspectiveCamera(
	75,
	sizes.width / sizes.height,
	0.1,
	300
);
camera.position.x = 15;
camera.position.y = 15;
camera.position.z = 15;
scene.add(camera);

// Controls
const controls = new OrbitControls(camera, canvas);
controls.enableDamping = true;

/**
 * Cube
 */

//
const pixel = new THREE.Mesh(
	new THREE.BoxGeometry(1, 1, 1),
	new THREE.MeshBasicMaterial({ color: "F68657" })
);

const grid = new THREE.Mesh(
	new THREE.BoxGeometry(10, 10, 10, 10, 10, 10),
	new THREE.MeshBasicMaterial({
		color: "#F68657",
		wireframe: true,
		opacity: 0.3,
		transparent: true,
	})
);

grid.name = 'grid'

grid.position.set(5, 5, 5);

pixel.position.copy(
	placeInGrid({
		x: 9,
		y: 0,
		z: 0,
		mesh: pixel,
	})
);

// scene.add(pixel);

scene.add(grid);

// const test = renderMeshSequence(mockSequence);

// scene.add(test);

/**
 * Renderer
 */
const renderer = new THREE.WebGLRenderer({
	canvas: canvas,
	antialias: true,
});
renderer.setSize(sizes.width, sizes.height);
renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
renderer.setClearColor(new THREE.Color("#1F2124"));
/*
 * 490a3d
 * bd1550
 * e97f02
 * f8ca00
 */
/**
 * Animate
 */
const clock = new THREE.Clock();
let lastElapsedTime = 0;

const tick = () => {
	const elapsedTime = clock.getElapsedTime();
	const deltaTime = elapsedTime - lastElapsedTime;
	lastElapsedTime = elapsedTime;

	// Update controls
	controls.update();

	// Render
	renderer.render(scene, camera);

	// Call tick again on the next frame
	window.requestAnimationFrame(tick);
};

tick();

let appState: AppState = {
	// Go state
	server: {
		MeshSequences: [],
	},
	drawnIds: [],
};

const timelines = initTimelines()
initHandlers(appState, {
	onPurge: () => {
		renderAppState(appState, scene);
	},
}, timelines);



setInterval(() => {
	axios.get("/poll").then((response) => {
		const data = response.data;
		appState.server = data;
		renderAppState(appState, scene);
	});
}, 2000);
