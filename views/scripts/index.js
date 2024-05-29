// htmx.logAll();
document.addEventListener("DOMContentLoaded", (event) => {
	document.addEventListener("htmx:beforeSwap", (evt) => {
		if (evt.detail.xhr.status == 422) {
			evt.detail.shouldSwap = true;
			evt.detail.isError = false;
            if (evt.detail.target.id == "tbody") {
                Swal.fire({
                    title: "Invalid input length",
                    icon: "warning",
                    // timer: 10000,
                    confirmButtonText: "OK",
                })
            }
		} else if (evt.detail.xhr.status == 404) { // handle player id not found
            evt.detail.isError = false;
            Swal.fire({
                title: "Player not found",
                icon: "error",
                // timer: 10000,
                confirmButtonText: "Return Home",
                allowOutsideClick: false,
                allowEscapeKey: false,
            }).then( () => {
                const btn = document.getElementById("return-btn");
                htmx.trigger(btn, "return");
            })
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
    // Handle winning condition
    document.addEventListener("htmx:afterSettle", (evt) => {
        if (evt.detail.target.id == "tbody" && evt.detail.xhr.status != 422) { 
            const digit = document.getElementById("form-container").getAttribute("data-digit");
            const matchDigit = document.getElementById("tbody").rows.length ? 
                document.getElementById("tbody").rows[0].cells[2].textContent[0] : null;
            const ans = document.getElementById("tbody").rows.length ? 
                document.getElementById("tbody").rows[0].cells[1].textContent.split(" ")[1] : null;
            if (matchDigit && matchDigit == digit) {
                Swal.fire({
                    title: "You won the game!",                     
                    icon: "success",
                    // timer: 10000,
                    confirmButtonText: "Return Home",
                    showCancelButton: true,
                    text: `The answer is ${ans}`,
                }).then( (result) => {
                    if (result.isConfirmed || result.dismiss == Swal.DismissReason.timer) {
                        const btn = document.getElementById("return-btn");
                        htmx.trigger(btn, "return");
                    }
                });
            }
		}
	});
 
    // Add param to request POST /guess
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
