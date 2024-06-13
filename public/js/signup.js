document.addEventListener('DOMContentLoaded', () => {
    const signupForm = document.getElementById('signupForm');
    const signupBtn = document.getElementById('signup-btn');

    signupBtn.disabled = true;

    let upper = false, lower = false, number = false, lenght = false;

    signupForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const response = await fetch('/api/auth/signup', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: signupForm.name.value,
                email: signupForm.email.value,
                username: signupForm.signupUsername.value,
                password: signupForm.signupPassword.value
            })
        });
        if (response.ok) {
            alert('Signup successful');
            window.location.href = 'login.html';
        } else {
            alert('Signup failed');
        }
    });

    var passwordInput = document.getElementById("signupPassword"); 
    var passwordMessageItems = document.getElementsByClassName("password-message-item"); 
    var passwordMessage = document.getElementById("password-message"); 
    
    
    passwordInput.onfocus = function () { 
        passwordMessage.style.display = "block"; 
    } 
    
    passwordInput.onkeyup = function () {
        let uppercaseRegex = /[A-Z]/g; 
        if (passwordInput.value.match(uppercaseRegex)) { 
            passwordMessageItems[1].classList.remove("invalid"); 
            passwordMessageItems[1].classList.add("valid"); 
            upper = true;
        } else { 
            passwordMessageItems[1].classList.remove("valid"); 
            passwordMessageItems[1].classList.add("invalid"); 
            lower = false;
        } 
 
        let lowercaseRegex = /[a-z]/g; 
        if (passwordInput.value.match(lowercaseRegex)) { 
            passwordMessageItems[0].classList.remove("invalid"); 
            passwordMessageItems[0].classList.add("valid"); 
            lower = true;
        } else { 
            passwordMessageItems[0].classList.remove("valid"); 
            passwordMessageItems[0].classList.add("invalid");
            lower = false; 
        } 
        let numbersRegex = /[0-9]/g; 
        if (passwordInput.value.match(numbersRegex)) { 
            passwordMessageItems[2].classList.remove("invalid"); 
            passwordMessageItems[2].classList.add("valid");
            number = true;; 
        } else { 
            passwordMessageItems[2].classList.remove("valid"); 
            passwordMessageItems[2].classList.add("invalid"); 
            number = false;
        } 

        if (passwordInput.value.length >= 8) { 
            passwordMessageItems[3].classList.remove("invalid"); 
            passwordMessageItems[3].classList.add("valid"); 
            lenght = true;
        } else { 
            passwordMessageItems[3].classList.remove("valid"); 
            passwordMessageItems[3].classList.add("invalid"); 
            lenght = false;
        } 
        if(upper && lower && number && lenght){
            signupBtn.disabled = false;
        }else {
            signupBtn.disabled = true;
        }
    }
});
