document.addEventListener('DOMContentLoaded', () => {
    const name = document.getElementById('name');
    const username = document.getElementById('username');
    const email = document.getElementById('email');
    const totalURLs = document.getElementById('total-links');
    const logo = document.getElementById('nav-logo');

    async function loadContent() {
        const user_id = localStorage.getItem('userId');
        if (!user_id) {
            window.location.href = 'login.html';
            alert('User not logged in');
            return;
        }

        const response = await fetch(`/api/users/${user_id}`);
        if(response.ok){
            const res = await response.json();
            name.value = res.data.user.name;
            username.value = res.data.user.username;
            email.value = res.data.user.email;
            totalURLs.value = res.data.user.total_urls;
            logo.href = 'create-url.html';
        }else {
            alert(response.json().err)
        }
    }

    loadContent();
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
