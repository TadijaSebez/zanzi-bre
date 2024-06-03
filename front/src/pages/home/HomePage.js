import { useEffect, useState } from "react"
import { usePopup } from "../../components/pop-up/PopUpFrame"
import EditNote from "./EditNote"
import "./HomePage.css"
import NotesList from "./NotesList"
import ShareSlideOut from "./ShareSlideOut"

export default function HomePage() {

    const sharePopup = usePopup()

    const handleShare = (note) => {
        sharePopup.showPopup(note)
    }

    const [notes, setNotes] = useState([])
    const fetchNotes = () => {
        let temp = []
        for(let i = 1; i <= 20; i++) {
            temp.push({
                id: i,
                title: "Note " + i, 
                content: "This an example of the Note! This one is a lot longer so we can see how does the thingy work",
                premission: i < 10 ? 'OWNER' : i < 15 ? 'EDITOR' : 'VIEWER'
            })
        }
        setNotes(temp)
    }
    // useEffect(() => fetchNotes(), [])

    const [noteInFocuse, setNoteInFocuse] = useState(undefined)
    const saveNote = (note) => {
        setNoteInFocuse(undefined)
        if(note.id == null) {
            note.id = notes.length
            setNotes(prev => [...prev, note])
        }
    }
    const newNote = () => {
        setNoteInFocuse({
            id: null,
            title: '', 
            content: '',
            premission: 'OWNER'
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
                    <button className="icon-button">
                        <span className="material-symbols-outlined icon">logout</span>
                    </button>
                </div>
                <NotesList notes={notes} inFocuse={noteInFocuse} onNoteEdit={(note) => setNoteInFocuse(note)}/>
            </div>
            <EditNote onShareCallback={handleShare} inFocuse={noteInFocuse} onSaveCallback={saveNote} onNewNoteRequested={newNote} onDiscardCallback={discard}/>
            <ShareSlideOut popup={sharePopup}/>
        </div>
    )
}