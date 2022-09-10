import gsap, { Back, Elastic } from "gsap";

export interface Timelines {
	toolbarTl: GSAPTimeline;
}

export function initTimelines(): Timelines {
	const toolbar = document.querySelector('.toolbar');
	const toolbarTl = gsap.timeline();

	const toggle = document.querySelector('.toolbar svg')

	toolbarTl.to(toolbar, { y: 0, duration: 0.5 });
	toolbarTl.to(toggle, { rotate: 0, duration: 0.2 }, 0.2);

	return {
		toolbarTl,
	};
}
