document.addEventListener("DOMContentLoaded", function() {
    const form = document.querySelector("form");
    const username = document.getElementById("username");
    const password = document.getElementById("password");
    const usernameError = document.createElement("span");
    const passwordError = document.createElement("span");

    usernameError.className = 'error';
    passwordError.className = 'error';
    username.parentNode.insertBefore(usernameError, username.nextSibling);
    password.parentNode.insertBefore(passwordError, password.nextSibling);

    function validateInput(inputElement, errorElement, minLength) {
        inputElement.style.borderColor = inputElement.value.trim().length < minLength ? 'red' : 'green';
        if (inputElement.value.trim().length < minLength) {
            errorElement.textContent = `Must be at least ${minLength} characters long.`;
            errorElement.style.color = 'red';
            return false;
        } else {
            errorElement.textContent = '';
            return true;
        }
    }

    username.addEventListener("input", () => validateInput(username, usernameError, 3));
    password.addEventListener("input", () => validateInput(password, passwordError, 6));

    form.onsubmit = function(e) {
        let isValidUsername = validateInput(username, usernameError, 3);
        let isValidPassword = validateInput(password, passwordError, 6);

        if (!isValidUsername || !isValidPassword) {
            e.preventDefault(); // Prevent form submission if validation fails
        }
    };
});
