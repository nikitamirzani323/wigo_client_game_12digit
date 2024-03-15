import './app.css'
import App from './App.svelte'

const app = new App({
  target: document.getElementById('app'),
  props: {
	  	path_api: "/",
	  	path_websocket: "pp88turnamen.com",
	  	// path_sse: "http://128.199.241.112:1117/",
		// path_api: "http://localhost:1116/",
	  	// path_websocket: "wigo.isbapps.com",
		// path_websocket: "localhost:1116",
		version: "0.0.1",
	}
})

export default app
