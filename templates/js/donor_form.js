document.getElementById('donorForm').addEventListener('submit', function(e) {
    e.preventDefault();

    // Get the phone number value and remove non-numeric characters
    const phoneInput = document.getElementById('phone');
    let rawPhoneNumber = phoneInput.value.replace(/\D/g, ''); // Remove formatting (non-numeric characters)
    
    // Add the +90 Turkish country code
    const formattedPhoneNumber = `+90${rawPhoneNumber}`;
    
    // Create a new FormData object to include the raw phone number
    const formData = new FormData(this);
    formData.set('PhoneNumber', formattedPhoneNumber); // Override the PhoneNumber field with the raw number

    // Submit the form data to the server
    fetch('/donor', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            document.getElementById('donorForm').style.display = 'none';
            document.getElementById('verificationSection').style.display = 'block';
            localStorage.setItem('donorId', data.id);
        }
    })
    .catch(error => console.error('Error:', error));
});

document.getElementById('verifyButton').addEventListener('click', function() {
    const verificationCode = document.getElementById('verification_code').value;
    const donorId = localStorage.getItem('donorId');
    fetch(`/verify-donor/${donorId}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ verificationCode })
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            alert('تم التحقق بنجاح');
            window.location.href = '/';
        }
    })
    .catch(error => console.error('Error:', error));
});