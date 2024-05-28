// htmx.logAll();
document.addEventListener("DOMContentLoaded", (event) => {
	document.addEventListener("htmx:beforeSwap", (evt) => {
		if (evt.detail.xhr.status == 422) {
			evt.detail.shouldSwap = true;
			evt.detail.isError = false;
		}
	});
	document.addEventListener("htmx:afterSwap", (evt) => {
		if (evt.detail.elt.id == "form-container") {
			const digit = parseInt(
				document.getElementById("form-container").getAttribute("data-digit")
			);
			addFields(digit);
        }
    });

    // TODO: there is a werid problem where this request's elt & target are both tbody. See if i can reproduce it!
    document.addEventListener("htmx:afterSettle", (evt) => {
        if (evt.detail.target.id == "tbody") { 
            const digit = document.getElementById("form-container").getAttribute("data-digit");
            const ans = document.getElementById("tbody").rows[0].cells[2].innerText[0];
            console.log(digit, ans);
            if (digit == ans) {
                Swal.fire({
                    title: "You won the game!",                     
                    confirmButtonText: "Return Home",
                    showCancelButton: true,
                    text: `The answer is ${ans}`,
                }).then( (result) => {
                    if (result.isConfirmed) {
                        const btn = document.getElementById("return-btn");
                        htmx.trigger(btn, "return");
                    }
                });
            }
		}
	});
	document.addEventListener("htmx:configRequest", (evt) => {
		if (evt.detail.elt.id == "submit-btn") {
			evt.detail.parameters["guess"] = calcInputs(); // add a new param to request
			var inputs = document.getElementsByClassName("digit-input"); // clear input
			Array.from(inputs).forEach((i) => {
				i.value = "";
			});
			const box = document.getElementById("digit1");
			box.focus();
		}
	});
});

// Dynamically generate input fields for game.html
function addFields(digit) {
	var parent = document.createElement("div");
	parent.id = "input-container";
	parent.className = "input-container";
	var form = document.getElementById("form-container");
	var container = form.insertBefore(parent, form.childNodes[4]);

	for (i = 0; i < digit; i++) {
		var input = document.createElement("input");
		input.type = "number";
		input.min = "0";
		input.max = "9";
		input.name = "digit" + (i + 1);
		input.id = "digit" + (i + 1);
		input.className = "digit-input";
		container.appendChild(input);
	}

	// Allow only single digit inputs
	var boxes = document.getElementsByClassName("digit-input");
	var secondPress = false;
	Array.from(boxes).forEach((box, index, array) => {
		box.addEventListener("input", (event) => {
			if (box.value.length == 1 && index != array.length - 1) {
				// Focus on the next sibling
				box.nextElementSibling.focus();
			}
		});
		box.addEventListener("keyup", (event) => {
			if (
				event.key == "Backspace" &&
				!box.value &&
				index > 0 &&
				secondPress
			) {
				// Focus on the previous sibling
				box.previousElementSibling.focus();
				secondPress = false;
			} else if (event.key == "Backspace" && !box.value && index > 0) {
				secondPress = true;
			} else if (
				event.key == "Enter" &&
				box.value &&
				index == array.length - 1
			) {
				document.getElementById("submit-btn").focus();
			}
		});
	});
}

// Concat strings from input boxes
function calcInputs() {
	var boxes = document.getElementsByClassName("digit-input");
	var val = "";
	Array.from(boxes).forEach((box) => {
		if (box.value) {
			val += box.value;
		}
	});
	return val;
}
