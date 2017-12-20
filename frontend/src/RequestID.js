import rp from 'require-promise';

function getPost(id, hosturi) {
    var options = {
        uri: hosturi+'/request/post/'+id,
        headers: {
            'User-Agent': 'Request-Promise'
        },
        json: true  
    };
    rp(options)
        .then(function (content) {
            return content.length;
        })
        .catch(function(err){
            return "something went wrong";
        });
}
