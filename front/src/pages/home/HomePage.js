import "./HomePage.css"
import NotesList from "./NotesList"

export default function HomePage() {
    return(
        <div className="home-wrapper">
            <div className="notes-column silent-scroll">
                
                <div className="main-header flex center space-between card shadow-elevated">
                    <p className="hero-title">Not<b>ER</b></p>
                    {/* <button className="solid-accent-button gap-xs flex center">
                        <span className="material-symbols-outlined icon">add</span>
                        New Note
                    </button> */}
                    <button className="icon-button">
                        <span className="material-symbols-outlined icon">logout</span>
                    </button>
                </div>
                <NotesList/>
            </div>
            <div className="">

            </div>
        </div>
    )
}