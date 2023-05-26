function encodeImageFileAsURL(file) {
    var reader = new FileReader();
    reader.onloadend = function () {
        console.log('Encoded image:', reader.result);
    }
    reader.readAsDataURL(file);
}


function postRequest(url,Body) {
    fetch(url, {
        method: "POST",
        body: Body,
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    })
        .then((response) => response.json())
        .then((json) => {
            let data=JSON.parse(json);
            return data;
        });
}



