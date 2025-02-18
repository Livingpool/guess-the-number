// dynamically generate input fields for game.html
function addFields(digit) {
    var parent = document.createElement('div');
    parent.id = 'input-container';
    parent.className = 'input-container';
    var form = document.getElementById('form-container');
    var container = form.insertBefore(parent, form.childNodes[4]);

    for (i = 0; i < digit; i++) {
        var input = document.createElement('input');
        input.type = 'number';
        input.min = '0';
        input.max = '9';
        input.name = 'digit' + (i + 1);
        input.id = 'digit' + (i + 1);
        input.className = 'digit-input';
        container.appendChild(input);
    }

    var boxes = document.getElementsByClassName('digit-input');
    var secondPress = false;
    document.getElementById('digit1').focus();

    // allow only single digit inputs
    Array.from(boxes).forEach((box, index, array) => {
        box.addEventListener('input', (event) => {
            if (box.value.length == 1 && index != array.length - 1) {
                // focus on the next sibling
                box.nextElementSibling.focus();
            }
        });
        box.addEventListener('keyup', (event) => {
            if (
                event.key == 'Backspace' &&
                !box.value &&
                index > 0 &&
                secondPress
            ) {
                // focus on the previous sibling
                box.previousElementSibling.focus();
                secondPress = false;
            } else if (event.key == 'Backspace' && !box.value && index > 0) {
                secondPress = true;
            } else if (
                event.key == 'Enter' &&
                box.value &&
                index == array.length - 1
            ) {
                document.getElementById('submit-btn').focus();
            }
        });
    });
}

// concat strings from input boxes
function calcInputs() {
    var boxes = document.getElementsByClassName('digit-input');
    var val = '';
    Array.from(boxes).forEach((box) => {
        if (box.value) {
            val += box.value;
        }
    });
    return val;
}
