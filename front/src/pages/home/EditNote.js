import { useState } from "react"

export default function EditNote() {
    const [title, setTitle] = useState("")
    const [content, setContent] = useState("")

    return(
            <div className="card edit-note-wrapper">
                <div className="v-spacer-s flex center gap-xs">
                    <div className="input-wrapper regular-border w-100" style={{paddingRight:'16px', paddingLeft:'8px'}}>
                        <span className="material-symbols-outlined icon input-icon">titlecase</span>
                        <input placeholder="Note Title" value={title} onChange={(e) => setTitle(e.target.value)}/>
                    </div>
                    <button className="surface-button gap-xs flex center">
                        <span className="material-symbols-outlined icon small-icon">share</span>
                        Share
                    </button>
                    <button className="solid-accent-button gap-xs flex center" style={{paddingRight:'16px'}}>
                        <span className="material-symbols-outlined icon small-icon">save</span>
                        Save
                    </button>

                </div>

                <div className="input-wrapper regular-border v-spacer-xs" style={{height:'calc(100% - 52px)'}}>
                    <span className="material-symbols-outlined icon input-icon" style={{alignSelf:'start', paddingTop:'14px'}}>description</span>
                    <textarea className="h-100" style={{padding:'16px 0 16px 12px'}} placeholder="Content" value={content} onChange={(e) => setContent(e.target.value)}/>
                </div>
            </div>
    )
}