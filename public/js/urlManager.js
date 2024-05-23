document.addEventListener('DOMContentLoaded', () => {
    const shortenForm = document.getElementById('shortenForm');
    const logoutButton = document.getElementById('logoutButton');
    const urlList = document.getElementById('urlList');

    shortenForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const original_url = shortenForm.longUrl.value;
        const user_id = localStorage.getItem('userId');
        const expiration_date = new Date(shortenForm.expirationDate.value).toISOString();

        if (!user_id) {
            window.location.href = 'login.html';
            alert('User not logged in');
            return;
        }
        
        const response = await fetch('/api/url/create-short-url', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({ original_url, user_id, expiration_date })
        });
        
        if (response.ok) {
            loadURLs();
        } else {
            alert('Failed to create short URL');
        }
    });

    logoutButton.addEventListener('click', async () => {
        const response = await fetch('/api/url/logout', { method: 'GET' });
        if (response.ok) {
            localStorage.clear();
            window.location.href = 'login.html';
        } else {
            alert('Logout failed');
        }
    });

    async function loadURLs() {
        const response = await fetch('/api/url/get-urls', { method: 'GET' });
        if (response.ok) {
            const urls = await response.json();
            console.log(urls);
            urlList.innerHTML = '';
            urls.data.forEach(url => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td><a>${url.short_url}</a></td>
                    <td><a>${url.original_url}</a></td>
                    <td>${url.visiter}</td>
                `;
                urlList.appendChild(tr);
            });
        } else {
            window.location.href = 'login.html';
            alert('Failed to load URLs');
        }
    }

    loadURLs();
});
