import './app.css'
import App from './App.svelte'

const app = new App({
  target: document.getElementById('app'),
  props: {
	  	path_api: "/",
		// path_api: "http://localhost:1116/",
		version: "0.0.1",
	}
})

export default app
