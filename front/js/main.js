function activateCustomPass() {
    var customPass = document.getElementById("customPass").checked;
    if (customPass == true) {
        document.getElementById('customField').disabled = false;
    } else {
        document.getElementById('customField').disabled = true;
    }
}

function modalIfNoPass() {
    var customPass = document.getElementById("customPass").checked;
    if (customPass == true) {
        document.getElementById("postForm").submit();
    } else {
        $('#noPassModal').modal(document)
    }
}

if(window.location.pathname == "/post/create")  {
    document.getElementById("newpost").classList.add('active');
}
else if(window.location.pathname == "/documentation"){
    document.getElementById("doc").classList.add('active');
}