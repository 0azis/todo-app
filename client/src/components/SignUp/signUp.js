import {IsAuth} from "../../isAuth/isAuth";
import {Navigate, useNavigate} from "react-router-dom";
import {useState} from "react";
import "./signUp.css"

export const SignUp = () => {
    if (IsAuth()) {
        return <Navigate to='/notes' />
    }
    // eslint-disable-next-line react-hooks/rules-of-hooks
    const [values, setValues] = useState({email: '', password: '' })
    // eslint-disable-next-line react-hooks/rules-of-hooks
    const [error, setError] = useState(false)
    // eslint-disable-next-line react-hooks/rules-of-hooks
    const [isLoading, setLoading] = useState(false)
    const submitHandler = e => {
        e.preventDefault()
        setLoading(true)
        fetch('https://notes-mzmm.onrender.com/api/auth/signup', {
            method: 'POST',
            credentials: 'include',
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify(values),
        }).then((response) => {
            if (response.ok) {
                setError(true)
            }
            setLoading(false)
        })
    }


    if (isLoading) return <h1>Регистрация аккаунта... </h1>
    if (error) return <Navigate to="/notes" />
    return (
        <div className="container">
            <h1 className="auth">Регистрация</h1>
            <form method="post" onSubmit={submitHandler}>
                <div className="inputs">
                    <input className="input" type="email" placeholder="Почта" required value={values.email} onChange={e => setValues({ ...values, email: e.target.value })} />
                    <input className="input" type="password" placeholder="Пароль" required value={values.password} onChange={e => setValues({ ...values, password: e.target.value })} />
                </div>
                <button type="submit" className="btn">Регистрация</button>
            </form>
            <a href="/signin" className="signin">Вход</a>
        </div>
    )
}