document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');
    const loginBtn = document.getElementById('login-btn');

    loginBtn.disabled = true;

    console.log(loginForm.loginUsername.value);
    console.log(loginForm.loginPassword.value);

    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                identity: loginForm.loginUsername.value,
                password: loginForm.loginPassword.value
            })
        });
        if (response.ok) {
            const res = await response.json();
            localStorage.setItem('userId', res.data.user.id);
            window.location.href = 'create-url.html';
        } else {
            alert('Login failed');
        }
    });

    var passwordInput = document.getElementById("loginPassword"); 
    var passwordMessageItems = document.getElementsByClassName("password-message-item"); 
    var passwordMessage = document.getElementById("password-message"); 
    var lenght = 0;
    
    passwordInput.onfocus = function () { 
        passwordMessage.style.display = "block"; 
    } 
    passwordInput.onblur = function () { 
        passwordMessage.style.display = "none"; 
    } 
    passwordInput.onkeyup = function () {
        if (passwordInput.value.length >= 8) { 
            passwordMessageItems[0].classList.remove("invalid"); 
            passwordMessageItems[0].classList.add("valid"); 
            lenght = true;
        } else { 
            passwordMessageItems[0].classList.remove("valid"); 
            passwordMessageItems[0].classList.add("invalid"); 
            lenght = false;
        } 
        if(lenght){
            loginBtn.disabled = false;
        }else {
            loginBtn.disabled = true;
        }
    }
});
