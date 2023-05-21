function encodeImageFileAsURL(file) {
    var reader = new FileReader();
    reader.onloadend = function() {
        console.log('Encoded image:', reader.result);
    }
    reader.readAsDataURL(file);
}

// Usage example:
var fileInput = document.querySelector('input[type="file"]');
fileInput.addEventListener('change', function() {
    var file = fileInput.files[0];
    encodeImageFileAsURL(file);
});





