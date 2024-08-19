document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById('registerForm');
    const statusDiv = document.getElementById('status');

    form.addEventListener('submit', function(event) {
        event.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        fetch('/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
            .then(response => {
                if (response.ok) {
                    statusDiv.innerHTML = `<p>Registration successful! <a href="/login.html">Login here</a></p>`;
                } else {
                    return response.text().then(text => {
                        statusDiv.innerHTML = `<p>Error: ${text}</p>`;
                    });
                }
            })
            .catch(error => {
                statusDiv.innerHTML = `<p>Error: ${error.message}</p>`;
            });
    });
});
