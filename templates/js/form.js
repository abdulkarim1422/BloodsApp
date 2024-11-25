// Format the phone number input as the user types
document.getElementById('phone').addEventListener('input', function () {
    let input = this.value.replace(/\D/g, ''); // Remove non-numeric characters
    input = input.substring(0, 10); // Limit to 10 digits for Turkish numbers

    // Format the input as (XXX) XXX XX XX
    const formatted = input.replace(
        /(\d{3})(\d{3})(\d{2})(\d{2})/,
        '($1) $2 $3 $4'
    );

    // Update the input value with the formatted version
    this.value = formatted;
});

// When the form is submitted, prepare the phone number for submission
document.getElementById('patientForm').addEventListener('submit', function(e) {
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
    fetch('/patient', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            document.getElementById('patientForm').style.display = 'none';
            document.getElementById('verificationSection').style.display = 'block';
            localStorage.setItem('patient_id', data.id);
        }
    })
    .catch(error => console.error('Error:', error));
});

