function Prompt() {
    // toast 
    let toast = function(c) {
        const {
            msg = "", 
            icon = "success",
            position = "top-end",
        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg, 
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
              toast.addEventListener('mouseenter', Swal.stopTimer)
              toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    // success 
    let success = function(c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;

        Swal.fire({
            icon : 'success',
            title: title,
            icon: icon,
            footer: footer,
        })
    }

    // error 
    let error = function(c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;

        Swal.fire({
            icon : 'error',
            title: title,
            text: msg, 
            footer: footer,
        })
    }

    // custom 
    async function custom(c) {
        const {
            msg = "", 
            title = "",
        } = c;

        const { value: result } = await Swal.fire({
            title: title,
            html: msg,
            // backdrop: true,
            focusConfirm: false,
            showCancelButton: true, 
            // allowOutsideClick: true,
            willOpen: () => {
                const elem = document.getElementById('reservation-dates-modal');
                const rp = new DateRangePicker(elem, {
                    format: 'yyyy-mm-dd',
                    showOnFocus: true,
                })
            },
            preConfirm: () => {
              return [
                document.getElementById('start').value,
                document.getElementById('end').value
              ]
            }, 
            didOpen: () => {
                document.getElementById('start').removeAttribute('disabled');
                document.getElementById('end').removeAttribute('disabled');
            }
        })

        console.log("1. got to result ")
        if (result) {
            // if user didn't hit the cancel button
            if (result.dismiss !== Swal.DismissReason.cancel) {
                // check values in result if result is not empty
                if (result.value !== "") {
                    // if there is a call back
                    if (c.callback !== undefined) {
                        // call back the result (ie what user entered)
                        c.callback(result);
                    } 
                } else {
                    // return a false 
                    c.callback(false);
                }
            } else {
                // else if user canceled 
                c.callback(false);
            }
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}