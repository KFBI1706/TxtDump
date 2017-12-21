var rp = require('request-promise');

//SUPER FREAKING JANKY WAY OF DOING THIS, but i really fucking hate javascript so super fucking janky will have to work for now

module.exports = {
    getPost: function (id, hosturi) {
        var options = {
            uri: hosturi+'/request/post/'+id,
            headers: {
                'User-Agent': 'Request-Promise'
            },
            json: true  
        };
        return rp(options)
            .then(function (htmlString) {
                console.log(htmlString)
                return htmlString;
            })
            .catch(function(err){
                return "something went wrong";
            });
    }
}

