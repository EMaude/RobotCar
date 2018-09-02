function sendCommand(api){
    $.ajax({
        url: "/api" + api
      }).done(function(resp) {
        console.log(resp);
        var temp = resp;
        return temp;
      });
}

$(function(){
    
});