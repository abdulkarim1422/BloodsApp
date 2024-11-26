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
