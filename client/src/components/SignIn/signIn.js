import {useState} from "react";
import {Navigate, useNavigate} from "react-router-dom";
import {IsAuth} from "../../isAuth/isAuth";


export const SignIn = () => {
    const navigate = useNavigate();
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
        fetch('http://localhost:8000/api/auth/signin', {
            method: 'POST',
            credentials: 'include',
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify(values),
        }).then((response) => {
            if (!response.ok) {
                setError(true)
            }
        }).then(() => navigate('/notes'))
    }
    if (error) return <h1>Ошибка при входе в аккаунт</h1>
    if (isLoading) return <h1>Вход в аккаунт...</h1>
    return (
        <div className="container">
            <h1 className="auth">Вход</h1>
            <form method="post" onSubmit={submitHandler}>
                <div className="inputs">
                    <input className="input" type="email" placeholder="Почта" required value={values.email} onChange={e => setValues({ ...values, email: e.target.value })} />
                    <input className="input" type="password" placeholder="Пароль" required value={values.password} onChange={e => setValues({ ...values, password: e.target.value })} />
                </div>
                <button type="submit" className="btn">Вход</button>
            </form>
            <a href="/signup" className="signin">Регистрация</a>
        </div>
    )
}