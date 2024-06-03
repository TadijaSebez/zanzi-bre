import { useEffect, useState } from "react"

export default function EditNote({onShareCallback, onSaveCallback, onNewNoteRequested, inFocuse}) {
    const [title, setTitle] = useState('')
    const [content, setContent] = useState('')
    const [note, setNote] = useState(undefined)

    useEffect(() => {
        setNote(inFocuse)
        if (inFocuse == undefined) {
            setTitle('')
            setContent('')
        }
        else {
            setTitle(inFocuse?.title)
            setContent(inFocuse?.content)
        }
    }, [inFocuse])

    return(
        <div className="edit-note-wrapper">
            {note == undefined && <div className="dashed-card h-100 flex column justify-center center">
                <p className="card-title">Note not Selected</p>
                <p className="card-body neutral v-spacer-s">Select a note from the list to edit it or create a new note!</p>
                <button className="solid-accent-button" onClick={onNewNoteRequested}>New Note</button>
            </div>}
            {note != undefined && <div className="card showing h-100">
                <div className="v-spacer-s flex center gap-xs">
                    <div className="input-wrapper regular-border w-100" style={{paddingRight:'16px', paddingLeft:'8px'}}>
                        <span className="material-symbols-outlined icon input-icon">titlecase</span>
                        <input placeholder="Note Title" value={title} onChange={(e) => setTitle(e.target.value)}/>
                    </div>
                    <button className="surface-button gap-xs flex center" onClick={() => onShareCallback(undefined)}>
                        <span className="material-symbols-outlined icon small-icon">share</span>
                        Share
                    </button>
                    <button className="solid-accent-button gap-xs flex center" style={{paddingRight:'16px'}} onClick={() => onSaveCallback(note)}>
                        <span className="material-symbols-outlined icon small-icon">save</span>
                        Save
                    </button>
    
                </div>
    
                <div className="input-wrapper regular-border v-spacer-xs" style={{height:'calc(100% - 52px)'}}>
                    <span className="material-symbols-outlined icon input-icon" style={{alignSelf:'start', paddingTop:'14px'}}>description</span>
                    <textarea className="h-100" style={{padding:'16px 0 16px 12px'}} placeholder="Content" value={content} onChange={(e) => setContent(e.target.value)}/>
                </div>
            </div>}
        </div>)
}