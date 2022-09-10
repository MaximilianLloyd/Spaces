import { AppState } from "./draw";
import axios from "axios";
import { Timelines } from "./animations";

interface Callbacks {
	onPurge: () => void;
}

export function initHandlers(
	appState: AppState,
	callbacks: Callbacks,
	timelines: Timelines
) {
	const purgeButton = document.querySelector(".purge-button");
	const toggleButton = document.querySelector(".toggle");

	purgeButton.addEventListener("click", async () => {
		console.log("Purge button");

		const { data } = await axios.get("/purge");
		appState.server = data;
		appState.drawnIds = [];
		callbacks.onPurge();
	});

	toggleButton.addEventListener("click", () => {	
		const tl = timelines.toolbarTl;
		if (tl.reversed()) {
			tl.play();
		} else {
			tl.reverse();
		}
	});
}
