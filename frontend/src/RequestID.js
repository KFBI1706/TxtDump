var rp = require('request-promise');

module.exports = {
    getPost: function (id, hosturi, result) {
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
                result(htmlString);
            })
            .catch(function(err){
                return "something went wrong";
            });
    }
}