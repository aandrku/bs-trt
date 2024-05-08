import {validateEmail} from "./utils/validate_email"
import {validatePassword} from "./utils/validate_password"
import {displayError} from "./utils/displayError"

document.addEventListener("DOMContentLoaded", async function(){
	let form = document.getElementById("login-form") as HTMLFormElement
	form?.addEventListener("submit", async function(event) {
		event.preventDefault()

		//get form data
		let formData = new FormData(form)
		let email = formData.get("email")
		let password = formData.get("password")

		//obtain a display html element
		let display = document.getElementById("error")
		if (display == null) return


		//validate email
		if (email == null) {
			displayError("Provide an email.", display)
			return
		}
		//get a string value from an email
		email = email.toString()
		let emailValidationResult = validateEmail(email)
		if (!emailValidationResult.isValid) {
			displayError(emailValidationResult.message, display)
		}

		//validate password
		if (password == null) { displayError("Provide a password", display)
			return
		}
		//get a string value from a password
		password = password.toString()
		let passwordValidationResult = validatePassword(password)
		if (!passwordValidationResult.isValid) {
			displayError(passwordValidationResult.message, display)
			return
		}	

		try {
			let data = new URLSearchParams();
			data.append("email", email)
			data.append("password", password)
			const response = await fetch("https://localhost:8080/login", {
				method: "POST",
				headers: {
					"Content-Type": "application/x-www-form-urlencoded",
				},
				body: data,
			})

			if (!response.ok) {
				throw new Error("Network response was not ok")
			}


			if (response.redirected) {
				window.location.href = response.url
				return
			}

			const responseData = await response.json()
			displayError(responseData.message, display)
			return

		} catch (error) {
			console.error("There was a problem")
		}
	})
})

