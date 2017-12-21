var rp = require('request-promise');

//SUPER FREAKING JANKY WAY OF DOING THIS, but i really fucking hate javascript so super fucking janky will have to work for now

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