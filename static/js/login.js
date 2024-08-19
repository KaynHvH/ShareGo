document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById('loginForm');
    const statusDiv = document.getElementById('status');

    form.addEventListener('submit', function(event) {
        event.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else {
                    return response.text().then(text => {
                        statusDiv.innerHTML = `<p>Error: ${text}</p>`;
                    });
                }
            })
            .then(data => {
                if (data.token) {
                    localStorage.setItem('token', data.token);
                    window.location.href = '/uploadfiles';
                }
            })
            .catch(error => {
                statusDiv.innerHTML = `<p>Error: ${error.message}</p>`;
            });
    });
});
