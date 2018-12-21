window.onload = init;
function activateCustomPass(checkbox = "customPass", element = "customField") {
    var customPass = document.getElementById(checkbox).checked;
    if (customPass == true) {
        document.getElementById(element).disabled = false;
    } else {
        document.getElementById(element).disabled = true;
    }
}

function init(){
	debug = document.getElementById("debug");
    if (debug) {
	    consolle.log(debug.innerHTML);
	}
}

function modalIfNoPass() {
	var customPass = document.getElementById("customPass").checked;
	var passField = document.getElementById("customField").value;

    if (customPass === true) {
        document.getElementById("postForm").submit();
    } else {
		if(document.getElementById("sel").value != 1 ){
			$('#noPassModal').modal(document)
		}
		else {
			document.getElementById("postForm").submit();
		}
    }
}

if(window.location.pathname == "/post/create")  {
    document.getElementById("newpost").classList.add('active');
}
else if(window.location.pathname == "/documentation"){
    document.getElementById("doc").classList.add('active');
}
$(document).ready(function () {
    var userInputs = new Array();
    userInputs.push("john.smith@test.com");
    $("#customField").zxcvbnProgressBar({
          passwordInput: "#customField",
          userInputs: userInputs,
    });
});

(function ($) {

	$.fn.zxcvbnProgressBar = function (options) {
		var settings = $.extend({
			passwordInput: '#Password',
			userInputs: [],
		}, options);

		return this.each(function () {
			UpdateProgressBar();
			$(settings.passwordInput).keyup(function (event) {
				UpdateProgressBar();
			});
		});

		function UpdateProgressBar() {
			var field = document.getElementById("customField");
			var password = $(settings.passwordInput).val();
			if (password) {
				var result = zxcvbn(password, settings.userInputs);

				if (result.score == 0) {
					//weak
					field.style.backgroundColor = "#dc3545";
					field.style.color = "#fff";
				}
				else if (result.score == 1) {
					//normal
					field.style.backgroundColor = "#ffc107";
					field.style.color = "#fff";
				}
				else if (result.score == 2) {
					//medium
					field.style.backgroundColor = "#17a2b8";
					field.style.color = "#fff";
				}
				else if (result.score == 3) {
					//strong
					field.style.backgroundColor = "#28a745";
					field.style.color = "#fff";
				}
				else if (result.score == 4) {
					//very strong
					field.style.backgroundColor = "#007bff";
					field.style.color = "#fff";
				}
			}
			else {
				field.style.backgroundColor = "#fff";
				field.style.color = "#333";
			}
		}
	};
})(jQuery);