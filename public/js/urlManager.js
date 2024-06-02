document.addEventListener('DOMContentLoaded', () => {
    const shortenForm = document.getElementById('shortenForm');
    const urlList = document.getElementById('urlList');
    const logo = document.getElementById('nav-logo');


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
            // add a copy link
        } else {
            alert('Failed to create short URL');
        }
    });

    

    async function loadContent() {
        logo.href = 'create-url.html';
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