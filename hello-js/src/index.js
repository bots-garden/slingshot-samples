import { ExecHandler } from "./core/slingshot"

function handle() {
	
	ExecHandler(param => {
		let output = "param: " + param
		let err = null

		return [output, err]
	})
}

module.exports = {handle}
