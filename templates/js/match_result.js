document.addEventListener('DOMContentLoaded', function() {
    const sendButton = document.getElementById('sendButton');
    sendButton.addEventListener('click', function() {
        console.log('Send button clicked'); // Debugging log
        sendButton.disabled = true; // Disable the button
        sendButton.style.backgroundColor = 'gray'; // Change button color to gray
        sendButton.style.cursor = 'not-allowed'; // Change cursor to not-allowed
        sendButton.innerText = 'تمّ تشغيل برنامج إرسال الرّسائل بنجاح'; // Change button text

        const selectedDonors = Array.from(document.querySelectorAll('.donor-checkbox:checked')).map(checkbox => {
            const row = checkbox.closest('tr');
            const cells = row.querySelectorAll('td');
            return parseInt(cells[1].innerText, 10); // Extract and parse the ID as an integer
        });

        if (selectedDonors.length === 0) {
            alert('Please select at least one donor.');
            sendButton.disabled = false; // Re-enable the button
            sendButton.style.backgroundColor = ''; // Reset button color
            sendButton.style.cursor = ''; // Reset cursor
            sendButton.innerText = 'Send Match Result'; // Reset button text
            return;
        }

        const matchSection = document.getElementById('matchSection');
        const message = matchSection.getAttribute('data-message');
        const patientId = matchSection.getAttribute('data-patient-id');

        const jsonResponse = {
            message: message,
            patient_id: parseInt(patientId),
            donors: selectedDonors
        };

        console.log('Sending JSON:', JSON.stringify(jsonResponse)); // Log the JSON data

        // Calculate the estimated time
        const estimatedTime = 20 + (18 * selectedDonors.length);
        document.getElementById('sendResult').innerText = `يتمّ التحميل الآن... الوقت التخميني : ${estimatedTime} ثانية`;
        document.getElementById('sendResult').innerText += '\nيرجى عدم إرسال أمر آخر أثناء عمل البرنامج';

        // Show and animate the progress bar
        const progressBarContainer = document.getElementById('progressBarContainer');
        const progressBar = document.getElementById('progressBar');
        progressBarContainer.style.display = 'block';
        let width = 0;
        const interval = setInterval(() => {
            if (width >= 100) {
                clearInterval(interval);
            } else {
                width++;
                progressBar.style.width = width + '%';
                progressBar.innerText = width + '%';
            }
        }, (estimatedTime * 10)); // Adjust the interval time to match the estimated time

        // Send the JSON data to the server
        fetch('/send-match-result', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonResponse)
        })
        .then(response => {
            console.log('Response received'); // Debugging log
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        }) // Parse the JSON response
        .then(data => {
            console.log('Data received:', data); // Debugging log
            document.getElementById('sendResult').innerText = data.message;
            clearInterval(interval); // Stop the progress bar animation
            progressBar.style.width = '100%'; // Set progress bar to 100%
            progressBar.innerText = '100%';
        })
        .catch(error => {
            console.error('Error:', error);
            document.getElementById('sendResult').innerText = 'Error: ' + error.message;
            clearInterval(interval); // Stop the progress bar animation
            progressBar.style.width = '100%'; // Set progress bar to 100%
            progressBar.style.backgroundColor = 'red'; // Change progress bar color to red
            progressBar.innerText = 'Error';
        });
    });
});


// Select all and deselect all buttons
document.getElementById('selectAllButton').addEventListener('click', function() {
    document.querySelectorAll('.donor-checkbox').forEach(checkbox => {
        checkbox.checked = true;
    });
});

document.getElementById('deselectAllButton').addEventListener('click', function() {
    document.querySelectorAll('.donor-checkbox').forEach(checkbox => {
        checkbox.checked = false;
    });
});