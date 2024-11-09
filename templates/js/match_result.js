document.getElementById('sendButton').addEventListener('click', function() {
    const selectedDonors = Array.from(document.querySelectorAll('.donor-checkbox:checked')).map(checkbox => {
        const row = checkbox.closest('tr');
        const cells = row.querySelectorAll('td');
        return {
            id: cells[1].innerText,
            first_name: cells[2].innerText,
            last_name: cells[3].innerText,
            bloodType: cells[4].innerText,
            phone_number: cells[5].innerText,
            address: cells[6].innerText,
            email: cells[7].innerText,
            age: cells[8].innerText,
            gender: cells[9].innerText
        };
    });

    const jsonResponse = {
        message: document.getElementById('message').innerText,
        patient_first_name: document.getElementById('patient_first_name').innerText,
        patient_last_name: document.getElementById('patient_last_name').innerText,
        patient_blood_type: document.getElementById('patient_blood_type').innerText,
        donors: selectedDonors
    };

    console.log('Sending JSON:', JSON.stringify(jsonResponse)); // Log the JSON data

    fetch('/send', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(jsonResponse)
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('sendResult').innerText = data.message;
    })
    .catch(error => {
        document.getElementById('sendResult').innerText = 'Error: ' + error;
    });
});