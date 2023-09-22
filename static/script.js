document.addEventListener("DOMContentLoaded", function () {
    const busForm = document.getElementById("busForm");
    const busResults = document.getElementById("busResults");

    busForm.addEventListener("submit", function (e) {
        e.preventDefault();

        const source = document.getElementById("source").value;
        const destination = document.getElementById("destination").value;

        fetch(`/getBuses?source=${source}&destination=${destination}`)
            .then((response) => response.json())
            .then((data) => {
                let html = "<h2>Bus Routes:</h2>";
                data.forEach((bus) => {
                    html += `
                        <p>Route: ${bus.route}</p>
                        <p>Source: ${bus.source}</p>
                        <p>Destination: ${bus.destination}</p>
                        <hr>
                    `;
                });
                busResults.innerHTML = html;
            })
            .catch((error) => {
                console.error("Error fetching data:", error);
                busResults.innerHTML = "Error fetching data. Please try again.";
            });
    });
});
