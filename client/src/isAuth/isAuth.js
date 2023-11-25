import Cookies from 'js-cookie'
export function IsAuth() {
    let cookieValue = Cookies.get("token")
    return cookieValue !== undefined;
}