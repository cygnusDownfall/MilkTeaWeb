function encodeImageFileAsURL(file) {
    var reader = new FileReader();
    reader.onloadend = function () {
        console.log('Encoded image:', reader.result);
    }
    reader.readAsDataURL(file);
}

async function postRequest(url, Body) {
    const res = await fetch(url, {
        method: "POST",
        body: JSON.stringify(Body),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    });
    return res.json()
}

function autoIncrease(previousID, increaseAmount) {
    let previousValue = parseInt(previousID.slice(2));
    let increasedValue = previousValue + increaseAmount;
    return previousID.slice(0, 2) + increasedValue.toString().padStart(3, '0');
}

function hashCode(str) {
    var hash = 0;
    for (var i = 0; i < str.length; i++) {
        hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    return hash;
}


function move(movingDiv,targetid) {
    let targetDiv = document.getElementById(targetid);
    console.log("tagDiv"+targetDiv,movingDiv);
    let targetPosition = targetDiv.getBoundingClientRect();

    movingDiv.style.position = "absolute";
    movingDiv.style.left = targetPosition.left + "px";
    movingDiv.style.top = targetPosition.top + "px";

}