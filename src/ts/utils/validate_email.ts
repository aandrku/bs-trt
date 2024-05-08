type EmailValidationResult = {
    isValid: boolean;
    message: string;
};

export function validateEmail(email: string): EmailValidationResult {
    // Regex source: https://emailregex.com/
    const regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    
    if (regex.test(email)) {
        return { isValid: true, message: "Email is valid." };
    } else {
        return { isValid: false, message: "Email is invalid." };
    }
}
