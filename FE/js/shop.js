const somererular=require("./main.js")

function loadProduct() {
    
    let itemcontain = document.getElementById("itemPanel");
    let itemtempl = document.getElementsByTagName("template")[0];
    console.log(itemcontain)
    console.log(itemtempl)
    let dataarray = postRequest("/Rproduct", "{'product':'all'}");
    for (let i = 0; i < dataarray.length; i++) {
        let node = itemtempl.content.querySelector("div");
        console.log(node)

        let tensptxt = node.content.querySelector("h3");
        console.log(tensptxt)
        tensptxt.content = dataarray[i].tensp;
        let icon = node.content.querySelector("img");
        icon.src = dataarray[i].iconencode;
        console.log(icon)

        itemcontain.appendChild(node);
    }
}
