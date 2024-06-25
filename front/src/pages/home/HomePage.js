import { useEffect, useState } from "react"
import { usePopup } from "../../components/pop-up/PopUpFrame"
import EditNote from "./EditNote"
import "./HomePage.css"
import NotesList from "./NotesList"
import ShareSlideOut from "./ShareSlideOut"
import axios from 'axios';
import {API} from '../../environment'

export default function HomePage() {

    const sharePopup = usePopup()

    const handleShare = (note) => {
        sharePopup.showPopup(note)
    }

    const [notes, setNotes] = useState([])
    const token = localStorage.getItem("token")

    const fetchNotes = () => {
        axios.get(API + "/note", { headers: {"Authorization" : `Bearer ${token}`}})
            .then(resp => {
                setNotes(resp.data)
            })
            .catch(e => alert("Couldn't fetch notes"))
    }
    useEffect(() => fetchNotes(), [])

    const [noteInFocuse, setNoteInFocuse] = useState(undefined)
    const saveNote = (note) => {
        setNoteInFocuse(undefined)

        let payload = {
            "noteId": note.noteId,
            "title": note.title,
            "content": note.content
        }

        axios.post(API + "/note", payload, { headers: {"Authorization" : `Bearer ${token}`}})
            .then(resp => {
                console.log(resp)
                console.log(resp.data)
                if(note.noteId == -1) {
                    setNotes(prev => [...prev, {...resp.data, permission: note.permission}])
                } else {
                    setNotes(prevNotes => 
                        prevNotes.map(note => 
                            note.noteId === resp.data.noteId ? {... resp.data, permission: note.permission} : note
                        )
                    );
                }
            })
            .catch(e => alert("Couldn't fetch notes"))

    }
    const newNote = () => {
        setNoteInFocuse({
            noteId: -1,
            title: '', 
            content: '',
            permission: 'owner',
        })
    }
    const discard = () => {
        setNoteInFocuse(undefined)
    }

    return(
        <div className="home-wrapper">
            <div className="notes-column silent-scroll">
                
                <div className="main-header flex center space-between card shadow-elevated">
                    <p className="hero-title">Not<b>ER</b></p>
                    <button className="icon-button" onClick={() => {window.location.href="/"}}>
                        <span className="material-symbols-outlined icon">logout</span>
                    </button>
                </div>
                <NotesList notes={notes} inFocuse={noteInFocuse} onNoteEdit={(note) => setNoteInFocuse(note)}/>
            </div>
            <EditNote onShareCallback={handleShare} inFocuse={noteInFocuse} onSaveCallback={saveNote} onNewNoteRequested={newNote} onDiscardCallback={discard}/>
            <ShareSlideOut popup={sharePopup} note={noteInFocuse}/>
        </div>
    )
}