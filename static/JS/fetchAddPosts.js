const AddPosts = () => {
    fetch("/tryAddPosts", {
        method: "POST",
        header: {
            "content-type": "application/json"
        },
        body: JSON.stringify({
            title: document.getElementById("title").value,
            theme: document.getElementById("theme").value,
            content: document.getElementById("content").value
        })
    })
    var value = document.getElementById("content");
    console.log(value);

}
