const phoneInput = document.getElementById('phone');

phoneInput.addEventListener('input', function (event) {
    let input = phoneInput.value.replace(/\D/g, ''); // Remove non-numeric characters
    input = input.substring(0, 10); // Limit to 10 digits

    // Format the input as (XXX) XXX XX XX
    const formatted = input.replace(
        /(\d{3})(\d{3})(\d{2})(\d{2})/,
        '($1) $2 $3 $4'
    );

    // Update the input value with the formatted version
    phoneInput.value = formatted;
});
