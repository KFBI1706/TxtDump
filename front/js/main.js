function activateCustomPass(){
    var customPass = document.getElementById("customPass").checked;
    if(customPass == true){
        document.getElementById('customField').disabled = false;
    }
    else{
        document.getElementById('customField').disabled = true;
    }
}