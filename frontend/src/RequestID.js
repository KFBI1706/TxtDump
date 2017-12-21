var rp = require('request-promise');

function getPost(id, hosturi) {
    var options = {
        uri: hosturi+'/request/post/'+id,
        headers: {
            'User-Agent': 'Request-Promise'
        },
        json: true  
    };
    rp(options)
        .then(function (htmlString) {
            console.log(htmlString)
        })
        .catch(function(err){
            return "something went wrong";
        });
}

