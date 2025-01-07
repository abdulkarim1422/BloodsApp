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

// Verify the patient's phone number
document.getElementById('verifyButton').addEventListener('click', function() {
    const verificationCode = document.getElementById('verification_code').value;
    const patient_id = localStorage.getItem('patient_id');
    fetch(`/verify-patient/${patient_id}`, {
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

// Toggle the display of the Red Crescent code field
function toggleRedCrescentCode() {
    var selectBox = document.getElementById("acceptsRedCrescent");
    var redCrescentCodeDiv = document.getElementById("redCrescentCodeDiv");
    var requiresNationalityDiv = document.getElementById("requiresNationalityDiv");
    if (selectBox.value === "true") {
        redCrescentCodeDiv.style.display = "block";
        requiresNationalityDiv.style.display = "none";
    } else {
        redCrescentCodeDiv.style.display = "none";
        requiresNationalityDiv.style.display = "block";
    }
}