document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');

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
            window.location.href = 'urlManager.html';
        } else {
            alert('Login failed');
        }
    });
});
