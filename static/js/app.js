function Prompt(){
    let toast = function (e){
        const {
            msg = "",
            icon = "success",
            position = "top-end",
        } = e;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position:position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.onmouseenter = Swal.stopTimer;
                toast.onmouseleave = Swal.resumeTimer;
            }
        });
        Toast.fire({});
    }
    let success = function (e){
        const {
            msg = "",
            title = "",
            footer = "",
        } = e
        Swal.fire({
            icon: "success",
            title: title,
            text: msg,
            footer: footer,
        });
    }

    let error = function (e){
        const {
            msg = "",
            title = "",
            footer = "",
        } = e
        Swal.fire({
            icon: "error",
            title: title,
            text: msg,
            footer: footer,
        });
    }

    async function custom(c) {
        const {
            icon = "",
            msg = "",
            title = "",
            showConfirmButton = true,
        } = c;

        const { value: result } = await Swal.fire({
            icon: icon,
            title:title,
            html:msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton: showConfirmButton,
            willOpen:() =>{
                if(c.willOpen !== undefined){
                    c.willOpen()
                }
            },
            didOpen: () => {
                if(c.didOpen !== undefined){
                    c.didOpen()
                }
            },
        });
        if (result) {
            if(result.dismiss !== Swal.DismissReason.cancel){
                if(result.value !== ""){
                    if(c.callback !== undefined){
                        c.callback(result);
                    }
                }else{
                    c.callback(false);
                }
            }else{
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