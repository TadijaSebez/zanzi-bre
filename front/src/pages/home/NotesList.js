import { useEffect, useState } from "react"

export default function NotesList() {
    const [notes, setNotes] = useState([
        // {
        //     title: "Note 1",
        //     content: "This an example of the Note!"
        // }
    ])

    useEffect(() => fetchNotes(), [])

    const fetchNotes = () => {
        let temp = []
        for(let i = 1; i <= 20; i++) {
            temp.push({
                title: "Note " + i, 
                content: "This an example of the Note! This one is a lot longer so we can see how does the thingy work",
                premission: i < 10 ? 'OWNER' : i < 15 ? 'EDITOR' : 'VIEWER'
            })
        }
        setNotes(temp)
    }

    return(
        <div style={{padding:'0 12px 12px 12px', marginTop:'12px'}}>
            {notes.map(item => {
                return(
                    <div className="card clickable v-spacer-xs" style={{position:'relative'}}>
                        <p className="card-title">{item.title}</p>
                        <p className="card-body neutral note-peek">{item.content}</p>
                        
                        {item?.premission == 'OWNER' && <div className="solid-chip small-chip tr-chip flex center gap-xxs owner-chip">
                            <span className="material-symbols-outlined icon small-icon">person</span>
                            <p className="card-body">Owner</p>
                        </div>}

                        {item?.premission == 'EDITOR' && <div className="solid-chip small-chip tr-chip flex center gap-xxs editor-chip">
                            <span className="material-symbols-outlined icon small-icon">edit</span>
                            <p className="card-body">Editor</p>
                        </div>}

                        {item?.premission == 'VIEWER' && <div className="solid-chip small-chip tr-chip flex center gap-xxs viewer-chip">
                            <span className="material-symbols-outlined icon small-icon">visibility</span>
                            <p className="card-body">Viewer</p>
                        </div>}
                    </div>
                )
            })}

            {notes?.length == 0 && <div className="dashed-card flex column center">
                <p className="card-title">Notes Empty</p>
                <p className="card-body">Add a new note to begin your jurney!</p>
            </div>}
        </div>
    )
}