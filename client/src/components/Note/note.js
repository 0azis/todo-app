import "./note.css"
import {useEffect, useState} from "react";
import {IsAuth} from '../../isAuth/isAuth'
import { Navigate } from 'react-router-dom';
import ImageNote from '../../note.svg'

import {ThreeCircles} from "react-loader-spinner";

export const Note = ({noteID, text, date}) => {
    const submitHandler = e => {
        e.preventDefault()

        fetch(`https://notes-mzmm.onrender.com/api/notes/?note_id=${noteID}`, {
            method: 'DELETE',
            credentials: 'include',
            headers: {
                "Content-type": "application/json",
            }
        }).then(() => {
            window.location.reload();
        })
    }
    return (
        <div className="note">
            <div className="note_info">
                <p className="note_text">{text}</p>
                <p className="note_date">{date}</p>
            </div>
            <div className="note_functional">
                <form method="post" onSubmit={submitHandler}>
                    <button type="submit" className="btn_delete">Удалить</button>
                </form>
            </div>
        </div>
    )
}


export const GetNotes = () => {
    const [data, setData] = useState(null)
    const [newNote, setNewNote] = useState({'text': ''})
    const [isLoading, setLoading] = useState(true)
    const [error, setError] = useState('')

    useEffect(() => {
        const fetchData = async () => {
            await fetch('https://notes-mzmm.onrender.com/api/notes', { credentials: 'include', method: 'GET' })
                .then(stats => stats.json())
                .then(data => setData(data))
                .catch(err => setError(err))

            setLoading(false)
        }

        fetchData()
    }, [])

    const submitHandler = e => {
        e.preventDefault()
        setLoading(true)
        fetch('https://notes-mzmm.onrender.com/api/notes', {
            method: 'POST',
            credentials: 'include',
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify(newNote),
        }).then(() => {
            window.location.reload();
        })
    }

    if (!IsAuth()) {
        return <Navigate to="/signup" />
    }

    if (isLoading) {
        return (
            <>
            <div className="notes">
                <form method="post" onSubmit={submitHandler}>
                    <input className="input_newnote" placeholder="Type a new note..." required value={newNote.text} onChange={e => setNewNote({ ...newNote, text: e.target.value })} />
                </form>
            </div>
            <div className="loader">
                <ThreeCircles width={50} height={50} color="#DBD8AE" />
            </div>
            </>
        )
    }
    if (error) return error
    console.log(data)
    if (data != null) {
        return (
            <>
                <div className="notes">
                    <form method="post" onSubmit={submitHandler}>
                        <input className="input_newnote" placeholder="Type a new note..." required value={newNote.text} onChange={e => setNewNote({ ...newNote, text: e.target.value })} />
                    </form>
                    {data.slice(0).reverse().map(note => (
                        <Note noteID={note.noteID} text={note.text} date={new Date(note.date).toLocaleDateString()} />
                    ))}
                </div>
            </>
        )
    }
    return (
        <div className="notes">
            <form method="post" onSubmit={submitHandler}>
                <input className="input_newnote" placeholder="Type a new note..." required value={newNote.text} onChange={e => setNewNote({ ...newNote, text: e.target.value })} />
                <h2 className="intro">Создайте вашу первую заметку</h2>
                <img src={ImageNote} width={70} />
            </form>

        </div>
    )
}