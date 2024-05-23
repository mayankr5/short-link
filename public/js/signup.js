document.addEventListener('DOMContentLoaded', () => {
    const signupForm = document.getElementById('signupForm');

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
});
