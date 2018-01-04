axios.get('http://localhost:1337/post/9175728/request')
  .then(function (response) {
    console.log(response.data);
    changeTest(response.data);
  })
  .catch(function (error) {
    console.log(error);
    changeTest();
  });
  

function changeTest(text = "something went wrong"){
   var trg = document.getElementsByClassName("postText")[0];
   trg.innerHTML = text.Content + text.Time;
}
