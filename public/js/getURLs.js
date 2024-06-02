document.addEventListener('DOMContentLoaded', () => {
    const logoutButton = document.getElementById('logoutButton');
    const logo = document.getElementById('nav-logo');
    logoutButton.addEventListener('click', async () => {
        const response = await fetch('/api/urls/logout', { method: 'GET' });
        if (response.ok) {
            localStorage.clear();
            window.location.href = 'login.html';
        } else {
            alert('Logout failed');
        }
    }, false);

    async function loadURLs() {
        const response = await fetch('/api/urls/get-urls', { method: 'GET' });
        if (response.ok) {
            logo.href = 'create-url.html';
            const urls = await response.json();
            console.log(urls);
            urlList.innerHTML = '';
            let i = 0
            urls.data.forEach(url => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${++i}</td>
                    <td><a>${url.short_url}</a></td>
                    <td><a>${url.original_url}</a></td>
                    <td>${url.visiter}</td>
                    <td>${url.created_at}</td>
                    <td>${url.validity}</td>
                `;
                urlList.appendChild(tr);
            });
        } else {
            window.location.href = 'login.html';
            alert('Failed to load URLs');
        }
    }

    loadURLs();
})

function myFunction() {
    document.getElementById("myDropdown").classList.toggle("show");
  }

window.onclick = function(event) {
if (!event.target.matches('.dropbtn')) {
    var dropdowns = document.getElementsByClassName("dropdown-content");
    var i;
    for (i = 0; i < dropdowns.length; i++) {
    var openDropdown = dropdowns[i];
    if (openDropdown.classList.contains('show')) {
        openDropdown.classList.remove('show');
    }
    }
}
}