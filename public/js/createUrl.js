document.addEventListener('DOMContentLoaded', () => {
    const shortenForm = document.getElementById('shortenForm');
    const logo = document.getElementById('nav-logo');
    const shortUrl = document.getElementById('short-url');
    const copyBtn = document.getElementById('copy');
    const modalBox = document.getElementById('modal')


    shortenForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const original_url = shortenForm.longUrl.value;
        const user_id = localStorage.getItem('userId');
        const date = document.getElementById('date');
        const time = document.getElementById('time');
        const expiration_date = new Date(date.value+'T'+time.value).toISOString();

        if (!user_id) {
            window.location.href = 'login.html';
            alert('User not logged in');
            return;
        }
        
        const response = await fetch('/api/urls/create-short-url', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({ original_url, user_id, expiration_date })
        });
        
        if (response.ok) {
            const res = await response.json();
            shortUrl.value = res.data.short_url;
            modalBox.style.display = "block";
            var span = document.getElementsByClassName("close")[0];
            span.onclick = function() {
                modal.style.display = "none";
            }
            window.onclick = function(event) {
                if (event.target == modal) {
                  modal.style.display = "none";
                }
            }

            copyBtn.addEventListener('click', async(e) => {
                copyUrl = shortUrl.select();
                await navigator.clipboard.writeText(copyUrl.value);

            })
        } else {
            alert('Failed to create short URL');
        }
    });

    

    async function loadContent() {
        const user_id = localStorage.getItem('userId');
        if (!user_id) {
            window.location.href = 'login.html';
            alert('User not logged in');
            return;
        }

        const response = await fetch(`/api/users/${user_id}`);
        if(response.ok){
            logo.href = 'create-url.html';
        }else {
            alert(response.json().err)
        }

        let currentDate = new Date();

        let day = currentDate.getDate() + 1;
        let month = currentDate.getMonth() + 1; 
        let year = currentDate.getFullYear();

        if (day < 10) {
            day = '0' + day;
        }
        if (month < 10) {
            month = '0' + month;
        }

        let formattedDate = year + '-' + month + '-' + day;

        document.getElementById('date').value = formattedDate;
        let hours = currentDate.getHours();
        let seconds = currentDate.getSeconds();

        if (hours < 10) {
        hours = '0' + hours;
        }
        if (seconds < 10) {
        seconds = '0' + seconds;
        }
        let formattedTime = hours + ':' + seconds;

        document.getElementById('time').value = formattedTime
    }

    loadContent();

    const logoutButton = document.getElementById('logoutButton');
    logoutButton.addEventListener('click', async () => {
        const response = await fetch('/api/urls/logout', { method: 'GET' });
        if (response.ok) {
            localStorage.clear();
            window.location.href = 'login.html';
        } else {
            alert('Logout failed');
        }
    }, false);
});

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