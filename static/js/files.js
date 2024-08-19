document.addEventListener("DOMContentLoaded", function() {
    fetch('/files', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(files => {
            const fileList = document.getElementById("file-list");
            fileList.innerHTML = "";

            if (files.length > 0) {
                files.forEach(file => {
                    const listItem = document.createElement("li");
                    
                    if (isImage(file.filename)) {
                        const img = document.createElement("img");
                        img.src = `/files/${file.id}/download`;
                        img.alt = file.filename;
                        img.classList.add("file-preview");
                        listItem.appendChild(img);
                    } else if (isPdf(file.filename)) {
                        const iframe = document.createElement("iframe");
                        iframe.src = `/files/${file.id}/download`;
                        iframe.style.width = "100%";
                        iframe.style.height = "500px";
                        iframe.frameBorder = 0;
                        listItem.appendChild(iframe);
                    } else {
                        listItem.textContent = file.filename;
                    }

                    const downloadLink = document.createElement("a");
                    downloadLink.href = `/files/${file.id}/download`;
                    downloadLink.textContent = "Download";
                    downloadLink.classList.add("download-button");

                    listItem.appendChild(downloadLink);
                    fileList.appendChild(listItem);
                });
            } else {
                fileList.innerHTML = "<li>No files found.</li>";
            }
        })
        .catch(error => {
            console.error("Error loading files:", error);
        });
});

function isImage(filename) {
    const extension = filename.split('.').pop().toLowerCase();
    return ['jpg', 'jpeg', 'png', 'gif'].includes(extension);
}

function isPdf(filename) {
    return filename.split('.').pop().toLowerCase() === 'pdf';
}
