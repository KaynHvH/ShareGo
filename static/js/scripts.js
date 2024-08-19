document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById('uploadForm');
    const statusDiv = document.getElementById('status');

    form.addEventListener('submit', function(event) {
        event.preventDefault();

        const fileInput = document.getElementById('fileInput');
        const file = fileInput.files[0];

        if (file) {
            const formData = new FormData();
            formData.append('file', file);
            formData.append('user_id', '1');

            fetch('/upload', {
                method: 'POST',
                body: formData
            })
                .then(response => response.json())
                .then(data => {
                    if (data.success) {
                        statusDiv.innerHTML = `<p>File uploaded successfully: ${data.filename}</p>`;
                    } else {
                        statusDiv.innerHTML = `<p>Error: ${data.error}</p>`;
                    }
                })
                .catch(error => {
                    statusDiv.innerHTML = `<p>Error: ${error.message}</p>`;
                });
        } else {
            statusDiv.innerHTML = `<p>No file selected.</p>`;
        }
    });

    const showFilesButton = document.getElementById('showFilesButton');
    showFilesButton.addEventListener('click', function() {
        window.location.href = '/files.html';
    });
});
