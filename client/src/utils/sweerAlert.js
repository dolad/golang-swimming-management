
import Swal from 'sweetalert2'

export const successAlertMesssage = () => {
    return Swal.fire({
        position: 'top-end',
        icon: 'success',
        title: 'Successfully loggedIn',
        showConfirmButton: false,
        timer: 1500
      })    
};

export const errorMessageAlertMessage = (message) => {
    return Swal.fire({
        icon: 'error',
        title: 'Oops...',
        text: `${message}`,
      })
}