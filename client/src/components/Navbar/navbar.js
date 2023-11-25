import "./navbar.css"
import {IsAuth} from "../../isAuth/isAuth";
import Cookies from "js-cookie";
import {redirect} from "react-router-dom";

const Logout = () => {
    const cookies = document.cookie.split(";");

    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i];
        const eqPos = cookie.indexOf("=");
        const name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
    }
}

export const NavBar = () => {
    if (IsAuth()) {
        return (
            <div className="navbar">
                <div className="logo">
                    <a href="/" className="logo">NNapp</a>
                </div>
                <div className="links">
                    <a href="" className="link" onClick={Logout}>Logout</a>
                </div>
            </div>
        );
    }
    return (
        <div className="navbar">
            <div className="logo">
                <a href="/" className="logo">NNapp</a>
            </div>
            <div className="links">
                <a href="https://github.com/0azis/" target="_blank" className="link">GitHub</a>
                <a href="" className="link">About</a>
            </div>
        </div>
    );
}