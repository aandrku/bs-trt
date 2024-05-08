import {validateEmail} from "./utils/validate_email"
import {validatePassword} from "./utils/validate_password"
import {displayError} from "./utils/displayError"

document.addEventListener("DOMContentLoaded", async function(){
	let form = document.getElementById("signup-form") as HTMLFormElement
	form?.addEventListener("submit", async function(event) {
		event.preventDefault()

		//get form data
		let formData = new FormData(form)
		let email = formData.get("email")
		let tag = formData.get("tag")
		let password = formData.get("password")
		let confirmPassword = formData.get("confirm-password")

		//obtain a display html element
		let display = document.getElementById("error")
		if (display == null) return

		//validate email
		if (email == null) {
			displayError("Provide an email.", display)
			return
		}
		email = email.toString()
		let emailValidationResult = validateEmail(email)
		if (!emailValidationResult.isValid) {
			displayError(emailValidationResult.message, display)
			return
		}

		//validate a tag
		if (tag == null) {
			displayError("Provide a game tag.", display)
			return
		}
		tag = tag.toString()

		//validate password
		if (password == null || confirmPassword == null) { displayError("Provide a password", display)
			return
		}
		password = password.toString()
		confirmPassword = confirmPassword.toString()
		if (password != confirmPassword) {
			displayError("Passwords should match", display)
		}
		let passwordValidationResult = validatePassword(password)
		if (!passwordValidationResult.isValid) {
			displayError(passwordValidationResult.message, display)
			return
		}	
		let confirmPasswordValidationResult = validatePassword(password)
		if (!confirmPasswordValidationResult.isValid) {
			displayError(confirmPasswordValidationResult.message, display)
			return
		}	

		try {
			let data = new URLSearchParams();
			data.append("email", email)
			data.append("tag", tag)
			data.append("password", password)
			data.append("confirm-password", confirmPassword)
			const response = await fetch("https://localhost:8080/signup", {
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
