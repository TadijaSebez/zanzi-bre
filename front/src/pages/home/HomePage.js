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

    return(
        <div className="home-wrapper">
            <div className="notes-column silent-scroll">
                
                <div className="main-header flex center space-between card shadow-elevated">
                    <p className="hero-title">Not<b>ER</b></p>
                    <button className="icon-button">
                        <span className="material-symbols-outlined icon">logout</span>
                    </button>
                </div>
                <NotesList/>
            </div>
            <EditNote onShareCallback={handleShare}/>
            <ShareSlideOut popup={sharePopup}/>
        </div>
    )
}